package repository

import (
	// "fmt"

	// "golay/helper"
	"golay/internal/domain/user/model"

	"gorm.io/gorm"
	"golay/internal/utils"
)

// UserRepository is an interface that defines the operations a user repository should implement.
type UserRepository interface {
	CreateUser(user *model.User) error
	GetUser(id uint) (*model.User, error)
	ListUsers() ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

// GORMUserRepository is a repository implementation using GORM.
type GORMUserRepository struct {
	db *gorm.DB
}

// NewGORMUserRepository creates a new GORMUserRepository instance.
func NewGORMUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{db: db}
}

// CreateUser creates a new user.
func (r *GORMUserRepository) CreateUser(user *model.User) error {

	
	pass,err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pass
	r.db.Create(user)
	return nil
}

// GetUser retrieves a user by ID.
func (r *GORMUserRepository) GetUser(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}
func (r *GORMUserRepository) ListUsers() ([]*model.User, error) {
	var users []*model.User
	err := r.db.Find(&users).Error
	return users, err
}

// UpdateUser updates an existing user.
func (r *GORMUserRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

// DeleteUser deletes a user by ID.
func (r *GORMUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
