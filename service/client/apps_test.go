package client

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/em"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("0.0.0.0", 1662)

	body := &em.App{
		HostName: "rc",
	}
	_, r := client.InstallApp(body)
	fmt.Println(r)
}
