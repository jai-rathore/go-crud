package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jai-rathore/lists/config"
	"github.com/jai-rathore/lists/docs"
	"github.com/jai-rathore/lists/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	r := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Authorization", "Content-Type"}
	corsConfig.AllowAllOrigins = true

	r.Use(gin.Recovery())
	r.Use(cors.New(corsConfig))
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/health"}}))

	InitAppRepositories(s)
	InitAppServices(s)

	handlers.NewAdminHandler(r)
	handlers.NewListHandler(listService, r)

	docs.SwaggerInfo.BasePath = "/"

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + s.cfg.Port)
}
