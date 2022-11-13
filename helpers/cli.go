package helpers

import (
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/edgebioscli"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
)

func GetEdgeClient(host *model.Host) *edgecli.Client {
	cli := edgecli.New(&edgecli.Client{
		Rest:          nil,
		Ip:            host.IP,
		Port:          host.Port,
		HTTPS:         host.HTTPS,
		ExternalToken: host.ExternalToken,
	})
	return cli
}

func GetEdgeBiosClient(host *model.Host) *edgebioscli.BiosClient {
	cli := edgebioscli.New(&edgebioscli.BiosClient{
		Rest:          nil,
		Ip:            host.IP,
		Port:          host.BiosPort,
		HTTPS:         host.HTTPS,
		ExternalToken: host.ExternalToken,
	})
	return cli
}
