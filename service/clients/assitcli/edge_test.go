package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

var appName = "flow-framework"
var buildName = "flow-framework"
var appVersion = "v0.6.1"

func TestClient_EdgeListApps(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeListApps("rc")
	fmt.Println(err)
	pprint.PrintJOSN(data)
}

func TestClient_EdgeListAppsAndService(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeListAppsAndService("rc")
	fmt.Println(err)
	pprint.PrintJOSN(data)
}
