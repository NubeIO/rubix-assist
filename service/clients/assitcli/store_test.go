package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_ListStore(t *testing.T) {
	client := New("0.0.0.0", 1662)
	store, err := client.ListStore()
	fmt.Println(err)
	if err != nil {
		return
	}

	pprint.PrintJOSN(store)

}
