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

// type Mine_Reward struct {
// 	Amount int
// 	Level  string
// }
