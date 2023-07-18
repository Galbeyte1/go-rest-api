package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file couldn't be loaded")
	}
	fmt.Println("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
		dbHost, 
		dbPort, 
		dbUsername, 
		dbTable, 
		dbPassword,
	)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		log.Println("Unable to connecte to Database")
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		log.Println("Unable to ping Database")
		return db, err
	}

	return db, nil
}