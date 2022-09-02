package wirescli

import (
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_InstallApp(t *testing.T) {
	client := New("192.168.15.171", 1313)

	body := &WiresTokenBody{
		Username: "admin",
		Password: "N00BWire",
	}
	token, _ := client.GetToken(body)
	pprint.PrintJSON(token)
	// data, _ := client.Backup(token.Token)
	// pprint.PrintJSON(data.Objects)
}
