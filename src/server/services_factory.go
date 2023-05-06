package server

import "github.com/jai-rathore/lists/services"

var (
	listService services.IListService
	userService services.IUserService
)

func InitAppServices(s *Server) {
	listService = services.NewListService(listRepository)
	userService = services.NewUserService(userRepository)
}
