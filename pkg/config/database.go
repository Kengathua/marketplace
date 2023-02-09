package config

import (
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgresDriver.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Printf("Successfully connected to db\n")
}
