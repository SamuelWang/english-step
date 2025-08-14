package database

import (
	"english-step/website-api/database/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init initializes the database connection and GORM instance for PostgreSQL.
func Init() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable" // default for local development
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC", host, user, password, name, port, sslmode)
 
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Printf("Connected to database")
	return db, nil
}

// MigrateDev runs auto-migration for all models in the development environment.
func MigrateDev(globalDB *gorm.DB) {
	env := os.Getenv("ENV")
	if env != "development" {
		log.Println("Skipping migration: not in development environment")
		return
	}

	err := globalDB.AutoMigrate(
		&models.SynonymExplanation{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	log.Println("Database migration completed (development environment)")
}
