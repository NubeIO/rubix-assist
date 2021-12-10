package main

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/pkg/config"
	"github.com/NubeIO/rubix-updater/pkg/database"
	"github.com/NubeIO/rubix-updater/pkg/logger"
	"github.com/NubeIO/rubix-updater/pkg/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.Println("try and start rubix-updater")
	if err := config.Setup(); err != nil {
		logger.Fatalf("config.Setup() error: %s", err)
	}

	if err := database.Setup(); err != nil {
		logger.Fatalf("database.Setup() error: %s", err)
	}

	db := database.GetDB()
	r := router.Setup(db)

	host := "0.0.0.0"
	if h := viper.GetString("server.host"); h != "" {
		host = h
	}
	logger.Infof("Server is starting at %s:%s", host, viper.GetString("server.port"))
	fmt.Printf("server is running at %s:%s Check logs for details\n", host, viper.GetString("server.port"))
	fmt.Println()
	logger.Fatalf("%v", r.Run(host+":"+viper.GetString("server.port")))
}
