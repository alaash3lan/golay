package main

import (
	"golay/internal/domain/user/seeder"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Set up the database connection
	dsn := "root:alaa@tcp(127.0.0.1:3306)/golay?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	seeder.SeedUsers(db)
}
