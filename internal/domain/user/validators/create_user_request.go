// validators/createUserRequest.go
package validators

import (
	"golay/internal/domain/user/model"

	"github.com/go-playground/validator/v10"
)

// CreateUserRequest represents the request data for creating a new user.
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Validate validates the CreateUserRequest fields and returns a map of field names to error messages.
func (r *CreateUserRequest) Validate() map[string]string {
	validate := validator.New()
	err := validate.Struct(r)

	if err == nil {
		return nil
	}

	validationErrors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		errorMessage := err.Tag()
		switch errorMessage {
		case "required":
			validationErrors[fieldName] = "This field is required"
		case "email":
			validationErrors[fieldName] = "Please enter a valid email address"
		case "min":
			validationErrors[fieldName] = "The password must be at least 8 characters"
		default:
			validationErrors[fieldName] = "Invalid value"
		}
	}

	return validationErrors
}

// ToUserModel converts the CreateUserRequest to a User model.
func (r *CreateUserRequest) ToUserModel() *model.User {
	return &model.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}