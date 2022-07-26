package store

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func Test_read(t *testing.T) {

	appName := "flow-framework"
	appVersion := "v0.6.0"

	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})

	details, err := getFileDetails("/data/store/apps/flow-framework/v0.6.0/")

	pprint.PrintJOSN(details)

	buildCheck, err := appStore.checkApp(appName, appVersion, "nubeio-flow-framework.service")
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(buildCheck)

}
