package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"example/slash/src/pkg/bot"
	"example/slash/src/pkg/db"
	"example/slash/src/pkg/profiles"
	"example/slash/src/shared"

	"github.com/bwmarrin/discordgo"

	_ "github.com/lib/pq"
)

func saveMineResults(mydb *sql.DB, discordId string, currTime time.Time, newBalance int) error {
	query := "UPDATE profiles SET last_mined_at = $1, balance = $2 WHERE discord_id = $3;"
	_, err := mydb.Exec(query, currTime, newBalance, discordId)
	if err != nil {
		return err
	}
	return nil
}

func calcTimeSinceLastMine(currTime time.Time, lastMinedAt time.Time) int64 {
	_, offset := currTime.Zone()
	unixDiff := currTime.Unix() - lastMinedAt.Unix()
	return unixDiff + int64(offset)
}

func getMineResults(secondsSinceLastMine int64) (int, string) {
	minutesSinceLastMine := secondsSinceLastMine / 60

	if secondsSinceLastMine == 69 {
		return 6969, "nice"
	}
	if minutesSinceLastMine > 50 {
		return shared.RandomNum(750, 800), "amazing"
	}
	if minutesSinceLastMine > 20 && minutesSinceLastMine <= 50 {
		return shared.RandomNum(400, 750), "great"
	}
	if minutesSinceLastMine > 10 && minutesSinceLastMine <= 20 {
		return shared.RandomNum(200, 300), "good"
	}
	if minutesSinceLastMine > 5 && minutesSinceLastMine <= 10 {
		return shared.RandomNum(75, 100), "decent"
	}
	if minutesSinceLastMine > 1 && minutesSinceLastMine <= 5 {
		return shared.RandomNum(25, 50), "fine"
	}
	if secondsSinceLastMine > 30 && secondsSinceLastMine <= 60 {
		return shared.RandomNum(10, 17), "sad"
	}
	if secondsSinceLastMine > 10 && secondsSinceLastMine <= 30 {
		return shared.RandomNum(4, 7), "pitiful"
	}
	return shared.RandomNum(1, 3), "depressing"
}

func Mine(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("init mine command...")

	mydb, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer mydb.Close()

	profile, err := profiles.Find(mydb, i.Member.User.ID)
	if err != nil {
		bot.Send("Mine Failed. No account found for Player '"+i.Member.User.GlobalName+"'", s, i)
		return
	}

	currTime := time.Now()
	secondsSinceLastMine := calcTimeSinceLastMine(currTime, profile.Last_Mined_At)

	mineAmt, mineDesc := getMineResults(secondsSinceLastMine)

	newBalance := profile.Balance + mineAmt

	mineUpdateError := saveMineResults(mydb, i.Member.User.ID, currTime, newBalance)
	if mineUpdateError != nil {
		log.Println(mineUpdateError)
		return
	}

	shared.PrettyLogJSON(profile)
	bot.Send("⛏️ " +profile.Discord_Global_Name+ " found "+fmt.Sprint(mineAmt)+" gold bits! That mine was "+mineDesc+"!\n" + "💰 "+ profile.Discord_Global_Name+ "'s wallet now has "+fmt.Sprint(newBalance)+"gb!", s, i)
}
