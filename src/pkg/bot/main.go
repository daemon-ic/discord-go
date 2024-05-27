package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var discordSession *discordgo.Session

func Start() *discordgo.Session {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("error loading .env")
	}

	var err error
	discordSession, err = discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	return discordSession
}
