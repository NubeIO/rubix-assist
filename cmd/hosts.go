package cmd

import (
	"fmt"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"

	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/database"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// model.Host
	hostsGet       bool
	hostAdd        bool
	hostsDrop      bool
	hostUpdateIp   bool
	hostUpdatePort bool
)

var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "hosts get/add and delete",
	Long:  ``,
	Run:   runHosts,
}

func initDB() *dbase.DB {
	if err := config.Setup(RootCmd); err != nil {
		log.Errorln("config.Setup() error: %s", err)
	}
	if err := database.Setup(); err != nil {
		log.Errorln("database.Setup() error: %s", err)
	}
	db := database.GetDB()
	appDB := &dbase.DB{
		DB: db,
	}

	return appDB
}

func getHosts(appDB *dbase.DB) {
	hosts, err := appDB.GetHosts()
	if err == nil {
		pprint.PrintJSON(hosts)
	} else {
		fmt.Println(err)
	}
}

func addHost(appDB *dbase.DB) {
	h := &model.Host{
		Name:     hostName,
		IP:       hostIP,
		Port:     hostPort,
		Username: hostUsername,
		Password: hostPassword,
	}
	host, err := appDB.CreateHost(h)
	if err == nil {
		pprint.PrintJSON(host)
	} else {
		fmt.Println(err)
	}
}

func updateHost(appDB *dbase.DB) {
	h := &model.Host{}
	if hostUpdateIp {
		h = &model.Host{
			IP: hostIP,
		}
	}
	if hostUpdatePort {
		h = &model.Host{
			Port: hostPort,
		}
	}
	host, err := appDB.UpdateHostByName(hostName, h)
	if err == nil {
		pprint.PrintJSON(host)
	} else {
		fmt.Println(err)
	}
}

func dropHosts(appDB *dbase.DB) {
	host, err := appDB.DropHosts()
	if err == nil {
		pprint.PrintJSON(host)
	} else {
		fmt.Println(err)
	}
}

func runHosts(cmd *cobra.Command, args []string) {
	appDB := initDB()

	if hostsGet {
		getHosts(appDB)
	}
	if hostAdd {
		addHost(appDB)
	}

	if hostsDrop {
		dropHosts(appDB)
	}

	if hostUpdateIp {
		updateHost(appDB)
	}

	if hostUpdatePort {
		updateHost(appDB)
	}
}

func init() {
	RootCmd.AddCommand(hostsCmd)
	hostsCmd.Flags().BoolVarP(&hostsGet, "all", "", false, "get all")
	hostsCmd.Flags().BoolVarP(&hostAdd, "new", "", false, "add a new")
	hostsCmd.Flags().BoolVarP(&hostsDrop, "drop", "", false, "delete all")
	hostsCmd.Flags().BoolVarP(&hostUpdateIp, "update-ip", "", false, "update host ip: hosts --update-ip=true --name=RC --ip=xx.xx.xx")
	hostsCmd.Flags().BoolVarP(&hostUpdatePort, "update-port", "", false, "update host port: hosts --update-port=true --name=RC --port=2022")
}
