package migrate

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql" // Import for MySQL driver

	_ "github.com/go-sql-driver/mysql" // Import for MySQL driver registration (underscore for blank import)
)

// Migrator defines the interface for migration operations.
type Migrator interface {
	Up(path string) error
	Down(path string) error
}

// MySQLMigrator implements the Migrator interface for MySQL database.
type MySQLMigrator struct {
	db *sql.DB
}

// NewMySQLMigrator creates a new MySQLMigrator instance.
func NewMySQLMigrator(db *sql.DB) Migrator {
	return &MySQLMigrator{db: db}
}

// Up performs the database migration.
func (m *MySQLMigrator) Up(path string) error {
	driver, err := mysql.WithInstance(m.db, &mysql.Config{}) // Initialize MySQL driver
	if err != nil {
		return fmt.Errorf("failed to initialize MySQL driver: %w", err)
	}

	// Create a new migration instance
	migrationInstance, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path), // Use the file URL scheme for migrations
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := migrationInstance.Up(); err != nil {
		return fmt.Errorf("failed to run migrations up: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

// Down reverts the database migration.
func (m *MySQLMigrator) Down(path string) error {
	driver, err := mysql.WithInstance(m.db, &mysql.Config{}) // Initialize MySQL driver
	if err != nil {
		return fmt.Errorf("failed to initialize MySQL driver: %w", err)
	}

	// Create a new migration instance
	migrationInstance, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path), // Use the file URL scheme for migrations
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	if err := migrationInstance.Down(); err != nil {
		return fmt.Errorf("failed to run migrations down: %w", err)
	}

	log.Println("Migrations reverted successfully")
	return nil
}
