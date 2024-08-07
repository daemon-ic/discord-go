package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"example/slash/src/pkg/bot"
	"example/slash/src/pkg/db"
	"example/slash/src/pkg/profiles"
	"example/slash/src/shared"

	"github.com/bwmarrin/discordgo"
)

// tentative rules
// roll 3d20
// win if you get about 30

// if you get a 20, you dont lose your wager if you lose
// if the numbers are all the same, you dont lose your wager if you lose, and get 150 gb bonus

// if you get 2 20s, you get a 500 gb bonus
// if you get 3 20s, you get a 1000 gb bonus

func getWager(i *discordgo.InteractionCreate) (int64, error) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	wager, ok := optionMap["wager"]
	if ok {
		return wager.IntValue(), nil
	} else {
		return 0, errors.New("issue getting wager option")
	}
}

func roll(label string, s *discordgo.Session, i *discordgo.InteractionCreate) int {
	result := shared.RollDice(20)
	time.Sleep(500 * time.Millisecond)

	msg := "🎲 " + i.Member.User.GlobalName + "s " + label + " Roll" + ": " + fmt.Sprint(result)

	log.Println(msg)
	bot.SendMsg(msg, s, i)
	return result
}

func countOccurences(rollArray []int) map[int]int {
	dict := make(map[int]int)

	for _, num := range rollArray {
		dict[num] = dict[num] + 1
	}

	return dict
}

func playGame(s *discordgo.Session, i *discordgo.InteractionCreate) (int, bool, bool, float64) {
	const WIN_THRESHOLD = 30
	gameWon := false
	immunity := false
	bonusMultiplier := 1.0
	score := 0
	code := ""

	roll1 := roll("1st", s, i)
	roll2 := roll("2nd", s, i)
	roll3 := roll("3rd", s, i)

	time.Sleep(500 * time.Millisecond)
	rollMap := countOccurences([]int{roll1, roll2, roll3})

	score = roll1 + roll2 + roll3
	gameWon = score >= WIN_THRESHOLD

	if score/3 == roll1 {
		if score == 3 {
			code = "MAXIMUM_FAILURE"
		} else {
			code = "TRIPLE"
			bonusMultiplier = 1.2
			immunity = true
		}
	}

	if rollMap[20] > 0 {
		code = "CRITICAL_SUCCESS"
		immunity = true
	}

	if rollMap[20] == 2 {
		code = "DOUBLE_SUCCESS"
		bonusMultiplier = 1.5
	}

	if rollMap[20] == 3 {
		code = "MAXIMUM_SUCCESS"
		bonusMultiplier = 3
	}

	codeDict := make(map[string]string)

	codeDict["MAXIMUM_FAILURE"] = " has gotten a Maximum Failure... Haven't added anything for this yet so just feel ashamed for now"
	codeDict["CRITICAL_SUCCESS"] = " has gotten a Critical Success! Rolling a 20 spares your wallet on losing!"
	codeDict["TRIPLE"] = " has gotten a Three Of A Kind! This grants a 20% bonus on winning, and spares your wallet on losing!"
	codeDict["DOUBLE_SUCCESS"] = " has gotten a Double Success! Rolling 2x 20's will grant a 50% bonus on winning!"
	codeDict["MAXIMUM_SUCESS"] = " has gotten a Maximum Success!!! Rolling 3x 20's will grant a 3x bonus on winning!"

	if _, ok := codeDict[code]; ok {
		bot.SendMsg(i.Member.User.GlobalName+codeDict[code], s, i)
	}

	return score, gameWon, immunity, bonusMultiplier
}

func generateResultMsgs(discordName string, score int, gameWon bool, immunity bool, gainWithBonus int, newBalance int) string {
	resultMsg := ""

	if gameWon {
		resultMsg += "🔥 Congrats! "
	} else {
		if immunity {
			resultMsg += "😓 Close call! "
		}
		if !immunity {
			resultMsg += "💸 Aw shucks! "
		}
	}

	resultMsg += discordName + " got a " + fmt.Sprint(score) + "! "

	if immunity && !gameWon {
		resultMsg += discordName + " was spared from losing any Gold Bits!\n"
	}
	if gameWon {
		resultMsg += "Won " + fmt.Sprint(gainWithBonus) + "gb!\n"
	} else {
		resultMsg += "Lost " + fmt.Sprint(gainWithBonus) + "gb!\n"
	}
	resultMsg += "💰 " + discordName + "'s wallet now has " + fmt.Sprint(newBalance) + "gb!\n"
	return resultMsg
}

func saveCasinoResults(mydb *sql.DB, discordId string, newBalance int) error {
	query := "UPDATE profiles SET balance = $1 WHERE discord_id = $2;"
	_, err := mydb.Exec(query, newBalance, discordId)
	if err != nil {
		return err
	}
	return nil
}

func Casino(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("init casino command...")

	mydb, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer mydb.Close()

	profile, err := profiles.Find(mydb, i.Member.User.ID)
	if err != nil {
		bot.Send("Casino Command Failed. No account found for Player '"+i.Member.User.GlobalName+"'", s, i)
		return
	}

	log.Println("found user profile:", profile.Discord_Global_Name)
	log.Println("pre gamble balance:", profile.Balance)

	unverifiedWager, err := getWager(i)
	if err != nil {
		bot.Send("Casino Command Failed. Issue Retrieving Wager for Player '"+i.Member.User.GlobalName+"'", s, i)
		return
	}

	if unverifiedWager <= 0 {
		bot.Send("Invalid Wager for Player '"+i.Member.User.GlobalName+"'", s, i)
		return
	}

	wager := math.Min(float64(unverifiedWager), float64(profile.Balance))
	log.Println("wager:", wager)

	score, gameWon, immunity, bonusMultiplier := playGame(s, i)

	gainWithBonus := int(math.Round(float64(wager) * bonusMultiplier))
	balanceChange := gainWithBonus

	if !immunity && !gameWon {
		balanceChange = -1 * gainWithBonus
	}

	log.Println("balance change:", balanceChange)

	newBalance := profile.Balance + balanceChange
	saveCasinoResults(mydb, profile.Discord_Id, newBalance)

	resultMsg := generateResultMsgs(profile.Discord_Global_Name, score, gameWon, immunity, gainWithBonus, newBalance)
	bot.SendMsg(resultMsg, s, i)
	bot.Send("✨ Game Complete\n", s, i)
	bot.EditInteractionResp("fuck you", s, i)

	// the results just didnt show the labels
	// one of the field results was a memeory value, and thats because i didnt try to log an actual field
	// apperantly if i try to log an object without hitting its field, the result will be a memory value.
	// "%+v\n" will be the value, combined with the fmt library in order to tell me the field names which apperanly is important
	// in order to use '%+v\n, i have to use fmt.Printf'
}
