package db

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

//go:embed migrations
var migrationFiles embed.FS

// Migrate runs the migrations
func Migrate(postgresDbUrl string) error {
	sourceDriver, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		return errors.Wrap(err, "Failed to create migration source driver")
	}

	db, err := sql.Open("postgres", postgresDbUrl)
	if err != nil {
		return errors.Wrap(err, "Failed to open database")
	}

	databaseDriver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "auctions_migrations",
	})
	if err != nil {
		return errors.Wrap(err, "Failed to create migration database driver")
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", databaseDriver)
	if err != nil {
		return errors.Wrap(err, "Failed to create migration instance")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "Failed to run migrations")
	}

	return nil
}
