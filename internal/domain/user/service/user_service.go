package service

import (
	"golay/internal/domain/user/model"
	"golay/internal/domain/user/repository"
)

// UserService is a service that handles user-related operations.
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService instance.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user *model.User) error {

	return s.repo.CreateUser(user)
}

// GetUser retrieves a user by ID.
func (s *UserService) ListUsers() ([]*model.User, error) {
	return s.repo.ListUsers()
}
func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.repo.GetUser(id)
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
