package main

import (
	"log"
	"time"

	"example/slash/src/pkg/db"
	"example/slash/src/shared"

	"github.com/google/uuid"
)

func getEntity() shared.Entity_Struct {
	e := shared.Entity_Struct{
		Id:         uuid.New(),
		Created_At: time.Now(),
		Name:       "batumaki",
		ImageUrl:   "",
		FlavorText: "",
		Power:      300,
		Rarity:     10,
		Prints:     0,
	}
	return e
}

func main() {
	log.Println("inity entity import")

	mydb, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer mydb.Close()

	e := getEntity()
	query := "INSERT INTO entities (id, created_at, name, image_url, flavor_text, power, rarity, prints) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

	_, err = mydb.Exec(
		query,
		e.Id,
		e.Created_At,
		e.Name,
		e.ImageUrl,
		e.FlavorText,
		e.Power,
		e.Rarity,
		e.Prints)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("import success")
}
