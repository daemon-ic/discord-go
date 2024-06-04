package handlers

import (
	"log"
	"time"

	"example/slash/src/pkg/db"
	"example/slash/src/pkg/profiles"
	"example/slash/src/shared"

	"github.com/bwmarrin/discordgo"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	mydb, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer mydb.Close()

	profiles.Register(mydb, shared.Profile_Struct{
		Id:               uuid.New(),
		Created_At:       time.Now(),
		Discord_Username: "test",
		Balance:          100,
		Last_Mined_At:    time.Now(),
		Collection:       "{}",
	})

	profiles.GetAll(mydb)

	log.Printf("%+v", i)
	log.Printf("%+v", s)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Register triggered",
		},
	})
}
