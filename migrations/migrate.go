package main

import (
	"os"

	"github.com/Kengathua/marketplace/pkg/config"
	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func init() {
	config.ConnectToDb()
}

func RunDBMigrationUp(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("Successfully migrated the database")
}

func main() {
	config.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;`)
	RunDBMigrationUp(os.Getenv("MIGRATION_URL"), os.Getenv("POSTGRESQL_URL"))
	config.DB.AutoMigrate(&models.CustomerOrder{})
	config.DB.AutoMigrate(&models.CustomerOrderItem{})
}
