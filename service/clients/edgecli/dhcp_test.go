package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	"github.com/NubeIO/rubix-edge/service/system"
	"testing"
)

var deviceIP = "192.168.15.191"
var deviceIface = "eth0"

func TestClient_DHCPPortExists(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	exists, err := cli.DHCPPortExists(&system.NetworkingBody{
		PortName: deviceIface,
	})
	fmt.Println(exists)
	fmt.Println(err)
	if err != nil {
		return
	}

}
func TestClient_DHCPSetAsAuto(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	exists, err := cli.DHCPSetAsAuto(&system.NetworkingBody{
		PortName: deviceIface,
	})
	fmt.Println(exists)
	fmt.Println(err)
	if err != nil {
		return
	}

}

func TestClient_DHCPSetStaticIP(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
		Port:  0,
		HTTPS: false,
	})

	exists, err := cli.DHCPSetStaticIP(&dhcpd.SetStaticIP{
		Ip:                   "192.168.15.191",
		NetMask:              "255.255.255.0",
		IFaceName:            deviceIface,
		GatewayIP:            "192.168.15.1",
		DnsIP:                "8.8.8.8",
		CheckInterfaceExists: false,
		SaveFile:             true,
	})
	fmt.Println(exists)
	fmt.Println(err)
	if err != nil {
		return
	}

}
