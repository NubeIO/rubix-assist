package cmd

import (
	"github.com/NubeIO/rubix-assist/drivers/flow"
	"github.com/spf13/cobra"
)

var (
	flowScan bool
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "flow-network",
	Long:  ``,
	Run:   runFlow,
}

func runFlow(cmd *cobra.Command, args []string) {

	if flowScan {
		scan := &flow.Scan{Debug: true}
		scan.Scan()
	}

}

func init() {
	RootCmd.AddCommand(flowCmd)
	flowCmd.Flags().BoolVarP(&flowScan, "scan", "", false, "do a network scan for iot devices")

}
