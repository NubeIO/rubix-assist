package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/edge"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("0.0.0.0", 1662)

	body := &edge.App{
		HostName: "rc",
	}
	_, r := client.InstallApp(body)
	fmt.Println(r)
}
