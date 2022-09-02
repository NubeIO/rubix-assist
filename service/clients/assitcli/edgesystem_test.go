package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-edge/service/system"
	"testing"
)

func TestClient_EdgeGetNetworks(t *testing.T) {
	data, err := client.EdgeGetNetworks("rc")
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeDHCPPortExists(t *testing.T) {
	data, err := client.EdgeDHCPPortExists("rc", &system.NetworkingBody{PortName: "eth0"})
	fmt.Println(err)
	fmt.Println(data)
}
