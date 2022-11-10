package appstore

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestStore_listAppsWithVersions(t *testing.T) {
	appStore, err := New(&Store{})
	app, err := appStore.ListAppsWithVersions()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(app)
}
