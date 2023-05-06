package services

import (
	"github.com/jai-rathore/lists/models"
	"github.com/jai-rathore/lists/repositories"
)

type IUserService interface {
	GetUserById(id int) (*models.User, error)
}

// UserService handles user requests
type UserService struct {
	userRepository repositories.IUserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// GetUserById returns a user by id
func (u *UserService) GetUserById(id int) (*models.User, error) {
	user, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
