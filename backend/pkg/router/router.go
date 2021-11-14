package router

import (
	"github.com/NubeIO/rubix-updater/controller"
	"github.com/NubeIO/rubix-updater/pkg/logger"
	"github.com/NubeIO/rubix-updater/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"os"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()
	// Write gin access log to file
	f, err := os.OpenFile("rubix.access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorf("Failed to create access log file: %v", err)
	} else {
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// Set default middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set custom middlewares
	r.Use(middleware.CORS())
	api := controller.Controller{DB: db}
	// Non-protected routes
	hosts := r.Group("/api/hosts")
	{
		hosts.GET("/", api.GetPosts)
		hosts.POST("/", api.CreateHost)
		hosts.GET("/:id", api.GetHost)
		hosts.PATCH("/:id", api.UpdateHost)
		hosts.DELETE("/:id", api.DeleteHost)
	}
	token := r.Group("/api/tokens")
	{
		token.GET("/", api.GetTokens)
		token.POST("/", api.CreateToken)
		token.GET("/:id", api.GetToken)
		token.PATCH("/:id", api.UpdateToken)
		token.DELETE("/:id", api.DeleteToken)
	}
	return r
}
