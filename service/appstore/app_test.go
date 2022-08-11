package appstore

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

func TestStore_listAppsWithVersions(t *testing.T) {
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	app, err := appStore.ListAppsWithVersions()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(app)
}

func TestStore_listAppsWithArch(t *testing.T) {

	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	app, err := appStore.ListAppsBuildDetails()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(app)
}

func TestStore_ListAppsFlow(t *testing.T) {
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})

	path, err := appStore.ListApps()
	fmt.Println(err)
	pprint.PrintJOSN(path)
	app, err := appStore.ListAppVersions("flow-framework")
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(app)
	appDetails, err := appStore.ListAppArchTypes("flow-framework", "v0.6.0")
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(appDetails)

}

func TestStore_ListAppsWires(t *testing.T) {
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})

	path, err := appStore.ListApps()
	fmt.Println(err, "ListApps")
	pprint.PrintJOSN(path)
	app, err := appStore.listAppBuilds("rubix-wires", "v2.7.3")
	fmt.Println(err, "listAppsBuilds")
	if err != nil {
		return
	}
	pprint.PrintJOSN(app)
	app, err = appStore.ListAppVersions("rubix-wires")
	fmt.Println(err, "ListAppsVersions")
	if err != nil {
		return
	}
	pprint.PrintJOSN(app)
	appDetails, err := appStore.ListAppArchTypes("rubix-wires", "v2.7.3")
	fmt.Println(err, "ListAppsArchTypes")
	if err != nil {
		return
	}
	pprint.PrintJOSN(appDetails)

}
