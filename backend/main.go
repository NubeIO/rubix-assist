package main

import (
	"fmt"
)



func main() {
	fmt.Println("try and start")
	//if err := config.Setup(); err != nil {
	//	logger.Fatalf("config.Setup() error: %s", err)
	//}
	//
	//if err := database.Setup(); err != nil {
	//	logger.Fatalf("database.Setup() error: %s", err)
	//}
	//
	//db := database.GetDB()
	//r := router.Setup(db)
	//
	//host := "127.0.0.1"
	//if h := viper.GetString("server.host"); h != "" {
	//	host = h
	//}
	//logger.Infof("Server is starting at %s:%s", host, viper.GetString("server.port"))
	//fmt.Printf("server is running at %s:%s Check logs for details\n", host, viper.GetString("server.port"))
	//fmt.Println()
	//logger.Fatalf("%v", r.Run(host+":"+viper.GetString("server.port")))
}
