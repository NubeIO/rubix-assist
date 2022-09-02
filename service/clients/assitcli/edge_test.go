package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/go-resty/resty/v2"
	"testing"
)

var appName = "flow-framework"
var buildName = "flow-framework"
var appVersion = "v0.6.1"

var client = New(&Client{
	Rest: &resty.Client{},
	URL:  "0.0.0.0",
	Port: 1662,
})

func TestClient_EdgeListApps(t *testing.T) {
	data, err := client.EdgeListApps("rc")
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeListAppsAndService(t *testing.T) {
	data, err := client.EdgeListAppsAndService("rc")
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeProductInfo(t *testing.T) {
	data, err := client.EdgeProductInfo("rc")
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeCtlAction(t *testing.T) {
	data, err := client.EdgeCtlAction("rc", &installer.CtlBody{
		AppName: "flow-framework",
		Service: "",
		Action:  "start",
	})
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeCtlStatus(t *testing.T) {
	data, err := client.EdgeCtlStatus("rc", &installer.CtlBody{
		AppName: "flow-framework",
		Service: "",
		Action:  "isInstalled",
	})
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeServiceMassStatus(t *testing.T) {
	data, err := client.EdgeServiceMassStatus("rc", &installer.CtlBody{
		AppNames: []string{"flow-framewor"},
		Service:  "",
		Action:   "isInstalled",
	})
	fmt.Println(err)
	pprint.PrintJSON(data)
}
