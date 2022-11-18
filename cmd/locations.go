package cmd

import (
	model2 "github.com/NubeIO/rubix-assist/model"
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
	loc := &model2.Location{
		Name: flgLocation.name,
	}

	location, err := db.CreateLocation(loc)
	if err != nil {
		return
	}

	network := &model2.Network{
		Name:         "",
		LocationUUID: location.UUID,
	}
	network, err = db.CreateHostNetwork(network)
	if err != nil {
		return
	}
	host := &model2.Host{
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
	pprint.PrintJSON(locations)
}

var flgLocation struct {
	name string
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flagSet := appsCmd.Flags()
	flagSet.StringVar(&flgLocation.name, "name", "", "name of the location")
}
