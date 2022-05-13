package router

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/controller"
	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/pkg/config"
	dbhandler "github.com/NubeIO/rubix-assist/pkg/handler"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/service/auth"
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
	conf := config.GetConfig()
	f, err := os.OpenFile(fmt.Sprintf("%s/rubix.access.log", conf.GetAbsDataDir()), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
	gg := new(dbhandler.Handler)
	gg.DB = GDB
	dbhandler.Init(gg)
	//GDB := dbase.DB{GORM: db}
	api := controller.Controller{DB: GDB, WS: ws}
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
		ws.HandleRequest(c.Writer, c.Request)
	})

	ws.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println(string(msg))
		ws.Broadcast(msg)
	})

	admin := r.Group("/api")
	hosts := admin.Group("/hosts")
	//hosts.Use(authMiddleware.MiddlewareFunc())
	{
		hosts.GET("/schema", api.HostsSchema)
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
	users.Use(authMiddleware.MiddlewareFunc())
	{
		users.GET("/schema", api.UsersSchema)
		users.GET("/", api.GetUsers)
		users.GET("/:uuid", api.GetUser)
		users.PATCH("/:uuid", api.UpdateUser)
		users.DELETE("/:uuid", api.DeleteUser)
		users.DELETE("/drop", api.DropUsers)
	}

	teams := admin.Group("/teams")
	teams.Use(authMiddleware.MiddlewareFunc())
	{
		teams.GET("/schema", api.TeamsSchema)
		teams.GET("/", api.GetTeams)
		teams.POST("/", api.CreateTeam)
		teams.GET("/:uuid", api.GetTeam)
		teams.PATCH("/:uuid", api.UpdateTeam)
		teams.DELETE("/:uuid", api.DeleteTeam)
		teams.DELETE("/drop", api.DropTeams)
	}

	alerts := admin.Group("/alerts")
	alerts.Use(authMiddleware.MiddlewareFunc())
	{
		alerts.GET("/schema", api.AlertsSchema)
		alerts.GET("/", api.GetAlerts)
		alerts.POST("/", api.CreateAlert)
		alerts.GET("/:uuid", api.GetAlert)
		alerts.PATCH("/:uuid", api.UpdateAlert)
		alerts.DELETE("/:uuid", api.DeleteAlert)
		alerts.DELETE("/drop", api.DropAlerts)
	}

	messages := admin.Group("/messages")
	messages.Use(authMiddleware.MiddlewareFunc())
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
	tools.Use(authMiddleware.MiddlewareFunc())
	{
		tools.GET("/endpoints", api.ToolsEndPoints)
		tools.GET("/edge/ip/schema", api.EdgeIPSchema)
		tools.POST("/edge/ip", api.EdgeSetIP)
		tools.POST("/edge/ip/dhcp", api.EdgeSetIP)
		tools.GET("/arch", api.ToolsGetArch)
		//tools.GET("/nodejs", api.NodeJsVersion)
		//tools.POST("/modbus/config", api.ModbusIOConfig)
		//tools.POST("/modbus/poll", api.ModbusPoll)
	}

	//ff := admin.Group("/ff")
	//ff.Use(authMiddleware.MiddlewareFunc())
	//{
	//	ff.POST("/flow_networks", api.FFFlowNetworkWizard)
	//	ff.POST("/networks", api.FFNetworkWizard)
	//}

	//master := admin.Group("/master")
	//master.Use(authMiddleware.MiddlewareFunc())
	//{
	//	master.GET("/discover/schema", api.RubixDiscoverSchema)
	//	master.GET("/slaves/schema", api.RubixMasterSchema)
	//}

	//wiresPlat := admin.Group("/plat")
	//wiresPlat.Use(authMiddleware.MiddlewareFunc())
	//{
	//	wiresPlat.GET("/schema", api.RubixPlatSchema)
	//}

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
	}
	//apps := r.Group("/api/apps")
	//{
	//	apps.POST("/full_install", api.AppsFullInstall)
	//}

	//programs := r.Group("/api/programs")
	//{
	//	programs.GET("/nodejs", api.NodeJsVersion)
	//	programs.POST("/nodejs", api.NodeJsInstall)
	//}

	uf := r.Group("/api/ufw")
	{
		uf.POST("/install", api.UFWInstall)
		uf.POST("/ports/open", api.UFWAddPort)
		uf.POST("/enable", api.UFWEnable)
		uf.POST("/disable", api.UFWDisable)
	}

	bios := r.Group("/api/bios")
	{
		bios.POST("/install", api.InstallBios)
		//bios.GET("/update_check", api.RubixServiceCheck)
		//bios.PUT("/upgrade_and_check", api.RubixServiceUpdate)
	}

	git := r.Group("/api/git")
	{
		git.GET("/:uuid", api.GitGetRelease)
	}

	//plugins := r.Group("/api/plugins")
	{
		//plugins.POST("/full_install", api.PluginFullInstall)
		//plugins.POST("/full_uninstall", api.PluginFullUninstall)
		//plugins.POST("/upgrade", api.FlowFrameworkUpgrade)
		//plugins.POST("/upload/:uuid", api.UploadPlugins)
		//plugins.POST("/delete/:uuid", api.DeleteAllPlugins)
	}
	upload := r.Group("/api/upload")
	{
		upload.POST("/:uuid", api.UploadZip)
		upload.POST("/unzip/:uuid", api.Unzip)
	}

	return r
}
