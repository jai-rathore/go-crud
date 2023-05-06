package server

import "github.com/jai-rathore/lists/repositories"

var listRepository repositories.IListRepository
var userRepository repositories.IUserRepository

func InitAppRepositories(s *Server) {
	listRepository = repositories.NewListRepository()
	userRepository = repositories.NewUserRepository()
}
