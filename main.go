package main

import (
	"fmt"
	"log"
	"os"

	"golay/internal/config"
	userDependency "golay/internal/domain/user/dependency"
	"golay/internal/routes"
	"golay/migrations"

	// userService "golay/internal/service/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {


	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	
	cfg,err := config.LoadConfig()
	if err != nil {
		logrus.Fatal("cannot load config: ", err)
	}


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
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
