package handlers

import (
	"bytes"
	"database/sql"
	"example/slash/src/pkg/profiles"
	"example/slash/src/shared"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Shop_Instance struct {
	S *discordgo.Session
	I *discordgo.InteractionCreate
}

// needs to move
func (shop Shop_Instance) ValidatePlayer(ggmDB *sql.DB) (shared.Profile_Struct, error) {
	log.Println("fetching player " + shop.I.Member.User.GlobalName)
	return profiles.Find(ggmDB, shop.I.Member.User.ID)
}

func _getImageFromUrl(url string) []byte {
	log.Println("fetching image :" + url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func _generateDisplayContent(displayItems shared.Shop_Display) (*string, *[]discordgo.MessageComponent) {
	Content := &displayItems.ItemName
	Components := &[]discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Emoji: &discordgo.ComponentEmoji{
						Name: "⬅️",
					},
					CustomID: "banner_prev",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					Label:    "Buy!",
					CustomID: "banner_buy",
					Style:    discordgo.PrimaryButton,
				},
				discordgo.Button{
					Emoji: &discordgo.ComponentEmoji{
						Name: "➡️",
					},
					CustomID: "banner_next",
					Style:    discordgo.SecondaryButton,
				},
			},
		},
	}
	return Content, Components
}

func (shop Shop_Instance) InitialDisplay(displayItems shared.Shop_Display) {
	Content, Components := _generateDisplayContent(displayItems)
	ImageFile := _getImageFromUrl(displayItems.ImageUrl)

	interactionErr := shop.S.InteractionRespond(
		shop.I.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    *Content,
				Components: *Components,
				Files: []*discordgo.File{
					{
						Name:        "test.png",
						ContentType: "image/png",
						Reader:      bytes.NewReader(ImageFile),
					},
				},
			},
		})
	if interactionErr != nil {
		panic(interactionErr)
	}
}

func (shop Shop_Instance) ChangeDisplay(displayItems shared.Shop_Display) {
	Content, Components := _generateDisplayContent(displayItems)

	interactionErr := shop.S.InteractionRespond(
		shop.I.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    *Content,
				Components: *Components,
			},
		})
	if interactionErr != nil {
		panic(interactionErr)
	}
}
