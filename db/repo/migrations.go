package repo

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // File source for migrations
)

// Migrate function applies migrations to the database.
func Migrate(dbURL string, migrationsPath string) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}
	urlPath := "file://" + filepath.ToSlash(absPath)

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		urlPath,
		dbURL,
	)
	if err != nil {
		return err
	}

	// Properly close migration, logging any errors
	defer func() {
		_, closeErr := m.Close()
		if closeErr != nil {
			log.Println("failed to close migration:", closeErr)
		}
	}()

	// Apply migrations
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

// MigrateDown function rolls back migrations from the database.
func MigrateDown(dbURL string, migrationsPath string) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		"file://"+filepath.ToSlash(absPath),
		dbURL,
	)
	if err != nil {
		return err
	}

	// Properly close migration, logging any errors
	defer func() {
		_, closeErr := m.Close()
		if closeErr != nil {
			log.Println("failed to close migration:", closeErr)
		}
	}()

	// Rollback migrations
	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
