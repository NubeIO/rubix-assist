package ffclient

import (
	"fmt"
	"testing"
)

func TestFlowClient_GetNetworksWithPoints(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	var err error
	nets, err := cli.GetFirstNetwork(true)
	if err != nil {
		return
	}
	for _, device := range nets.Devices {
		fmt.Println(device)
	}
	fmt.Println(nets)
	//for _, network := range *nets {
	//	fmt.Println(network)
	//}

}
