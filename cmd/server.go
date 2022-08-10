package cmd

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/database"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  "",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	setup()
	db := database.GetDB()
	r := router.Setup(db)
	r.MaxMultipartMemory = 250 << 20 // 250 mb

	host := "0.0.0.0"
	port := config.Config.GetPort()
	logger.Infof("Server is starting at %s:%s", host, port)
	fmt.Printf("server is running at %s:%s Check logs for details\n", host, port)
	log.Fatalf("%v", r.Run(host+":"+port))
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func setup() {
	logger.Init()
	logger.SetLogLevel(logrus.InfoLevel)
	logger.InfoLn("try and start rubix-updater")
	if err := config.Setup(RootCmd); err != nil {
		logger.Errorf("config.Setup() error: %s", err)
	}
	if err := os.MkdirAll(config.Config.GetAbsDataDir(), 0755); err != nil {
		panic(err)
	}
	if err := database.Setup(); err != nil {
		logger.Fatalf("database.Setup() error: %s", err)
	}
}
