package cmd

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/database"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var (
	//model.Host
	hostName     string
	hostIP       string
	hostPort     int
	hostUsername string
	hostPassword string

	rubixPort     int
	rubixUsername string
	rubixPassword string

	iface string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nube-cli",
	Short: "description",
	Long:  `description`,
	Run:   runServer,
}

var flgRoot struct {
	startApp bool
}

func setup() {
	logger.Init()
	logger.SetLogLevel(logrus.InfoLevel)
	logger.InfoLn("try and start rubix-updater")
	if err := config.Setup(); err != nil {
		logger.Errorf("config.Setup() error: %s", err)
	}
	if err := database.Setup(); err != nil {
		logger.Fatalf("database.Setup() error: %s", err)
	}

}

func runServer(cmd *cobra.Command, args []string) {

	if flgRoot.startApp {
		setup()
		db := database.GetDB()
		r := router.Setup(db)
		r.MaxMultipartMemory = 250 << 20 //250 mb
		host := "0.0.0.0"
		if h := viper.GetString("server.host"); h != "" {
			host = h
		}
		logger.Infof("Server is starting at %s:%s", host, viper.GetString("server.port"))
		fmt.Printf("server is running at %s:%s Check logs for details\n", host, viper.GetString("server.port"))
		log.Fatalf("%v", r.Run(host+":"+viper.GetString("server.port")))
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&flgRoot.startApp, "server", "", false, "start rubix assist")
	RootCmd.PersistentFlags().StringVarP(&hostName, "host", "", "RC", "host name (default RC)")
	RootCmd.PersistentFlags().StringVarP(&hostIP, "ip", "", "192.168.15.10", "host ip (default 192.168.15.10)")
	RootCmd.PersistentFlags().IntVarP(&hostPort, "port", "", 22, "SSH Port")
	RootCmd.PersistentFlags().StringVarP(&iface, "iface", "", "", "pc or host network interface example: eth0")
	RootCmd.PersistentFlags().StringVarP(&hostUsername, "host-user", "", "pi", "host/linux username (default pi)")
	RootCmd.PersistentFlags().StringVarP(&hostPassword, "host-pass", "", "N00BRCRC", "host/linux password")
	RootCmd.PersistentFlags().IntVarP(&rubixPort, "rubix-port", "", 1616, "rubix port (default 1616)")
	RootCmd.PersistentFlags().StringVarP(&rubixUsername, "rubix-user", "", "admin", "rubix username (default admin)")
	RootCmd.PersistentFlags().StringVarP(&rubixPassword, "rubix-pass", "", "N00BWires", "rubix password")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
