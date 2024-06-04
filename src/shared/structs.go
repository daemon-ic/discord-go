package shared

import (
	"time"

	"github.com/google/uuid"
)

type Weather_Struct struct {
	City    string
	Temp_Lo int
	Temp_Hi int
	Prcp    float32
	Date    string
}

type Profile_Struct struct {
	Id               uuid.UUID
	Created_At       time.Time
	Discord_Username string
	Balance          int
	Last_Mined_At    time.Time
	Collection       string
}
