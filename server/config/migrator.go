package config

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *DBConfig) error {
	dbURL := DbMigrationURL(db)
	migrationsPath := fmt.Sprintf("file://%s", "./server/migrations")
	m, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Println(dbURL)
		return fmt.Errorf("ERROR: failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("INFO: no changes to migrate")
			return nil
		}
		return fmt.Errorf("ERROR: failed to apply migrations: %v", err)
	}

	log.Println("INFO: Migrations applied successfully")
	return nil
}

func DbMigrationURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
