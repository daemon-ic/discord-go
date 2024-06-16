package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"example/slash/src/pkg/bot"
	"example/slash/src/pkg/handlers"

	"github.com/bwmarrin/discordgo"
)

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"casino":   handlers.Casino,
	"register": handlers.Register,
	"mine":     handlers.Mine,
	// "basic-command":            handlers.BasicCommand,
	// "basic-command-with-files": handlers.BasicCommandWithFiles,
	// "localized-command":        handlers.LocalizedCommand,
	// "options":                  handlers.Options,
	// "permission-overview":      handlers.PermissionOverview,
	// "subcommands":              handlers.Subcommands,
	// "responses":                handlers.Responses,
	// "followups":                handlers.FollowUps,
}

var GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")

func main() {
	discordSession := bot.Start()

	discordSession.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if handler, ok := Handlers[interaction.ApplicationCommandData().Name]; ok {
			handler(session, interaction)
		}
	})

	discordSession.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", session.State.User.Username, session.State.User.Discriminator)
	})

	err := discordSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	defer discordSession.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down.")
}
