package cmd

import (
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/pkg/model"
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
		Name: flgLocation.name,
	}

	location, err := db.CreateLocation(loc)
	if err != nil {
		return
	}

	network := &model.Network{
		Name:         "",
		LocationUUID: location.UUID,
	}
	network, err = db.CreateHostNetwork(network)
	if err != nil {
		return
	}
	host := &model.Host{
		Name:        "",
		NetworkUUID: network.UUID,
	}

	host, err = db.CreateHost(host)
	if err != nil {
		return
	}

	locations, err := db.GetLocations()
	if err != nil {
		return
	}
	pprint.PrintJOSN(locations)

}

var flgLocation struct {
	name string
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flagSet := appsCmd.Flags()
	flagSet.StringVar(&flgLocation.name, "name", "", "name of the location")

}
