package router

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/controller"
	dbase "github.com/NubeIO/rubix-assist/database"
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

	r.POST("/api/users/login", api.Login)

	handleAuth := func(c *gin.Context) { c.Next() }
	if config.Config.Auth() {
		//handleAuth = api.HandleAuth() // TODO add back in auth
	}
	admin := r.Group("/api", handleAuth)

	appStore := admin.Group("/store/apps")
	{
		appStore.GET("/", api.ListAppsWithVersions)
		appStore.GET("/details", api.ListAppsBuildDetails)
		appStore.POST("/", api.AddUploadStoreApp)
	}

	storePlugins := admin.Group("/store/plugins")
	{
		storePlugins.GET("/", api.StoreListPlugins)
		storePlugins.POST("/", api.StoreUploadPlugin)
	}

	edge := admin.Group("/edge")
	{
		edge.GET("/system/product", api.EdgeProductInfo)
	}

	edgeApps := admin.Group("/edge/apps")
	{
		edgeApps.GET("/", api.EdgeListApps)
		edgeApps.GET("/services", api.EdgeListAppsAndService)
		edgeApps.GET("/services/nube", api.EdgeListNubeServices)
		edgeApps.POST("/add", api.AddUploadEdgeApp)
		edgeApps.POST("/service/upload", api.GenerateUploadEdgeService)
		edgeApps.POST("/service/install", api.InstallEdgeService)
		edgeApps.DELETE("/", api.EdgeUninstallApp)
	}

	edgeConfig := admin.Group("/edge/config")
	{
		edgeConfig.POST("/", api.EdgeReplaceConfig)
	}

	edgePlugins := admin.Group("/edge/plugins")
	{
		edgePlugins.GET("/", api.EdgeListPlugins)
		edgePlugins.POST("/", api.EdgeUploadPlugin)
		edgePlugins.DELETE("/", api.EdgeDeletePlugin)
		edgePlugins.DELETE("/all", api.EdgeDeleteAllPlugins)
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
	{
		hosts.GET("/schema", api.GetHostSchema)
		hosts.GET("/", api.GetHosts)
		hosts.POST("/", api.CreateHost)
		hosts.GET("/:uuid", api.GetHost)
		hosts.PATCH("/:uuid", api.UpdateHost)
		hosts.DELETE("/:uuid", api.DeleteHost)
		hosts.DELETE("/drop", api.DropHosts)
	}

	teams := admin.Group("/teams")
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
	{
		messages.GET("/schema", api.TransactionsSchema)
		messages.GET("/", api.GetTransactions)
		messages.POST("/", api.CreateTransaction)
		messages.GET("/:uuid", api.GetTransaction)
		messages.PATCH("/:uuid", api.UpdateTransaction)
		messages.DELETE("/:uuid", api.DeleteTransaction)
		messages.DELETE("/drop", api.DropTransactions)
	}

	// tools := admin.Group("/tools")
	// //tools.Use(authMiddleware.MiddlewareFunc())
	// {
	//
	//	tools.GET("/edgeapi/ip/schema", api.EdgeIPSchema)
	//	tools.POST("/edgeapi/ip", api.EdgeSetIP)
	//	tools.POST("/edgeapi/ip/dhcp", api.EdgeSetIP)
	//
	// }

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

	files := admin.Group("/files")
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

	dirs := admin.Group("/dirs")
	{
		dirs.POST("/create", api.CreateDir)
		dirs.POST("/copy", api.CopyDir)
		dirs.DELETE("/delete", api.DeleteDir)
	}

	zip := admin.Group("/zip")
	{
		zip.POST("/unzip", api.Unzip)
		zip.POST("/zip", api.ZipDir)
	}

	user := admin.Group("/users")
	{
		user.PUT("", api.UpdateUser)
		user.GET("", api.GetUser)
	}

	token := admin.Group("/tokens")
	{
		token.GET("", api.GetTokens)
		token.POST("/generate", api.GenerateToken)
		token.PUT("/:uuid/block", api.BlockToken)
		token.PUT("/:uuid/regenerate", api.RegenerateToken)
		token.DELETE("/:uuid", api.DeleteToken)
	}

	r.Any("/proxy/*proxyPath", api.Proxy)

	return r
}
