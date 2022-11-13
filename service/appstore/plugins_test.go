package appstore

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func Test_ListPlugins(t *testing.T) {
	appStore, err := New(&Store{})
	details, err := appStore.GetPluginsStorePlugins()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(details)
}

func Test_getPluginPath(t *testing.T) {
	appStore, err := New(&Store{})
	pluginPath, err := appStore.GetPluginsStorePluginFile(&Plugin{
		Name:    "bacnetserver",
		Arch:    "amd64",
		Version: "v0.6.6",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(pluginPath)
}
