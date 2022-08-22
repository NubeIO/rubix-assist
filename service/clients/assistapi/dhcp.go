package assistapi

import (
	"fmt"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-edge/service/system"
)

func (inst *Client) EdgeDHCPPortExists(hostIDName string, body *system.NetworkingBody) (*system.Message, error) {
	url := fmt.Sprintf("proxy/api/networking/interfaces/exists/")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&system.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeDHCPSetAsAuto(hostIDName string, body *system.NetworkingBody) (*system.Message, error) {
	url := fmt.Sprintf("proxy/api/networking/interfaces/auto/")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&system.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeDHCPSetStaticIP(hostIDName string, body *dhcpd.SetStaticIP) (string, error) {
	url := fmt.Sprintf("proxy/api/networking/interfaces/static/")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetBody(body).
		Post(url))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
