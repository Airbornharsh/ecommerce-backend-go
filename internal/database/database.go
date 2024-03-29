package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func DBInit() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	DB_Uri := os.Getenv("DB_URI")

	db, err := sql.Open("postgres", DB_Uri)

	if err != nil {
		fmt.Println(err)
	}

	DB = db

	err = DB.Ping()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to database")
	// MakeTable(DB)
	// DropTable(DB)
}

func GetDb() (*sql.DB, error) {
	err := DB.Ping()

	if err != nil {
		return nil, err
	}

	return DB, err
}
