package shared

import (
	"time"

	"github.com/google/uuid"
)

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
