package assistcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/go-resty/resty/v2"
	"testing"
)

var client = New(&Client{
	Rest: &resty.Client{},
	Ip:   "0.0.0.0",
	Port: 1662,
})

func TestClient_EdgeProductInfo(t *testing.T) {
	data, err := client.EdgeProductInfo("rc")
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeCtlAction(t *testing.T) {
	data, err := client.EdgeSystemCtlAction("rc", &installer.SystemCtlBody{
		AppName: "flow-framework",
		Action:  "start",
	})
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeCtlStatus(t *testing.T) {
	data, err := client.EdgeSystemCtlStatus("rc", &installer.SystemCtlBody{
		AppName: "flow-framework",
		Action:  "isInstalled",
	})
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeServiceMassStatus(t *testing.T) {
	data, err := client.EdgeServiceMassStatus("rc", &installer.SystemCtlBody{
		AppNames: []string{"flow-framewor"},
		Action:   "isInstalled",
	})
	fmt.Println(err)
	pprint.PrintJSON(data)
}
