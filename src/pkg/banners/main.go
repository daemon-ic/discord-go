package banners

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"example/slash/src/shared"

	"github.com/lib/pq"
)

func GetAll(mydb *sql.DB) ([]shared.Banner_Struct, error) {
	query := "SELECT * FROM banners;"
	rows, err := mydb.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var bannerRows []shared.Banner_Struct

	for rows.Next() {
		var banner shared.Banner_Struct

		if err := rows.Scan(
			&banner.Id,
			&banner.Created_At,
			&banner.Name,
			pq.Array(&banner.Entities),
			&banner.ImageUrl,
			&banner.Description,
			&banner.Price,
		); err != nil {
			log.Fatal(err)
		}
		bannerRows = append(bannerRows, banner)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bannerRows, nil
}

func GetPage(mydb *sql.DB, page int) ([]shared.Banner_Struct, error) {
	if page < 1 {
		return nil, errors.New("invalid page num : " + strconv.Itoa(page))
	}

	offset := page - 1
	query := "SELECT * FROM banners ORDER BY created_at DESC LIMIT 1 OFFSET " + strconv.Itoa(offset) + ";"
	rows, err := mydb.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var bannerRows []shared.Banner_Struct

	for rows.Next() {
		var banner shared.Banner_Struct

		if err := rows.Scan(
			&banner.Id,
			&banner.Created_At,
			&banner.Name,
			pq.Array(&banner.Entities),
			&banner.ImageUrl,
			&banner.Description,
			&banner.Price,
		); err != nil {
			log.Fatal(err)
		}
		bannerRows = append(bannerRows, banner)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bannerRows, nil
}
