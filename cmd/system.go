package cmd

import (
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/remote"
	"github.com/NubeIO/rubix-assist/service/remote/ssh"
	"github.com/spf13/cobra"
)

var (
	getArch bool
	reboot  bool //reboot the host
)

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "system admin",
	Long:  `pass in the host name and do operation like check arch type of the host`,
	Run:   runSystem,
}

func initSession() *remote.Admin {

	db := initDB()

	host, name := db.GetHostByName(hostName)
	if name != nil {
		return nil
	}
	session := &remote.Admin{
		SSH: &ssh.Host{
			Host: &model.Host{
				IP:       host.IP,
				Port:     host.Port,
				Username: host.Username,
				Password: host.Password,
			},
		},
	}
	return remote.New(session)

}

func runSystem(cmd *cobra.Command, args []string) {
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
	RootCmd.AddCommand(systemCmd)
	systemCmd.Flags().BoolVarP(&reboot, "reboot", "", false, "reboot the host")
	systemCmd.Flags().BoolVarP(&getArch, "arch", "", false, "get arch type of host")

}
