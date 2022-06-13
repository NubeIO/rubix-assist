package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("14.92.222.8", 1662)

	//body := &edgeapi.App{
	//	HostName: "rc",
	//}
	d, r := client.GetHostSchema()
	fmt.Println(r.StatusCode)
	pprint.PrintJOSN(d)
}
