package cmd

import (
	"github.com/spf13/cobra"
)

var (
	hostName     string
	hostIP       string
	hostPort     int
	hostUsername string
	hostPassword string

	rubixPort     int
	rubixUsername string
	rubixPassword string

	iface string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nube-cli",
	Short: "description",
	Long:  `description`,
}

var flgRoot struct {
	startApp  bool
	prod      bool
	auth      bool
	port      int
	rootDir   string
	appDir    string
	dataDir   string
	configDir string
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&flgRoot.prod, "prod", "", false, "prod")
	RootCmd.PersistentFlags().BoolVarP(&flgRoot.prod, "auth", "", true, "auth")
	RootCmd.PersistentFlags().IntVarP(&flgRoot.port, "port", "p", 1662, "port (default 1662)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.rootDir, "root-dir", "r", "./", "root dir") // in production it will be `/data`
	RootCmd.PersistentFlags().StringVarP(&flgRoot.appDir, "app-dir", "a", "./", "app dir")    // in production it will be `/rubix-assist`
	RootCmd.PersistentFlags().StringVarP(&flgRoot.dataDir, "data-dir", "d", "data", "data dir")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.configDir, "config-dir", "c", "config", "config dir")
	RootCmd.PersistentFlags().StringVarP(&hostName, "host", "", "rc", "host name (default rc)")
	RootCmd.PersistentFlags().StringVarP(&hostIP, "ip", "", "192.168.15.10", "host ip (default 192.168.15.10)")
	RootCmd.PersistentFlags().IntVarP(&hostPort, "ssh-port", "", 22, "SSH Port")
	RootCmd.PersistentFlags().StringVarP(&iface, "iface", "", "", "pc or host network interface example: eth0")
	RootCmd.PersistentFlags().StringVarP(&hostUsername, "host-user", "", "pi", "host/linux username (default pi)")
	RootCmd.PersistentFlags().StringVarP(&hostPassword, "host-pass", "", "N00BRCRC", "host/linux password")
	RootCmd.PersistentFlags().IntVarP(&rubixPort, "rubix-port", "", 1616, "rubix port (default 1616)")
	RootCmd.PersistentFlags().StringVarP(&rubixUsername, "rubix-user", "", "admin", "rubix username (default admin)")
	RootCmd.PersistentFlags().StringVarP(&rubixPassword, "rubix-pass", "", "N00BWires", "rubix password")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
