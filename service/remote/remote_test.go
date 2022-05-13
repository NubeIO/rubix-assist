package remote

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/remote/ssh"
	"testing"
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
				IP:       "123.209.234.118",
				Port:     2022,
				Username: "pi",
				Password: "N00BRCRC",
			},
		},
	}
	run := New(host)
	out := run.Uptime()
	fmt.Println(out)

}
