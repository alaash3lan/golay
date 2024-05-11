package main

import (
	"log"
	"os"

	userDependency "golay/internal/domain/user/dependency"
	"golay/internal/routes"
	"golay/migrations"

	// userService "golay/internal/service/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	migrations.RunMigrations(db)
	router := gin.Default()

	userHandler, err := userDependency.SetupUserDependencies(db)
	if err != nil {
		panic(err)
	}

	// Setup user routes
	routes.SetupUserRoutes(router, userHandler)

	// Run the server

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} // Start the server
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
