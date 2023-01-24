package router

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/controller"
	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/installer"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/global"
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
		ctx.JSON(http.StatusNotFound, amodel.Message{Message: message})
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
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "host-uuid", "host-name",
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
	global.Installer = installer.New(&installer.Installer{})
	makeStore, _ := appstore.New(&appstore.Store{DB: appDB})
	api := controller.Controller{DB: appDB, Store: makeStore, FileMode: global.Installer.FileMode}

	r.POST("/api/users/login", api.Login)
	publicSystemApi := r.Group("/api/system")
	{
		publicSystemApi.GET("/ping", api.SystemPing)
	}

	handleAuth := func(c *gin.Context) { c.Next() }
	if config.Config.Auth() {
		handleAuth = api.HandleAuth()
	}

	apiRoutes := r.Group("/api", handleAuth)
	apiProxyRoutes := r.Group("/proxy", handleAuth)
	apiProxyRoutes.Any("/*proxyPath", api.Proxy)

	appStore := apiRoutes.Group("/store/apps")
	{
		appStore.POST("", api.UploadAddOnAppStore)
		appStore.GET("/exists", api.CheckAppExistence)
	}

	storePlugins := apiRoutes.Group("/store/plugins")
	{
		storePlugins.GET("", api.GetPluginsStorePlugins)
		storePlugins.POST("", api.UploadPluginStorePlugin)
	}

	edgeBiosApps := apiRoutes.Group("/eb/re")
	{
		edgeBiosApps.POST("/upload", api.EdgeBiosRubixEdgeUpload)
		edgeBiosApps.POST("/install", api.EdgeBiosRubixEdgeInstall)
		edgeBiosApps.GET("/version", api.EdgeBiosGetRubixEdgeVersion)
	}

	edgeApps := apiRoutes.Group("/edge/apps")
	{
		edgeApps.POST("/upload", api.EdgeAppUpload)
		edgeApps.POST("/install", api.EdgeAppInstall)
		edgeApps.POST("/uninstall", api.EdgeAppUninstall)
		edgeApps.GET("/status", api.EdgeListAppsStatus)
		edgeApps.GET("/status/:app_name", api.EdgeGetAppStatus)
	}

	edgePlugins := apiRoutes.Group("/edge/plugins")
	{
		edgePlugins.GET("", api.EdgeListPlugins)
		edgePlugins.POST("/upload", api.EdgeUploadPlugin)
		edgePlugins.POST("/move-from-download-to-install", api.EdgeMoveFromDownloadToInstallPlugins)
		edgePlugins.DELETE("/name/:plugin_name", api.EdgeDeletePlugin)
		edgePlugins.DELETE("/download-plugins", api.EdgeDeleteDownloadPlugins)
	}

	edgeConfig := apiRoutes.Group("/edge/config")
	{
		edgeConfig.GET("", api.EdgeReadConfig)
		edgeConfig.POST("", api.EdgeWriteConfig)
	}

	locations := apiRoutes.Group("/locations")
	{
		locations.GET("/schema", api.GetLocationSchema)
		locations.GET("", api.GetLocations)
		locations.POST("/wizard", api.CreateLocationWizard)
		locations.POST("", api.CreateLocation)
		locations.GET("/:uuid", api.GetLocation)
		locations.PATCH("/:uuid", api.UpdateLocation)
		locations.DELETE("/:uuid", api.DeleteLocation)
		locations.DELETE("/drop", api.DropLocations)
	}

	hostNetworks := apiRoutes.Group("/networks")
	{
		hostNetworks.GET("/schema", api.GetNetworkSchema)
		hostNetworks.GET("", api.GetHostNetworks)
		hostNetworks.POST("", api.CreateHostNetwork)
		hostNetworks.GET("/:uuid", api.GetHostNetwork)
		hostNetworks.PATCH("/:uuid", api.UpdateHostNetwork)
		hostNetworks.DELETE("/:uuid", api.DeleteHostNetwork)
		hostNetworks.DELETE("/drop", api.DropHostNetworks)
	}

	hosts := apiRoutes.Group("/hosts")
	{
		hosts.GET("/schema", api.GetHostSchema)
		hosts.GET("", api.GetHosts)
		hosts.POST("", api.CreateHost)
		hosts.GET("/:uuid", api.GetHost)
		hosts.PATCH("/:uuid", api.UpdateHost)
		hosts.DELETE("/:uuid", api.DeleteHost)
		hosts.DELETE("/drop", api.DropHosts)
		hosts.GET("update-status", api.UpdateStatus)
		hosts.GET("/:uuid/configure-openvpn", api.ConfigureOpenVPN)
	}

	teams := apiRoutes.Group("/teams")
	{
		teams.GET("/schema", api.TeamsSchema)
		teams.GET("", api.GetTeams)
		teams.POST("", api.CreateTeam)
		teams.GET("/:uuid", api.GetTeam)
		teams.PATCH("/:uuid", api.UpdateTeam)
		teams.DELETE("/:uuid", api.DeleteTeam)
		teams.DELETE("/drop", api.DropTeams)
	}

	Tasks := apiRoutes.Group("/tasks")
	{
		Tasks.GET("/schema", api.TasksSchema)
		Tasks.GET("", api.GetTasks)
		Tasks.POST("", api.CreateTask)
		Tasks.GET("/:uuid", api.GetTask)
		Tasks.PATCH("/:uuid", api.UpdateTask)
		Tasks.DELETE("/:uuid", api.DeleteTask)
		Tasks.DELETE("/drop", api.DropTasks)
	}

	messages := apiRoutes.Group("/transactions")
	{
		messages.GET("/schema", api.TransactionsSchema)
		messages.GET("", api.GetTransactions)
		messages.POST("", api.CreateTransaction)
		messages.GET("/:uuid", api.GetTransaction)
		messages.PATCH("/:uuid", api.UpdateTransaction)
		messages.DELETE("/:uuid", api.DeleteTransaction)
		messages.DELETE("/drop", api.DropTransactions)
	}

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
		files.GET("/exists", api.FileExists)            // needs to be a file
		files.GET("/walk", api.WalkFile)                // similar as find in linux command
		files.GET("/list", api.ListFiles)               // list all files and folders
		files.POST("/create", api.CreateFile)           // create file only
		files.POST("/copy", api.CopyFile)               // copy either file or folder
		files.POST("/rename", api.RenameFile)           // rename either file or folder
		files.POST("/move", api.MoveFile)               // move file only
		files.POST("/upload", api.UploadFile)           // upload single file
		files.POST("/download", api.DownloadFile)       // download single file
		files.GET("/read", api.ReadFile)                // read single file
		files.PUT("/write", api.WriteFile)              // write single file
		files.DELETE("/delete", api.DeleteFile)         // delete single file
		files.DELETE("/delete-all", api.DeleteAllFiles) // deletes file or folder
	}

	dirs := apiRoutes.Group("/dirs")
	{
		dirs.GET("/exists", api.DirExists)  // needs to be a folder
		dirs.POST("/create", api.CreateDir) // create folder
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
		token.GET("/:uuid", api.GetToken)
		token.POST("/generate", api.GenerateToken)
		token.PUT("/:uuid/block", api.BlockToken)
		token.PUT("/:uuid/regenerate", api.RegenerateToken)
		token.DELETE("/:uuid", api.DeleteToken)
	}

	return r
}
