package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DBInit() {
	DB_Uri := os.Getenv("DB_URI")

	db, err := sql.Open("postgres", DB_Uri)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	DB = db
}

func GetDb() (*sql.DB, error) {
	err := DB.Ping()

	if err != nil {
		return nil, err
	}

	return DB, err
}
