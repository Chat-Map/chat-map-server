package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

// New creates a new postgres connection with the given connection string
func New(url string) (*sql.DB, error) {
	// Creates a new postgres conn
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %s", err)
	}
	return conn, nil
}

// Migrate runs the migrations on the database
func Migrate(conn *sql.DB, migrationDir string) error {
	// Create a new pg instance
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create driver for migration: %s", err)
	}

	// Load migration files
	m, err := migrate.NewWithDatabaseInstance(migrationDir, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create new pg instance for migration: %s", err)
	}

	// Push migration changes
	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
		log.Println("Database is up to date, No migration made")
	}
	return nil
}
