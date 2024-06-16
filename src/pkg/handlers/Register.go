package handlers

import (
	"log"
	"time"

	"example/slash/src/pkg/bot"
	"example/slash/src/pkg/db"
	"example/slash/src/pkg/profiles"
	"example/slash/src/shared"

	"github.com/bwmarrin/discordgo"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("init register command...")

	mydb, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer mydb.Close()

	result := profiles.Register(mydb, shared.Profile_Struct{
		Id:                  uuid.New(),
		Created_At:          time.Now(),
		Discord_Username:    i.Member.User.Username,
		Discord_Id:          i.Member.User.ID,
		Discord_Global_Name: i.Member.User.GlobalName,
		Balance:             100,
		Last_Mined_At:       time.Now(),
		Collection:          "{}",
	})

	message := ""

	switch result {
	case "OK":
		message = "Player '" + i.Member.User.GlobalName + "' successfully registered!"

	case "ERR:unique_discord_ids":
		message = "Player '" + i.Member.User.GlobalName + "' already registered!"

	default:
		message = "Unknown issue registering Player '" + i.Member.User.GlobalName + "'"
	}

	profiles.GetAll(mydb)

	shared.PrettyLogJSON(i)

	bot.Send(message, s, i)
}
