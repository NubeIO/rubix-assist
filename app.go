package main

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/database"
	"github.com/NubeIO/rubix-assist/pkg/jobs"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/router"
	"github.com/NubeIO/rubix-assist/service/ping"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

func setup() {
	logger.Init()
	logger.SetLogLevel(logrus.InfoLevel)
	logger.InfoLn("try and start rubix-updater")
	if err := config.Setup(); err != nil {
		logger.Fatalf("config.Setup() error: %s", err)
	}
	if err := database.Setup(); err != nil {
		logger.Fatalf("database.Setup() error: %s", err)
	}
	j := new(jobs.Jobs)
	j.InitCron()

}

func main() {
	setup()
	db := database.GetDB()
	r := router.Setup(db)
	ping.TEST()

	host := "0.0.0.0"
	if h := viper.GetString("server.host"); h != "" {
		host = h
	}
	logger.Infof("Server is starting at %s:%s", host, viper.GetString("server.port"))
	fmt.Printf("server is running at %s:%s Check logs for details\n", host, viper.GetString("server.port"))
	log.Fatalf("%v", r.Run(host+":"+viper.GetString("server.port")))
}
