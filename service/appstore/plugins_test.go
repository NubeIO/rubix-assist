package appstore

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func Test_pluginDetails(t *testing.T) {

	//fmt.Println(PluginDetails("influx-amd64.so"))
}
func Test_PluginZipDetails(t *testing.T) {
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	details := appStore.PluginZipDetails("bacnetserver-0.6.6-a79d6c29.armv7.zip")
	if err != nil {
		return
	}
	pprint.PrintJOSN(details)
}

func Test_ListPlugins(t *testing.T) {
	appStore, err := New(&Store{
		App:  &installer.App{},
		Perm: nonRoot,
	})
	details, err := appStore.StoreListPlugins()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(details)
}

func Test_getPluginPath(t *testing.T) {
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	details, name, err := appStore.GetPluginPath(&Plugin{
		PluginName: "bacnetserver",
		Arch:       "amd64",
		Version:    "v0.6.6",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(details)
	fmt.Println(name)
}
