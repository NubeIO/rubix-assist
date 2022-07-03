package ffclient

import (
	"fmt"
	"testing"
)

func TestFlowClient_GetNetworksWithPoints(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	var err error
	nets, err := cli.GetNetworkByPluginName("bacnetmaster", true)
	if err != nil {
		return
	}
	for _, device := range nets.Devices {
		fmt.Println(device)
	}
	fmt.Println(nets.UUID)

	network, err := cli.GetNetwork(nets.UUID, false)
	if err != nil {
		return
	}
	fmt.Println(network.Devices)
	for _, d := range network.Devices {
		fmt.Println(d)
	}

}
