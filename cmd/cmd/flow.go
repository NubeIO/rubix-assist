package cmd

import (
	"github.com/NubeIO/rubix-assist/drivers/flow"
	"github.com/spf13/cobra"
)

var (
	flowScan      bool
	ip            string
	interfaceName string
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "flow-network",
	Long:  ``,
	Run:   runFlow,
}

func runFlow(cmd *cobra.Command, args []string) {
	if flowScan {
		scan := &flow.Scan{}
		scan.Scan(ip, 254, interfaceName)
	}

}

func init() {
	RootCmd.AddCommand(flowCmd)
	flowCmd.Flags().BoolVarP(&flowScan, "scan", "", false, "do a network scan for iot devices")
	flowCmd.Flags().StringVarP(&ip, "ip", "", "192.168.15.1", "scan range ip example is 192.168.15.1-10 would scan the 10 ips")
	flowCmd.Flags().StringVarP(&interfaceName, "interface", "", "", "host interface")

}
