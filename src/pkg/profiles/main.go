package profiles

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"example/slash/src/shared"
)

func Exists() {}

func Register(mydb *sql.DB, p shared.Profile_Struct) string {
	query := "INSERT INTO profiles (id, created_at, discord_username, discord_id, discord_global_name, balance, last_mined_at, collection) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
	_, err := mydb.Exec(
		query,
		p.Id,
		p.Created_At,
		p.Discord_Username,
		p.Discord_Id,
		p.Discord_Global_Name,
		p.Balance,
		p.Last_Mined_At,
		p.Collection)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "unique_discord_ids"):
			log.Println(err)
			return "ERR:unique_discord_ids"

		default:
			log.Fatal(err)
		}
	}

	log.Println("success: " + query)
	return "OK"
}

func Find(mydb *sql.DB, id string) (shared.Profile_Struct, error) {
	var profile shared.Profile_Struct

	if err := mydb.QueryRow(
		"SELECT * FROM profiles WHERE discord_id = $1;",
		id,
	).Scan(
		&profile.Id,
		&profile.Created_At,
		&profile.Discord_Username,
		&profile.Balance,
		&profile.Last_Mined_At,
		&profile.Collection,
		&profile.Discord_Id,
		&profile.Discord_Global_Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return profile, errors.New("no profile found")
		}
		return profile, err
	}
	return profile, nil
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
			&profile.Discord_Id,
			&profile.Discord_Global_Name,
		); err != nil {
			log.Fatal(err)
		}
		profileRows = append(profileRows, profile)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	shared.PrettyLogJSON(profileRows)
}
