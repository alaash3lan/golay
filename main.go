package main

import (
	"log"
	"os"

	userDependency "golay/internal/domain/user/dependency"
	"golay/internal/domain/user/model"
	"golay/internal/routes"

	// userService "golay/internal/service/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

)

func main() {

	// Set up the database connection
	dsn := "root:alaa@tcp(127.0.0.1:3306)/golay?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

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
	}	// Start the server
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
