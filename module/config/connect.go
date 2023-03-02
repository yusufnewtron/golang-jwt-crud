package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDataBase() (*sql.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	// DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	// DbPort := os.Getenv("DB_PORT")
	DbSsl := os.Getenv("DB_SSL")

	DBURL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", DbUser, DbPassword, DbName, DbSsl)

	db, err := sql.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil

}
