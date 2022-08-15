package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// EdgeProductInfo get edge product info
func (inst *Client) EdgeProductInfo(hostIDName string) (*installer.Product, error) {
	url := fmt.Sprintf("/api/edge/system/product")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.Product{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.Product), nil
}

func (inst *Client) EdgeCtlAction(hostIDName string, body *installer.CtlBody) (*systemctl.SystemResponse, error) {
	url := fmt.Sprintf("/api/edge/control/action")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&systemctl.SystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemResponse), nil
}

func (inst *Client) EdgeServiceMassAction(hostIDName string, body *installer.CtlBody) ([]systemctl.MassSystemResponse, error) {
	url := fmt.Sprintf("/api/edge/control/action/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]systemctl.MassSystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]systemctl.MassSystemResponse)
	return *data, nil
}

func (inst *Client) EdgeCtlStatus(hostIDName string, body *installer.CtlBody) (*systemctl.SystemState, error) {
	url := fmt.Sprintf("/api/edge/control/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&systemctl.SystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemState), nil
}

func (inst *Client) EdgeServiceMassStatus(hostIDName string, body *installer.CtlBody) ([]systemctl.SystemState, error) {
	url := fmt.Sprintf("/api/edge/control/status/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]systemctl.SystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]systemctl.SystemState)
	return *data, nil
}
