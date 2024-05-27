package main

import (
	"flag"
	"log"

	"example/slash/src/pkg/bot"
)

func init() { flag.Parse() }

var GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")

func init() { flag.Parse() }

func main() {
	discordSession := bot.Start()

	err := discordSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Removing commands...")
	registeredCommands, err := discordSession.ApplicationCommands(discordSession.State.User.ID, *GuildID)
	if err != nil {
		log.Fatalf("could not fetch registered commands: '%v'", err)
	}

	for _, value := range registeredCommands {
		err := discordSession.ApplicationCommandDelete(discordSession.State.User.ID, *GuildID, value.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", value.Name, err)
		} else {
			log.Printf("deleted command: '%v'", value.Name)
		}
	}
	log.Println("sucessfully deleted commands")
	defer discordSession.Close()
}
