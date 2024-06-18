package main

import (
	"log"
	"time"

	"example/slash/src/pkg/db"
	"example/slash/src/shared"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func getBanners() shared.Banner_Struct {
	b := shared.Banner_Struct{
		Id:          uuid.New(),
		Created_At:  time.Now(),
		Name:        "PROTOTYPE",
		Entities:    []string{"ENTITY_ID_1", "ENTITY_ID_2", "60211b35-27b6-4e85-91dc-555e94da8585"},
		ImageUrl:    "",
		Description: "This is a prototype of a banner used for testing purposes.",
		Price:       1000,
	}
	return b
}

func main() {
	log.Println("importing banner(s)")

	mydb, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer mydb.Close()

	b := getBanners()
	query := "INSERT INTO banners (id, created_at, name, entities, image_url, description, price) VALUES ($1, $2, $3, $4, $5, $6, $7);"

	_, err = mydb.Exec(
		query,
		&b.Id,
		&b.Created_At,
		&b.Name,
		(*pq.StringArray)(&b.Entities),
		&b.ImageUrl,
		&b.Description,
		&b.Price,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully imported banner(s)")
}
