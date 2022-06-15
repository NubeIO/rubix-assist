package wirescli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("192.168.15.113", 1313)

	body := &WiresTokenBody{
		Username: "admin",
		Password: "aaaa",
	}
	d, r := client.GetToken(body)
	fmt.Println(r)
	pprint.PrintJOSN(d)
}
