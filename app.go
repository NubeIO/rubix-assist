package main

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/database"
	"github.com/NubeIO/rubix-assist/pkg/jobs"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/router"
	"github.com/NubeIO/rubix-assist/service/ping"
	log "github.com/sirupsen/logrus"
	"os"
)

func setup() {
	log.Println("try and start rubix-assist")
	conf := config.Setup()
	if err := os.MkdirAll(conf.GetAbsDataDir(), 0755); err != nil {
		panic(err)
	}
	logger.Setup()
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

	conf := config.GetConfig()
	logger.Infof("Server is starting at %s:%s", conf.Server.ListenAddr, conf.Server.Port)
	fmt.Printf("server is running at %s:%s Check logs for details\n", conf.Server.ListenAddr, conf.Server.Port)
	logger.Fatalf("%v", r.Run(conf.Server.ListenAddr+":"+conf.Server.Port))
}
