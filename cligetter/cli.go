package cligetter

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/edgebioscli"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	"github.com/NubeIO/rubix-assist/service/clients/openvpncli"
)

func GetEdgeClient(host *amodel.Host) *edgecli.Client {
	cli := edgecli.New(&edgecli.Client{
		Rest:          nil,
		Ip:            host.IP,
		Port:          host.Port,
		HTTPS:         host.HTTPS,
		ExternalToken: host.ExternalToken,
	})
	return cli
}

func GetEdgeClientFastTimeout(host *amodel.Host) *edgecli.Client {
	cli := edgecli.NewFastTimeout(&edgecli.Client{
		Rest:          nil,
		Ip:            host.IP,
		Port:          host.Port,
		HTTPS:         host.HTTPS,
		ExternalToken: host.ExternalToken,
	})
	return cli
}

func GetEdgeBiosClient(host *amodel.Host) *edgebioscli.BiosClient {
	cli := edgebioscli.New(&edgebioscli.BiosClient{
		Rest:          nil,
		Ip:            host.IP,
		Port:          host.BiosPort,
		HTTPS:         host.HTTPS,
		ExternalToken: host.ExternalToken,
	})
	return cli
}

func GetOpenVPNClient() (*openvpncli.OpenVPNClient, error) {
	return openvpncli.Get()
}
