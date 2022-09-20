package appstore

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestStore_listAppsWithVersions(t *testing.T) {
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir: "/data",
		},
	})
	app, err := appStore.ListAppsWithVersions()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(app)
}
