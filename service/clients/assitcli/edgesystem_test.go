package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_EdgeGetNetworks(t *testing.T) {
	data, err := client.EdgeGetNetworks("rc")
	fmt.Println(err)
	pprint.PrintJOSN(data)
}
