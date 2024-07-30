package handlers

import (
	"encoding/json"
	httpReq "example/slash/src/api"
	"github.com/bwmarrin/discordgo"
	"net/url"

	"example/slash/src/shared"
)

// func _generateDisplayContentNative(displayItems shared.Shop_Display) (*string, *[]discordgo.MessageComponent) {
// 	Content := &displayItems.ItemName
//
// 	// ComponentsNative :=
//
// 	Components := &[]discordgo.MessageComponent{
// 		discordgo.ActionsRow{
// 			Components: []discordgo.MessageComponent{
// 				discordgo.Button{
// 					Emoji: &discordgo.ComponentEmoji{
// 						Name: "⬅️",
// 					},
// 					CustomID: "banner_prev",
// 					Style:    discordgo.SecondaryButton,
// 				},
// 				discordgo.Button{
// 					Label:    "Buy!",
// 					CustomID: "banner_buy",
// 					Style:    discordgo.PrimaryButton,
// 				},
// 				discordgo.Button{
// 					Emoji: &discordgo.ComponentEmoji{
// 						Name: "➡️",
// 					},
// 					CustomID: "banner_next",
// 					Style:    discordgo.SecondaryButton,
// 				},
// 			},
// 		},
// 	}
// 	return Content, Components
// }

func (shop Shop_Instance) InitialDisplayNative(displayItems shared.Shop_Display) {

	interactionResp := shared.Interaction_Response_Json{
		Type: 4, // CHANNEL_MESSAGE_WITH_SOURCE
		Data: shared.Interaction_Callback_Data{
			Content: displayItems.ItemName,
			Components: shared.Components{
				discordgo.Button{
					Emoji: &discordgo.ComponentEmoji{
						Name: "⬅️",
					},
					CustomID: "banner_prev",
					Style:    2, // secondary button
				},
				discordgo.Button{
					Label:    "Buy!",
					CustomID: "banner_buy",
					Style:    1, // primary button
				},
				discordgo.Button{
					Emoji: &discordgo.ComponentEmoji{
						Name: "➡️",
					},
					CustomID: "banner_next",
					Style:    2, // secondary button
				},
			},
		},
	}

	uri := shared.DiscordBaseUri + "/interactions/" + shop.I.ID + "/" + shop.I.Token + "/callback"

	// the headers to pass
	headers := map[string]string{
		"Accept": "application/vnd.github.v3+json",
	}

	// the query parameters to pass
	queryParams := url.Values{}
	queryParams.Add("per_page", "1")

	// the body to pass
	buf, err := json.Marshal(&interactionResp)
	if err != nil {
		panic(err)
	}

	//body := bytes.NewBufferString(interactionJson)

	// the type to unmarshal the response into
	var respType map[string]interface{}

	// resp, err := http.Post(uri, "application/json", interactionResp)
	resp, err := httpReq.Make(uri, "POST", headers, queryParams, buf, respType)
	if err != nil {
		panic(err)
	}

	println(resp)
}

// func (shop Shop_Instance) InitialDisplayNative(displayItems shared.Shop_Display) {
// 	Content, Components := _generateDisplayContent(displayItems)
//
// 	interactionErr := shop.S.InteractionRespond(
// 		shop.I.Interaction,
// 		&discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Content:    *Content,
// 				Components: *Components,
// 				// Attachments: attachments,
// 			},
// 		})
// 	if interactionErr != nil {
// 		panic(interactionErr)
// 	}
// }

// func (shop Shop_Instance) ChangeDisplayNative(displayItems shared.Shop_Display) {
// 	Content, Components := _generateDisplayContent(displayItems)
//
// 	interactionErr := shop.S.InteractionRespond(
// 		shop.I.Interaction,
// 		&discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseUpdateMessage,
// 			Data: &discordgo.InteractionResponseData{
// 				Content:    *Content,
// 				Components: *Components,
// 			},
// 		})
// 	if interactionErr != nil {
// 		panic(interactionErr)
// 	}
// }
