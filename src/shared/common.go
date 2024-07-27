package shared

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

func PrettyLogJSON(value any) {
	json, err := json.MarshalIndent(value, "", "    ")
	if err != nil {
		fmt.Print("PrettyLogJSON Error")
	}
	fmt.Println(string(json))
}

func RandomNum(min, max int) int {
	return rand.IntN(max-min) + min
}

func RollDice(max int) int {
	return rand.IntN(max-1) + 1
}
