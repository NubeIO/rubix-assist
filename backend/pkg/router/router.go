package router

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/controller"
	dbase "github.com/NubeIO/rubix-updater/database"
	"github.com/NubeIO/rubix-updater/pkg/logger"
	"github.com/NubeIO/rubix-updater/service/auth"
	jwt "github.com/appleboy/gin-jwt/v2"
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
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders: []string{
			"X-FLOW-Key", "Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
		},
		ExposeHeaders:          []string{"Content-Length"},
		AllowCredentials:       true,
		AllowAllOrigins:        true,
		AllowBrowserExtensions: true,
		MaxAge:                 12 * time.Hour,
	}))

	GDB := new(dbase.DB)
	GDB.DB = db

	//GDB := dbase.DB{GORM: db}
	api := controller.Controller{DB: GDB, WS: ws}
	identityKey := "id"

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "go-proxy-service",
		Key:           []byte(os.Getenv("JWTSECRET")),
		Timeout:       time.Hour * 1000,
		MaxRefresh:    time.Hour,
		IdentityKey:   identityKey,
		PayloadFunc:   auth.MapClaims,
		Authenticator: api.Login,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization",
		TimeFunc:    time.Now,
	})

	//web socket route
	r.GET("/ws", func(c *gin.Context) {
		ws.HandleRequest(c.Writer, c.Request)
	})

	ws.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println(string(msg))
		ws.Broadcast(msg)
	})

	admin := r.Group("/api")
	hosts := admin.Group("/hosts")
	hosts.Use(authMiddleware.MiddlewareFunc())
	{
		hosts.GET("/schema", api.HostsSchema)
		hosts.GET("/", api.GetHosts)
		hosts.POST("/", api.CreateHost)
		hosts.GET("/:id", api.GetHost)
		hosts.PATCH("/:id", api.UpdateHost)
		hosts.DELETE("/:id", api.DeleteHost)
		hosts.DELETE("/drop", api.DropHosts)
		hosts.POST("/ops", api.MassOperations)
	}

	r.POST("/api/users", api.AddUser)
	r.POST("/api/users/login", authMiddleware.LoginHandler)

	users := admin.Group("/users")
	users.Use(authMiddleware.MiddlewareFunc())
	{
		users.GET("/schema", api.UsersSchema)
		users.GET("/", api.GetUsers)
		users.GET("/:id", api.GetUser)
		users.PATCH("/:id", api.UpdateUser)
		users.DELETE("/:id", api.DeleteUser)
		users.DELETE("/drop", api.DropUsers)
	}

	wiresPlat := admin.Group("/plat")
	wiresPlat.Use(authMiddleware.MiddlewareFunc())
	{
		wiresPlat.GET("/schema", api.RubixPlatSchema)
	}

	proxyRubix := r.Group("/api/rubix/proxy")
	{
		proxyRubix.GET("/*proxy", api.RubixProxyRequest)
		proxyRubix.POST("/*proxy", api.RubixProxyRequest)
		proxyRubix.PUT("/*proxy", api.RubixProxyRequest)
		proxyRubix.PATCH("/*proxy", api.RubixProxyRequest)
		proxyRubix.DELETE("/*proxy", api.RubixProxyRequest)
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
		apps.POST("/full_install", api.AppsFullInstall)
	}

	programs := r.Group("/api/programs")
	{
		programs.GET("/nodejs", api.NodeJsVersion)
		programs.POST("/nodejs", api.NodeJsInstall)
	}

	uf := r.Group("/api/ufw")
	{
		uf.POST("/install", api.InstallUFW)
	}

	bios := r.Group("/api/bios")
	{
		bios.POST("/install", api.InstallBios)
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
	}

	return r
}
