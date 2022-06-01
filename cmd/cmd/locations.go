package cmd

import (
	"fmt"
	"github.com/NubeIO/rubix-assist-model/model"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "location",
	Short: "add a new location",
	Long:  ``,
	Run:   runApps,
}

func runApps(cmd *cobra.Command, args []string) {
	db := initDB()
	loc := &model.Location{
		Name: "",
	}

	location, err := db.CreateLocation(loc)
	if err != nil {
		return
	}
	pprint.PrintJOSN(location)

	network := &model.Network{
		Name:         "",
		LocationUUID: location.UUID,
	}
	network, err = db.CreateHostNetwork(network)
	if err != nil {
		return
	}
	pprint.PrintJOSN(network)
	host := &model.Host{
		Name:        "",
		NetworkUUID: network.UUID,
	}

	host, err = db.CreateHost(host)
	fmt.Println(err, 999)
	if err != nil {
		return
	}
	pprint.PrintJOSN(host)

}

var flgLocation struct {
	name string
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flagSet := appsCmd.Flags()
	flagSet.StringVar(&flgLocation.name, "name", "", "name of the location")

}
