package pgrepo

import (
	"errors"
	"fmt"
	"log"

	"xamss.diploma/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrate struct {
	Config  *config.Config
	Migrate *migrate.Migrate
}

func NewMigrate(config *config.Config) *Migrate {
	m := new(Migrate)

	migr, err := migrate.New(
		fmt.Sprintf("file://%s", config.DB.MigrationPath),
		fmt.Sprintf(config.DB.DSN))
	if err != nil {
		log.Fatal(err)
	}

	m.Migrate = migr
	m.Config = config

	return m
}

func (m *Migrate) Up() error {
	if err := m.Migrate.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate up err: %w", err)
		}
	}

	return nil
}

func (m *Migrate) Down() error {
	if err := m.Migrate.Down(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate down err: %w", err)
		}
	}

	return nil
}

func (m *Migrate) MigrateToVersion(version uint) error {
	log.Printf("migrate to version: %d started", version)
	if err := m.Migrate.Migrate(version); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Printf("migrate to version: %d finished with err", version)
			return fmt.Errorf("migrate down err: %w", err)
		}
	}
	log.Printf("migrate to version: %d finished", version)
	return nil
}
