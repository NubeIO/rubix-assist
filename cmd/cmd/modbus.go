package cmd

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/drivers/modbus"
	"github.com/spf13/cobra"
)

var (
	modbusIp        string
	modbusPort      int
	modbusRegNumber int
	modbusRegCount  int
)

var modbusCmd = &cobra.Command{
	Use:   "modbus",
	Short: "modbus read and write",
	Long:  ``,
	Run:   runModbus,
}

func modbusInit() {
	mbClient := &modbus.Client{
		HostIP:   modbusIp,
		HostPort: modbusPort,
	}
	mbClient, err := modbus.SetClient(mbClient)
	if err != nil {
		fmt.Println("HERE")
		return
	}
	mbClient.TCPClientHandler.Address = fmt.Sprintf("%s:%d", modbusIp, modbusPort)
	mbClient.TCPClientHandler.SlaveID = byte(1)
	coils, err := mbClient.Client.ReadCoils(uint16(modbusRegNumber), uint16(modbusRegCount))
	if err != nil {
		fmt.Println("coils", err)
	}
	fmt.Println("coils", coils)
	return
}

func runModbus(cmd *cobra.Command, args []string) {
	modbusInit()
}

func init() {
	RootCmd.AddCommand(modbusCmd)
	modbusCmd.Flags().StringVarP(&modbusIp, "ip", "", "192.168.15.202", "device ip")
	modbusCmd.Flags().IntVarP(&modbusPort, "port", "", 502, "device port")
	modbusCmd.Flags().IntVarP(&modbusRegNumber, "register", "", 1, "read register")
	modbusCmd.Flags().IntVarP(&modbusRegCount, "count", "", 1, "register count")
}
