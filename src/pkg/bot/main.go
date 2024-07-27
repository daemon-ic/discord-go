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

func EditMsg(msgId string, content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.ChannelMessageEdit(i.ChannelID, msgId, content)
}

func SendMsg(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.ChannelMessageSend(i.ChannelID, content)
}

func Send(content string, s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	// sends and interaction response which apperantly ends the interaction ?

	// recieve the value of s and i, return value of interaction response... whcih is what i want right?

	// interaction response is going to be the reference of the result of the function.

	// im not sure why i am allowed to pass it as a result

	interactionResp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	}

	s.InteractionRespond(i.Interaction, interactionResp)
	return interactionResp
}

func EditInteractionResp(content string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	newresp := &discordgo.WebhookEdit{
		Content: &content,
	}

	s.InteractionResponseEdit(i.Interaction, newresp)
}
