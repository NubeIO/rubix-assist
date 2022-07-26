package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/store"
	"testing"
)

var appName = "flow-framework"
var buildName = "flow-framework"
var appVersion = "v0.6.1"

func TestClient_UploadEdgeApp(t *testing.T) {

	client := New("0.0.0.0", 1662)

	listStore, err := client.ListStore()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(listStore)

	app, err := client.UploadEdgeApp("rc", &store.EdgeApp{

		Name:      appName,
		BuildName: buildName,
		Version:   appVersion,
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(app)

}

func TestClient_InstallEdgeApp(t *testing.T) {

	client := New("0.0.0.0", 1662)

	listStore, err := client.ListStore()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(listStore)

	app, err := client.InstallEdgeApp("rc", &installer.Install{
		Name:      appName,
		BuildName: buildName,
		Version:   appVersion,
		Source:    "/data/tmp/tmp_891F067EDA1E/flow-framework-0.6.1-6cfec278.amd64.zip",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(app)

}
