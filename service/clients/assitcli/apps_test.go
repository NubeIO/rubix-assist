package assitcli

import (
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("0.0.0.0", 1662)
	d, _ := client.GetLocationSchema()

	pprint.PrintJOSN(d)

}
