package handlers

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"example/slash/src/pkg/profiles"
	"example/slash/src/shared"

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

func (shop Shop_Instance) FormatDataToDisplay(displayItems shared.Shop_Display) {
}

func (shop Shop_Instance) InitialDisplay(displayItems shared.Shop_Display) {
	Content, Components := _generateDisplayContent(displayItems)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	attachments := &[]*discordgo.MessageAttachment{
		{
			ID:       timestamp,
			URL:      displayItems.ImageUrl,
			Filename: "banner.png",
		},
	}

	// interactionErr := shop.S.InteractionRespond(
	// 	shop.I.Interaction,
	// 	&discordgo.InteractionResponse{
	// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 		Data: &discordgo.InteractionResponseData{
	// 			Content:     *Content,
	// 			Components:  *Components,
	// 			Attachments: attachments,
	// 		},
	// 	})
	interactionErr := shop.S.InteractionRespond(
		shop.I.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:     *Content,
				Components:  *Components,
				Attachments: attachments,
			},
		})
	if interactionErr != nil {
		panic(interactionErr)
	}
}

func (shop Shop_Instance) ChangeDisplay(displayItems shared.Shop_Display) {
	Content, Components := _generateDisplayContent(displayItems)

	// ID          string `json:"id"`
	// URL         string `json:"url"`
	// ProxyURL    string `json:"proxy_url"`
	// Filename    string `json:"filename"`
	// ContentType string `json:"content_type"`
	// Width       int    `json:"width"`
	// Height      int    `json:"height"`
	// Size        int    `json:"size"`
	// Ephemeral   bool   `json:"ephemeral"`
	//
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	attachments := &[]*discordgo.MessageAttachment{
		{
			ID:       timestamp,
			URL:      displayItems.ImageUrl,
			Filename: "banner.png",
		},
	}

	interactionErr := shop.S.InteractionRespond(
		shop.I.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:     *Content,
				Components:  *Components,
				Attachments: attachments,
			},
		})
	if interactionErr != nil {
		panic(interactionErr)
	}
}
