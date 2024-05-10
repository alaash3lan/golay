// internal/domain/user/user_seeder.go

package seeder

import (
	"golay/internal/domain/user/model"

	"gorm.io/gorm"
)

// SeedUsers populates the database with sample users.
func SeedUsers(db *gorm.DB) error {
	users := []model.User{
		{Name: "Alice", Email: "alice@example.com", Password: "password1"},
		{Name: "Bob", Email: "bob@example.com", Password: "password2"},
		{Name: "Charlie", Email: "charlie@example.com", Password: "password3"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
