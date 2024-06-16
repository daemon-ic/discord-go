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

func SendMsg(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.ChannelMessageSend(i.ChannelID, content)
}

func Send(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	// sends and interaction response which apperantly ends the interaction ?
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
