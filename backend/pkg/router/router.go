package router

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/controller"
	"github.com/NubeIO/rubix-updater/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
	"io"
	"os"
	"time"
)



func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()
	var ws = melody.New()
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
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:    []string{
		"X-FLOW-Key", "Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins: true,
		AllowBrowserExtensions: true,
		MaxAge: 12 * time.Hour,
	}))

	//web socket route
	r.GET("/ws", func(c *gin.Context) {
		ws.HandleRequest(c.Writer, c.Request)
	})

	ws.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println(string(msg))
		ws.Broadcast(msg)
	})

	//r.Use(middleware.CORS())
	api := controller.Controller{DB: db, WS: ws}
	hosts := r.Group("/api/hosts")
	{
		hosts.GET("/", api.GetHosts)
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
	apps := r.Group("/api/apps")
	{
		apps.GET("/:id", api.GetApps)
		apps.POST("/download/:id", api.DownloadApp)
		apps.POST("/full_install/:id", api.FullInstall)
		apps.POST("/install/:id", api.InstallApp)
		apps.GET("/state/:id", api.GetDownloadSate)
		apps.DELETE("/state/:id", api.DeleteDownloadSate)
	}

	git := r.Group("/api/git")
	{
		git.GET("/:id", api.GitGetRelease)
	}

	plugins := r.Group("/api/plugins")
	{
		plugins.POST("/full_install/:id", api.UpdatePlugins)
		plugins.POST("/upload/:id", api.UploadPlugins)
		plugins.POST("/delete/:id", api.DeleteAllPlugins)
	}
	upload := r.Group("/api/upload")
	{
		upload.POST("/:id", api.UploadZip)
		upload.POST("/unzip/:id", api.Unzip)
		//token.POST("/", api.CreateToken)
		//token.GET("/:id", api.GetToken)
		//token.PATCH("/:id", api.UpdateToken)
		//token.DELETE("/:id", api.DeleteToken)
	}

	proxyRubix := r.Group("/api/rubix/proxy")
	{
		proxyRubix.GET("/*id", api.FFPoints)
		proxyRubix.PATCH("/*id", api.FFPoints)
	}

	flowFramework := r.Group("/api/ff")
	{
		flowFramework.GET("/localstorage_flow_network/:id", api.FFGetLocalStorage)
		flowFramework.PATCH("/localstorage_flow_network/:id", api.FFUpdateLocalStorage)
		flowFramework.GET("/points/proxy/*id", api.FFPoints)
		flowFramework.PATCH("/points/proxy/*id", api.FFPoints)

	}
	return r
}
