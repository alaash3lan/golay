package model

import "gorm.io/gorm"

// User represents a user in the system.
type User struct {
	gorm.Model
	Name     string 
	Email    string `gorm:"unique"`
	Password string
}

