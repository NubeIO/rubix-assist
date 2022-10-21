package assistcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_GetLocalStorage(t *testing.T) {
	https := false
	cli := New(&Client{
		Rest:  nil,
		Ip:    "",
		Port:  0,
		HTTPS: &https,
	})

	storage, err := cli.GetLocalStorage("rc")
	if err != nil {
		return
	}
	pprint.PrintJSON(storage)
}

func TestClient_UpdateLocalStorage(t *testing.T) {
	https := false
	cli := New(&Client{
		Rest:  nil,
		Ip:    "",
		Port:  0,
		HTTPS: &https,
	})

	storage, err := cli.UpdateLocalStorage("rc", &model.LocalStorageFlowNetwork{
		FlowIP:       "192.168.15.68",
		FlowPort:     1660,
		FlowHTTPS:    nil,
		FlowUsername: "",
		FlowPassword: "",
		FlowToken:    "",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(storage)
}
