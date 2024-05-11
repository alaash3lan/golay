package migrations

import (
	"log"
	"golay/internal/domain/user/model"
	"gorm.io/gorm"
)

// RunMigrations runs auto-migrations for the database schema
func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	
}