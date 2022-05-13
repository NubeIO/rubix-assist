package cmd

import (
	"fmt"
	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/database"
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

func read(cmd *cobra.Command, args []string) {

	if err := config.Setup(); err != nil {
		log.Errorln("config.Setup() error: %s", err)
	}
	if err := database.Setup(); err != nil {
		log.Errorln("database.Setup() error: %s", err)
	}
	gormDB := database.GetDB()
	appDB := &dbase.DB{
		DB: gormDB,
	}

	if hostsGet {
		hosts, err := appDB.GetHosts()
		fmt.Println(err)
		for i, h := range hosts {
			fmt.Println(i, h.IP)
		}
	}
	if hostAdd {
		host, err := appDB.CreateHost(&model.Host{})
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
