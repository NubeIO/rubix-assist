package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"github.com/NubeIO/rubix-edge/service/system"
)

func (inst *Client) DHCPPortExists(body *system.NetworkingBody) (*system.DHCPPortExists, error) {
	url := fmt.Sprintf("/api/networking/interfaces/exists")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&system.DHCPPortExists{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.DHCPPortExists), nil
}

func (inst *Client) DHCPSetAsAuto(body *system.NetworkingBody) (*system.Message, error) {
	url := fmt.Sprintf("/api/networking/interfaces/auto")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&system.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) DHCPSetStaticIP(body *dhcpd.SetStaticIP) (string, error) {
	url := fmt.Sprintf("/api/networking/interfaces/static")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		Post(url))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
