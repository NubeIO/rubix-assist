package assitcli

import (
	"fmt"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("0.0.0.0", 1662)
	d, err := client.ProxyGET("rc", "/api/networking/networks")
	fmt.Println(err)
	fmt.Println(d.String())

}
