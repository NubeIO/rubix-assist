package router

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/controller"
	dbase "github.com/NubeIO/rubix-assist/database"
	dbhandler "github.com/NubeIO/rubix-assist/pkg/handler"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/service/auth"
	"github.com/NubeIO/rubix-assist/service/em"
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
	appDB := &dbase.DB{
		DB: db,
	}
	dbHandler := &dbhandler.Handler{
		DB: appDB,
	}
	dbhandler.Init(dbHandler) // TODO REMOVE THIS
	edgeManger := em.New(&em.EdgeManager{
		DB: appDB,
	})
	api := controller.Controller{DB: appDB, WS: ws, Edge: edgeManger}
	identityKey := "uuid"
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
		err := ws.HandleRequest(c.Writer, c.Request)
		fmt.Println(err)
		//if err != nil {
		//	return
		//}
	})

	ws.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println(string(msg))
		ws.Broadcast(msg)
	})

	admin := r.Group("/api")

	locations := admin.Group("/locations")
	//hosts.Use(authMiddleware.MiddlewareFunc())
	{
		locations.GET("/schema", api.GetLocationSchema)
		locations.GET("/", api.GetLocations)
		locations.POST("/wizard", api.CreateLocationWizard)
		locations.POST("/", api.CreateLocation)
		locations.GET("/:uuid", api.GetLocation)
		locations.PATCH("/:uuid", api.UpdateLocation)
		locations.DELETE("/:uuid", api.DeleteLocation)
		locations.DELETE("/drop", api.DropLocations)
	}

	hostNetworks := admin.Group("/networks")
	//hosts.Use(authMiddleware.MiddlewareFunc())
	{
		hostNetworks.GET("/schema", api.GetNetworkSchema)
		hostNetworks.GET("/", api.GetHostNetworks)
		hostNetworks.POST("/", api.CreateHostNetwork)
		hostNetworks.GET("/:uuid", api.GetHostNetwork)
		hostNetworks.PATCH("/:uuid", api.UpdateHostNetwork)
		hostNetworks.DELETE("/:uuid", api.DeleteHostNetwork)
		hostNetworks.DELETE("/drop", api.DropHostNetworks)
	}

	hosts := admin.Group("/hosts")
	//hosts.Use(authMiddleware.MiddlewareFunc())
	{
		hosts.GET("/schema", api.GetHostSchema)
		hosts.GET("/", api.GetHosts)
		hosts.POST("/", api.CreateHost)
		hosts.GET("/:uuid", api.GetHost)
		hosts.PATCH("/:uuid", api.UpdateHost)
		hosts.DELETE("/:uuid", api.DeleteHost)
		hosts.DELETE("/drop", api.DropHosts)
		hosts.POST("/ops", api.MassOperations)
	}

	r.POST("/api/users", api.AddUser)
	r.POST("/api/users/login", authMiddleware.LoginHandler)

	users := admin.Group("/users")
	//users.Use(authMiddleware.MiddlewareFunc())
	{
		users.GET("/schema", api.UsersSchema)
		users.GET("/", api.GetUsers)
		users.GET("/:uuid", api.GetUser)
		users.PATCH("/:uuid", api.UpdateUser)
		users.DELETE("/:uuid", api.DeleteUser)
		users.DELETE("/drop", api.DropUsers)
	}

	teams := admin.Group("/teams")
	//teams.Use(authMiddleware.MiddlewareFunc())
	{
		teams.GET("/schema", api.TeamsSchema)
		teams.GET("/", api.GetTeams)
		teams.POST("/", api.CreateTeam)
		teams.GET("/:uuid", api.GetTeam)
		teams.PATCH("/:uuid", api.UpdateTeam)
		teams.DELETE("/:uuid", api.DeleteTeam)
		teams.DELETE("/drop", api.DropTeams)
	}

	alerts := admin.Group("/tasks")
	//alerts.Use(authMiddleware.MiddlewareFunc())
	{
		alerts.GET("/schema", api.AlertsSchema)
		alerts.GET("/", api.GetAlerts)
		alerts.POST("/", api.CreateAlert)
		alerts.GET("/:uuid", api.GetAlert)
		alerts.PATCH("/:uuid", api.UpdateAlert)
		alerts.DELETE("/:uuid", api.DeleteAlert)
		alerts.DELETE("/drop", api.DropAlerts)
	}

	messages := admin.Group("/transactions")
	//messages.Use(authMiddleware.MiddlewareFunc())
	{
		messages.GET("/schema", api.MessagesSchema)
		messages.GET("/", api.GetMessages)
		messages.POST("/", api.CreateMessage)
		messages.GET("/:uuid", api.GetMessage)
		messages.PATCH("/:uuid", api.UpdateMessage)
		messages.DELETE("/:uuid", api.DeleteMessage)
		messages.DELETE("/drop", api.DropMessages)
	}

	tools := admin.Group("/tools")
	//tools.Use(authMiddleware.MiddlewareFunc())
	{

		tools.GET("/edge/ip/schema", api.EdgeIPSchema)
		tools.POST("/edge/ip", api.EdgeSetIP)
		tools.POST("/edge/ip/dhcp", api.EdgeSetIP)
		tools.POST("/zip", api.ZipUpload)

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
		token.GET("/:uuid", api.GetToken)
		token.PATCH("/:uuid", api.UpdateToken)
		token.DELETE("/:uuid", api.DeleteToken)
		token.DELETE("/drop", api.DropTokens)
	}
	git := r.Group("/api/git")
	{
		git.GET("/:uuid", api.GitGetRelease)
	}

	edge := r.Group("/api/edge")
	{
		edge.POST("/apps/install", api.InstallApp)
		edge.POST("/apps/pipeline/install", api.InstallPipeline)
	}

	return r
}
