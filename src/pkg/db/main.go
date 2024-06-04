package db

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	host   = os.Getenv("POSTGRES_HOST")
	port   = os.Getenv("POSTGRES_POST")
	user   = os.Getenv("POSTGRES_USER")
	dbname = os.Getenv("POSTGRES_DBNAME")
)

func Conn() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	return sql.Open("postgres", psqlInfo)
}

// TODO: need to handle table existence
// SELECT EXISTS (
// SELECT 1
// FROM information_schema.tables
// WHERE table_schema = 'public'
// AND table_name = 'my_table'
// );

//  CREATE TABLE profiles (id varchar(80), created_at date, discord_username varchar(80), balance int, last_mined_at date, collection json);