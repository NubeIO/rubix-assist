package router

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/controller"
	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/auth"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := fmt.Sprintf("%s %s [%d]: %s", ctx.Request.Method, ctx.Request.URL, 404, "api not found")
		ctx.JSON(http.StatusNotFound, controller.Message{Message: message})
	}
}

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.NoRoute(NotFound())
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
	makeStore, _ := appstore.New(&appstore.Store{App: &installer.App{}, DB: appDB})
	api := controller.Controller{DB: appDB, Store: makeStore}
	identityKey := "uuid"
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
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

	admin := r.Group("/api")
	appStore := admin.Group("/store")

	{
		appStore.GET("/apps", api.ListAppsWithVersions)
		appStore.GET("/apps/details", api.ListAppsBuildDetails)
		appStore.POST("/add", api.AddUploadStoreApp)
		appStore.POST("/upload/plugin", api.UploadStorePlugin)

	}

	edgeApps := admin.Group("/edge")
	{
		edgeApps.GET("/system/product", api.EdgeProductInfo)
		edgeApps.GET("/apps", api.EdgeListApps)
		edgeApps.GET("/apps/services", api.EdgeListAppsAndService)
		edgeApps.GET("/apps/services/nube", api.EdgeListNubeServices)
		edgeApps.POST("/apps/add", api.AddUploadEdgeApp)
		edgeApps.POST("/plugins/add", api.EdgeUploadPlugin)
		edgeApps.POST("/apps/service/upload", api.GenerateUploadEdgeService)
		edgeApps.POST("/apps/service/install", api.InstallEdgeService)
		edgeApps.DELETE("/apps", api.EdgeUninstallApp)
	}

	edgeAppsControl := admin.Group("/edge/control")
	{
		edgeAppsControl.POST("/action", api.EdgeCtlAction)              // start, stop
		edgeAppsControl.POST("/action/mass", api.EdgeServiceMassAction) // mass operation start, stop
		edgeAppsControl.POST("/status", api.EdgeCtlStatus)              // isRunning, isInstalled and so on
		edgeAppsControl.POST("/status/mass", api.EdgeServiceMassStatus) // mass isRunning, isInstalled and so on
	}

	locations := admin.Group("/locations")

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

	Tasks := admin.Group("/tasks")
	//Tasks.Use(authMiddleware.MiddlewareFunc())
	{
		Tasks.GET("/schema", api.TasksSchema)
		Tasks.GET("/", api.GetTasks)
		Tasks.POST("/", api.CreateTask)
		Tasks.GET("/:uuid", api.GetTask)
		Tasks.PATCH("/:uuid", api.UpdateTask)
		Tasks.DELETE("/:uuid", api.DeleteTask)
		Tasks.DELETE("/drop", api.DropTasks)
	}

	messages := admin.Group("/transactions")
	//messages.Use(authMiddleware.MiddlewareFunc())
	{
		messages.GET("/schema", api.TransactionsSchema)
		messages.GET("/", api.GetTransactions)
		messages.POST("/", api.CreateTransaction)
		messages.GET("/:uuid", api.GetTransaction)
		messages.PATCH("/:uuid", api.UpdateTransaction)
		messages.DELETE("/:uuid", api.DeleteTransaction)
		messages.DELETE("/drop", api.DropTransactions)
	}

	tools := admin.Group("/tools")
	//tools.Use(authMiddleware.MiddlewareFunc())
	{

		tools.GET("/edgeapi/ip/schema", api.EdgeIPSchema)
		tools.POST("/edgeapi/ip", api.EdgeSetIP)
		tools.POST("/edgeapi/ip/dhcp", api.EdgeSetIP)

	}

	files := admin.Group("/files")
	{
		files.POST("/upload", api.UploadFile)
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

	wires := admin.Group("/wires")
	{
		wires.POST("/upload", api.WiresUpload)
		wires.GET("/backup", api.WiresBackup)
	}

	system := admin.Group("/system")
	{
		system.GET("/ping", api.Ping)
		system.GET("/time", api.HostTime)
	}

	networking := admin.Group("/networking")
	{
		networking.GET("/networks", api.Networking)
		networking.GET("/interfaces", api.GetInterfacesNames)
		networking.GET("/internet", api.InternetIP)
	}
	r.Any("/proxy/*proxyPath", api.Proxy)

	return r
}
