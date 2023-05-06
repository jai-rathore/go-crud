package handlers

import (
	"github.com/jai-rathore/lists/services"
)

// UserHandler handles user requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
