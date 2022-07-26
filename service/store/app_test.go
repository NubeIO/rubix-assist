package store

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestStore_addApp(t *testing.T) {
	appName := "rubix-wires"
	appVersion := "v2.7.4"

	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	app, err := appStore.AddApp(&App{Name: appName, Version: appVersion})
	fmt.Println(err)
	if err != nil {
		return
	}

	pprint.PrintJOSN(app)

}
