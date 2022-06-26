package wirescli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {

	client := New("0.0.0.0", 1313)

	body := &WiresTokenBody{
		Username: "admin",
		Password: "admin",
	}
	token, r := client.GetToken(body)
	fmt.Println(r)
	pprint.PrintJOSN(token)
	ok, res := client.Backup(token.Token)

	fmt.Println(ok, res)
}
