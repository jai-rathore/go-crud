package main

import (
	"github.com/jai-rathore/lists/config"
	"github.com/jai-rathore/lists/logger"
	"github.com/jai-rathore/lists/repositories"
	"github.com/jai-rathore/lists/server"
)

// @title           Lists Service
// @version         1.0

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	config.LoadAppConfig()
	logger.InitLogger()
	_, err := repositories.ConnectDB(config.AppConfig.DBConnectionString)
	if err != nil {
		logger.Fatal("Error connecting to database")
	}
	s := server.NewServer(config.AppConfig)
	s.Run()
}
