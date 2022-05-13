package cmd

import (
	"fmt"
	dbase "github.com/NubeIO/rubix-updater/database"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/config"
	"github.com/NubeIO/rubix-updater/pkg/database"
	dbhandler "github.com/NubeIO/rubix-updater/pkg/handler"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	hostsGet bool
	hostAdd  bool
)

// whoIsCmd represents the whoIs command
var readCmd = &cobra.Command{
	Use:   "hosts",
	Short: "BACnet device discovery",
	Long: `whoIs does a bacnet network discovery to find devices in the network
 given the provided range.`,
	Run: read,
}

var DB *dbase.DB

func read(cmd *cobra.Command, args []string) {

	if err := config.Setup(); err != nil {
		log.Errorln("config.Setup() error: %s", err)
	}
	if err := database.Setup(); err != nil {
		log.Errorln("database.Setup() error: %s", err)
	}
	db := database.GetDB()
	GDB := new(dbase.DB)
	GDB.DB = db
	gg := new(dbhandler.Handler)
	gg.DB = GDB

	if hostsGet {
		hosts, err := gg.DB.GetHosts()
		fmt.Println(err)
		for i, h := range hosts {
			fmt.Println(i, h.IP)
		}
	}
	if hostAdd {
		host, err := gg.DB.CreateHost(&model.Host{})
		fmt.Println(err)
		fmt.Println("NEW HOST", host.UUID)

	}

	fmt.Println("READ")

}

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().BoolVarP(&hostsGet, "all", "", false, "get all")
	readCmd.Flags().BoolVarP(&hostAdd, "add", "", false, "add a new")

}
