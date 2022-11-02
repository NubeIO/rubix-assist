package router

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/controller"
	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := fmt.Sprintf("%s %s [%d]: %s", ctx.Request.Method, ctx.Request.URL, 404, "rubix-assist: api not found")
		ctx.JSON(http.StatusNotFound, model.Message{Message: message})
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

	r.POST("/api/users/login", api.Login)
	publicSystemApi := r.Group("/api/system")
	{
		publicSystemApi.GET("/ping", api.SystemPing)
	}

	handleAuth := func(c *gin.Context) { c.Next() }
	if config.Config.Auth() {
		// handleAuth = api.HandleAuth() // TODO add back in auth
	}

	apiRoutes := r.Group("/api", handleAuth)
	apiProxyRoutes := r.Group("/proxy", handleAuth)
	apiProxyRoutes.Any("/*proxyPath", api.Proxy)

	appStore := apiRoutes.Group("/store/apps")
	{
		appStore.GET("/", api.ListAppsWithVersions)
		appStore.POST("/", api.UploadAddOnAppStore)
	}

	storePlugins := apiRoutes.Group("/store/plugins")
	{
		storePlugins.GET("/", api.GetPluginsStorePlugins)
		storePlugins.POST("/", api.UploadPluginStorePlugin)
	}

	edgeBios := apiRoutes.Group("/edge-bios")
	{
		edgeBios.POST("/users/login", api.EdgeBiosLogin)
		edgeBios.POST("/tokens", api.EdgeTokens)
		edgeBios.GET("/system/ping", api.EdgeBiosPing)
	}

	edgeBiosApps := apiRoutes.Group("/edge-bios/edge")
	{
		edgeBiosApps.GET("/upload", api.EdgeBiosEdgeUpload)
	}

	edge := apiRoutes.Group("/edge/system")
	{
		edge.GET("/ping", api.EdgePing)
		edge.GET("/device", api.EdgeGetDeviceInfo)
		edge.GET("/product", api.EdgeProductInfo)
	}

	edgeApps := apiRoutes.Group("/edge/apps")
	{
		edgeApps.GET("/", api.EdgeListApps)
		edgeApps.GET("/status", api.EdgeListAppsStatus)
		edgeApps.POST("/upload", api.EdgeUploadApp)
		edgeApps.POST("/service/upload", api.GenerateServiceFileAndEdgeUpload)
		edgeApps.POST("/service/install", api.EdgeInstallService)
		edgeApps.DELETE("/", api.EdgeUninstallApp)
	}

	edgeAppsControl := apiRoutes.Group("/edge/control")
	{
		edgeAppsControl.POST("/action", api.EdgeSystemCtlAction)
		edgeAppsControl.POST("/status", api.EdgeSystemCtlStatus)
		edgeAppsControl.POST("/action/mass", api.EdgeServiceMassAction)
		edgeAppsControl.POST("/status/mass", api.EdgeServiceMassStatus)
	}

	edgeConfig := apiRoutes.Group("/edge/config")
	{
		edgeConfig.GET("/", api.EdgeReadConfig)
		edgeConfig.POST("/", api.EdgeWriteConfig)
	}

	edgeFiles := apiRoutes.Group("/edge/files")
	{
		edgeFiles.GET("/exists", api.EdgeFileExists)
		edgeFiles.GET("/read", api.EdgeReadFile)
		edgeFiles.POST("/create", api.EdgeCreateFile)
		edgeFiles.POST("/write/string", api.EdgeWriteString)
		edgeFiles.POST("/write/json", api.EdgeWriteFileJson)
		edgeFiles.POST("/write/yml", api.EdgeWriteFileYml)
	}

	edgeDirs := apiRoutes.Group("/edge/dirs")
	{
		edgeDirs.GET("/exists", api.EdgeDirExists)
	}

	edgePlugins := apiRoutes.Group("/edge/plugins")
	{
		edgePlugins.GET("/", api.EdgeListPlugins)
		edgePlugins.POST("/", api.EdgeUploadPlugin)
		edgePlugins.DELETE("/", api.EdgeDeletePlugin)
		edgePlugins.DELETE("/all", api.EdgeDeleteAllPlugins)
	}

	locations := apiRoutes.Group("/locations")
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

	hostNetworks := apiRoutes.Group("/networks")
	{
		hostNetworks.GET("/schema", api.GetNetworkSchema)
		hostNetworks.GET("/", api.GetHostNetworks)
		hostNetworks.POST("/", api.CreateHostNetwork)
		hostNetworks.GET("/:uuid", api.GetHostNetwork)
		hostNetworks.PATCH("/:uuid", api.UpdateHostNetwork)
		hostNetworks.DELETE("/:uuid", api.DeleteHostNetwork)
		hostNetworks.DELETE("/drop", api.DropHostNetworks)
	}

	hosts := apiRoutes.Group("/hosts")
	{
		hosts.GET("/schema", api.GetHostSchema)
		hosts.GET("/", api.GetHosts)
		hosts.POST("/", api.CreateHost)
		hosts.GET("/:uuid", api.GetHost)
		hosts.PATCH("/:uuid", api.UpdateHost)
		hosts.DELETE("/:uuid", api.DeleteHost)
		hosts.DELETE("/drop", api.DropHosts)
	}

	teams := apiRoutes.Group("/teams")
	{
		teams.GET("/schema", api.TeamsSchema)
		teams.GET("/", api.GetTeams)
		teams.POST("/", api.CreateTeam)
		teams.GET("/:uuid", api.GetTeam)
		teams.PATCH("/:uuid", api.UpdateTeam)
		teams.DELETE("/:uuid", api.DeleteTeam)
		teams.DELETE("/drop", api.DropTeams)
	}

	Tasks := apiRoutes.Group("/tasks")
	{
		Tasks.GET("/schema", api.TasksSchema)
		Tasks.GET("/", api.GetTasks)
		Tasks.POST("/", api.CreateTask)
		Tasks.GET("/:uuid", api.GetTask)
		Tasks.PATCH("/:uuid", api.UpdateTask)
		Tasks.DELETE("/:uuid", api.DeleteTask)
		Tasks.DELETE("/drop", api.DropTasks)
	}

	messages := apiRoutes.Group("/transactions")
	{
		messages.GET("/schema", api.TransactionsSchema)
		messages.GET("/", api.GetTransactions)
		messages.POST("/", api.CreateTransaction)
		messages.GET("/:uuid", api.GetTransaction)
		messages.PATCH("/:uuid", api.UpdateTransaction)
		messages.DELETE("/:uuid", api.DeleteTransaction)
		messages.DELETE("/drop", api.DropTransactions)
	}

	// tools := apiRoutes.Group("/tools")
	// //tools.Use(authMiddleware.MiddlewareFunc())
	// {
	//
	//	tools.GET("/edgeapi/ip/schema", api.EdgeIPSchema)
	//	tools.POST("/edgeapi/ip", api.EdgeSetIP)
	//	tools.POST("/edgeapi/ip/dhcp", api.EdgeSetIP)
	//
	// }

	wires := apiRoutes.Group("/wires")
	{
		wires.POST("/upload", api.WiresUpload)
		wires.GET("/backup", api.WiresBackup)
	}

	system := apiRoutes.Group("/system")
	{
		system.GET("/time", api.HostTime)
	}

	networking := apiRoutes.Group("/networking")
	{
		networking.GET("/networks", api.Networking)
		networking.GET("/interfaces", api.GetInterfacesNames)
		networking.GET("/internet", api.InternetIP)
	}

	files := apiRoutes.Group("/files")
	{
		files.GET("/walk", api.WalkFile)
		files.GET("/list", api.ListFiles) // /api/files/list?file=/data
		files.POST("/rename", api.RenameFile)
		files.POST("/copy", api.CopyFile)
		files.POST("/move", api.MoveFile)
		files.POST("/upload", api.UploadFile)
		files.POST("/download", api.DownloadFile)
		files.DELETE("/delete", api.DeleteFile)
		files.DELETE("/delete/all", api.DeleteAllFiles)
	}

	dirs := apiRoutes.Group("/dirs")
	{
		dirs.POST("/create", api.CreateDir)
		dirs.POST("/copy", api.CopyDir)
		dirs.DELETE("/delete", api.DeleteDir)
	}

	zip := apiRoutes.Group("/zip")
	{
		zip.POST("/unzip", api.Unzip)
		zip.POST("/zip", api.ZipDir)
	}

	user := apiRoutes.Group("/users")
	{
		user.PUT("", api.UpdateUser)
		user.GET("", api.GetUser)
	}

	token := apiRoutes.Group("/tokens")
	{
		token.GET("", api.GetTokens)
		token.POST("/generate", api.GenerateToken)
		token.PUT("/:uuid/block", api.BlockToken)
		token.PUT("/:uuid/regenerate", api.RegenerateToken)
		token.DELETE("/:uuid", api.DeleteToken)
	}

	return r
}
