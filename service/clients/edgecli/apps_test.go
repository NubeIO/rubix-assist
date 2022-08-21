package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

var appName = "flow-framework"
var appVersion = "v0.6.1"
var fileName = "flow-framework-0.6.1-6cfec278.amd64.zip"
var source = "/data/tmp/tmp_E57DA9ED2A7B/flow-framework-0.6.1-6cfec278.amd64.zip"

func Test_EdgeProductInfo(*testing.T) {
	cli := New(&Client{})
	apps, err := cli.EdgeProductInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)

}

func Test_ListApps(*testing.T) {
	cli := New(&Client{})
	apps, err := cli.ListApps()
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)

}

func Test_ListAppsAndService(*testing.T) {
	cli := New(&Client{})
	apps, err := cli.ListAppsAndService()
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)

}

func Test_ListNubeServices(*testing.T) {
	cli := New(&Client{})
	apps, err := cli.ListNubeServices()
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)

}

func Test_InstallApp(*testing.T) {

	cli := New(&Client{})
	file, err := cli.InstallApp(&installer.Install{
		Name: appName,

		Version: appVersion,
		Source:  source,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(file)

}
