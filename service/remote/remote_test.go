package remote

import (
	"fmt"
	"testing"

	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/remote/ssh"
)

func TestLocalConnection(t *testing.T) {
	host := &Admin{
		SSH: &ssh.Host{
			Host: &model.Host{
				IsLocalhost: nils.NewBool(true),
			},
		},
	}
	run := New(host)
	run.Uptime()

}

func TestRemoteConnection(t *testing.T) {
	host := &Admin{
		SSH: &ssh.Host{
			Host: &model.Host{
				IP:       "192.168.15.103",
				Port:     22,
				Username: "debian",
				Password: "N00B2828",
			},
		},
	}
	run := New(host)
	out, err := run.EdgeSetIP(&EdgeNetworking{IPAddress: "192.168.15.103", SubnetMask: "255.255.255.0", Gateway: "192.168.15.1"})
	fmt.Println(out, err)

}
