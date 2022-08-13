package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/appstore"
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

func TestClient_EdgeProductInfo(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeProductInfo("rc")
	fmt.Println(err)
	pprint.PrintJOSN(data)
}

func TestClient_EdgeCtlAction(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeCtlAction("rc", &installer.CtlBody{
		AppName: "flow-framework",
		Service: "",
		Action:  "start",
	})
	fmt.Println(err)
	pprint.PrintJOSN(data)
}

func TestClient_EdgeCtlStatus(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeCtlStatus("rc", &installer.CtlBody{
		AppName: "flow-framework",
		Service: "",
		Action:  "isInstalled",
	})
	fmt.Println(err)
	pprint.PrintJOSN(data)
}

func TestClient_EdgeUploadPlugin(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeUploadPlugin("rc", &appstore.Plugin{
		PluginName: "bacnetserver",
		Arch:       "amd64",
		Version:    "v0.6.6",
	})
	fmt.Println(err)
	pprint.PrintJOSN(data)
}
