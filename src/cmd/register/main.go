package main

import (
	"flag"
	"log"

	"example/slash/src/pkg/bot"
	"example/slash/src/pkg/commands"

	"github.com/bwmarrin/discordgo"
)

var GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")

func main() {
	discordSession := bot.Start()

	err := discordSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.List))

	for idx, value := range commands.List {
		cmd, err := discordSession.ApplicationCommandCreate(discordSession.State.User.ID, *GuildID, value)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", value.Name, err)
		} else {
			registeredCommands[idx] = cmd
			log.Printf("added command: '%v'", value.Name)
		}
	}
	log.Println("Sucessufully added commands")
	defer discordSession.Close()
}
