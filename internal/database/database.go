package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func DBInit() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_Uri := os.Getenv("DB_URI")

	db, err := sql.Open("postgres", DB_Uri)

	if err != nil {
		panic(err)
	}

	DB = db

	err = DB.Ping()

	if err != nil {
		panic(err)
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
