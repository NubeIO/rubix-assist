package edgecli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-edge/service/system"
	"testing"
)

func TestClient_UWFStatusList(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	resp, err := cli.UWFStatusList()
	pprint.PrintJSON(resp)
	fmt.Println(err)
	if err != nil {
		return
	}
}

func TestClient_UWFDisable(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	resp, err := cli.UWFDisable()
	pprint.PrintJSON(resp)
	fmt.Println(err)
	if err != nil {
		return
	}
}

func TestClient_UWFEnable(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	resp, err := cli.UWFEnable()
	pprint.PrintJSON(resp)
	fmt.Println(err)
	if err != nil {
		return
	}
}

func TestClient_UWFOpenPort(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	resp, err := cli.UWFOpenPort(system.UFWBody{Port: 1884})
	pprint.PrintJSON(resp)
	fmt.Println(err)
	if err != nil {
		return
	}
}

func TestClient_UWFClosePort(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	resp, err := cli.UWFClosePort(system.UFWBody{Port: 1884})
	pprint.PrintJSON(resp)
	fmt.Println(err)
	if err != nil {
		return
	}
}
