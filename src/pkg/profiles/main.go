package profiles

import (
	"database/sql"
	"log"

	"example/slash/src/shared"
)

func Exists() {}

func Register(mydb *sql.DB, p shared.Profile_Struct) {
	query := "INSERT INTO profiles (id, created_at, discord_username, balance, last_mined_at, collection) VALUES ($1, $2, $3, $4, $5, $6);"
	_, err := mydb.Exec(
		query,
		p.Id,
		p.Created_At,
		p.Discord_Username,
		p.Balance,
		p.Last_Mined_At,
		p.Collection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success: " + query)
}

func GetAll(mydb *sql.DB) {
	query := "SELECT * FROM profiles;"
	rows, err := mydb.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var profileRows []shared.Profile_Struct

	for rows.Next() {
		var profile shared.Profile_Struct

		if err := rows.Scan(
			&profile.Id,
			&profile.Created_At,
			&profile.Discord_Username,
			&profile.Balance,
			&profile.Last_Mined_At,
			&profile.Collection,
		); err != nil {
			log.Fatal(err)
		}
		profileRows = append(profileRows, profile)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", profileRows)
}
