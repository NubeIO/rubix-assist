package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/store"
	"testing"
)

func TestClient_AddApp(t *testing.T) {

	client := New("0.0.0.0", 1662)

	newApp, err := client.AddApp(&store.App{
		Name:        "flow-framework",
		Version:     "v0.6.1",
		ServiceFile: "",
	})
	fmt.Println(err)
	pprint.PrintJOSN(newApp)
	store, err := client.ListStore()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(store)

}
