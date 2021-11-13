package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yakuter/ugin/controller"
	"github.com/yakuter/ugin/pkg/logger"
	"github.com/yakuter/ugin/pkg/middleware"
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
	posts := r.Group("/posts")
	{
		posts.GET("/", api.GetPosts)
	}
	return r
}
