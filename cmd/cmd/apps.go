package cmd

import (
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/spf13/cobra"
)

var (
	getApps bool
	//reboot  bool //reboot the host
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "manage rubix service apps",
	Long:  `do things like install an app, the device must have internet access to download the apps`,
	Run:   runApps,
}

func runApps(cmd *cobra.Command, args []string) {
	host := initSession()
	if getArch {
		out := host.Uptime()
		pprint.PrintJOSN(out)
	}
	if reboot {
		out := host.HostReboot()
		pprint.PrintJOSN(out)
	}

}

func init() {
	RootCmd.AddCommand(appsCmd)
	appsCmd.Flags().BoolVarP(&getApps, "get", "", false, "reboot the host")

}
