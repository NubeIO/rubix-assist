package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"testing"
)

func TestClient_EdgeUploadPlugin(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeUploadPlugin("rc", &appstore.Plugin{
		PluginName: "bacnetserver",
		Arch:       "amd64",
		Version:    "v0.6.6",
	})
	fmt.Println(err)
	pprint.PrintJOSN(data)
}

func TestClient_EdgeListPlugins(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeListPlugins("rc")
	fmt.Println(err)
	pprint.PrintJOSN(data)
}

func TestClient_EdgeDeletePlugin(t *testing.T) {
	client := New("0.0.0.0", 1662)
	data, err := client.EdgeDeletePlugin("rc", &appstore.Plugin{
		PluginName: "bacnetserver",
		Arch:       "amd64",
	})
	fmt.Println(err)
	pprint.PrintJOSN(data)
}
