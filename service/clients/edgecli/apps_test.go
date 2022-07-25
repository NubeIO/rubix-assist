package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"os"
	"testing"
)

var appName = "flow-framework"
var buildName = "flow-framework"
var appVersion = "v0.6.1"
var fileName = "flow-framework-0.6.1-6cfec278.amd64.zip"
var source = "/data/tmp/tmp_E57DA9ED2A7B/flow-framework-0.6.1-6cfec278.amd64.zip"

func Test_UploadApp(*testing.T) {

	cli := New("", 0)
	reader, err := os.Open(fmt.Sprintf("/data/store/apps/%s/%s/%s", appName, appVersion, fileName))
	fmt.Println(err)
	if err != nil {
		return
	}

	file, err := cli.UploadApp(appName, appVersion, buildName, fileName, reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(file)

}

func Test_InstallApp(*testing.T) {

	cli := New("", 0)
	file, err := cli.InstallApp(&installer.Install{
		Name:      appName,
		BuildName: buildName,
		Version:   appVersion,
		Source:    source,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(file)

}
