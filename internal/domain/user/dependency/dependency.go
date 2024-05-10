package dependency

import (
	"golay/internal/domain/user/handler"
	"golay/internal/domain/user/repository"
	"golay/internal/domain/user/service"

	"gorm.io/gorm"
)

func SetupUserDependencies(db *gorm.DB) (*handler.UserHandler, error) {
	userRepo := repository.NewGORMUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	return userHandler, nil
}
