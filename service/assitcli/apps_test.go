package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/edgeapi"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("0.0.0.0", 1662)

	body := &edgeapi.App{
		HostName: "rc",
	}
	_, r := client.InstallApp(body)
	fmt.Println(r)
}
