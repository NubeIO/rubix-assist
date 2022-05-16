package cmd

import (
	"github.com/NubeIO/rubix-assist/drivers/flow"
	"github.com/spf13/cobra"
)

var (
	flowScan bool
	ipRange  string
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "flow-network",
	Long:  ``,
	Run:   runFlow,
}

func runFlow(cmd *cobra.Command, args []string) {
	if flowScan {
		scan := &flow.Scan{IP: ipRange, Iface: iface, Debug: true}
		scan.Scan()
	}

}

func init() {
	RootCmd.AddCommand(flowCmd)
	flowCmd.Flags().BoolVarP(&flowScan, "scan", "", false, "do a network scan for iot devices")
	flowCmd.Flags().StringVarP(&ipRange, "ip-range", "", "192.168.15.1-254", "scan range ip example is 192.168.15.1-10 would scan the 10 ips")

}
