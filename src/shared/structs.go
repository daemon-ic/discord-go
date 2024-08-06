package shared

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

// responding = initial message ----
type Interaction_Respond_Json struct {
	Type int                      `json:"type"`
	Data Interaction_Respond_Data `json:"components"`
}

type Interaction_Respond_Data struct {
	Content    string                       `json:"content"`
	Components []discordgo.MessageComponent `json:"components"`
}

// response = followup message ----
type Interaction_Response_Json struct {
	Type int                       `json:"type"`
	Data Interaction_Callback_Data `json:"data"`
}

type Interaction_Callback_Data struct {
	Content     string                       `json:"content"`
	Components  []discordgo.MessageComponent `json:"components"`
	Attachments Attachment_Object            `json:"attachments"`
}

//type Message_Component struct {
//	Type       int                   `json:"type"`
//	Components this.MessageComponent `json:"components"`
//}

type Attachment_Object struct {
	Id          string `json:"id"`
	Filename    string `json:"filename"`
	Url         string `json:"url"`
	ContentType string `json:"content_type"`
}

type Profile_Struct struct {
	Id                  uuid.UUID
	Created_At          time.Time
	Discord_Username    string
	Discord_Id          string
	Discord_Global_Name string
	Balance             int
	Last_Mined_At       time.Time
	Collection          string
}

type Entity_Struct struct {
	Id         uuid.UUID
	Created_At time.Time
	Name       string
	ImageUrl   string
	FlavorText string
	Power      int
	Rarity     int
	Prints     int
}

type Banner_Struct struct {
	Id          uuid.UUID
	Created_At  time.Time
	Name        string
	Entities    []string
	ImageUrl    string
	Description string
	Price       int
}

type Created_Card_Struct struct {
	Id         string
	Created_At time.Time
	OwnerId    uuid.UUID
	EntityId   uuid.UUID
	Print      int
	Condition  string
}

// type Shop_Button struct {
// 	Action string
// 	Label  string
// }

type Shop_Display struct {
	CurrentPage int
	TotalPages  int
	ItemName    string
	ImageUrl    string
	ItemPrice   int
}

// type Shop_Instance struct {
// 	S *discordgo.Session
// 	I *discordgo.InteractionCreate
// }

// type Mine_Reward struct {
// 	Amount int
// 	Level  string
// }
