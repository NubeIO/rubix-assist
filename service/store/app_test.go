package store

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestStore_addApp(t *testing.T) {
	appName := "rubix-wires"
	appVersion := "v2.7.2"

	app, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	err = app.AddApp(appName, appVersion)
	if err != nil {
		return
	}
	apps, err := app.ListApps()
	if err != nil {
		return
	}
	pprint.PrintJOSN(apps)

}
