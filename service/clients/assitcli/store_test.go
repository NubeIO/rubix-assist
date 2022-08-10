package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_ListAppsWithVersions(t *testing.T) {
	client := New("0.0.0.0", 1662)
	store, err := client.ListAppsWithVersions()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(store)

}

func TestClient_ListAppsBuildDetails(t *testing.T) {
	client := New("0.0.0.0", 1662)
	store, err := client.ListAppsBuildDetails()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(store)

}
