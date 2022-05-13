package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ip string
var modbusIp string
var modbusPort int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nube-cli",
	Short: "description",
	Long:  `description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&modbusIp, "modbus-ip", "", "192.168.15.93", "host ip")
	RootCmd.PersistentFlags().IntVarP(&modbusPort, "modbus-port", "", 502, "Port")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
