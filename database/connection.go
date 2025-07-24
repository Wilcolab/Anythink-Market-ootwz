package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"

	"github.com/Wilcolab/Anythink-Market-ootwz/config"
	_ "github.com/lib/pq"
)

// Connect establishes a connection to the PostgreSQL database
func Connect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("✅ Connected to PostgreSQL database")
	return db, nil
}

// RunMigrations executes all SQL migration files in the migrations directory
func RunMigrations(db *sql.DB, migrationsPath string) error {
	log.Println("🔄 Running database migrations...")

	files, err := filepath.Glob(filepath.Join(migrationsPath, "*.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration files: %w", err)
	}

	if len(files) == 0 {
		log.Println("📝 No migration files found")
		return nil
	}

	// Sort files to ensure they run in order
	sort.Strings(files)

	for _, file := range files {
		log.Printf("📄 Executing migration: %s", filepath.Base(file))

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}
	}

	log.Println("✅ All migrations completed successfully")
	return nil
}
