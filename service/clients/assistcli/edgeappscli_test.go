package assistcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_EdgeListApps(t *testing.T) {
	data, err := client.EdgeListApps("rc")
	fmt.Println(err)
	pprint.PrintJSON(data)
}
