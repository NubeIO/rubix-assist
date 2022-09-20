package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
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

// EdgePing ping a device
func (inst *Client) EdgePing(hostIDName string) (*model.Message, error) {
	url := fmt.Sprintf("/api/edge/system/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

func (inst *Client) EdgeSystemCtlAction(hostIDName string, body *installer.SystemCtlBody) (*installer.SystemResponse, error) {
	url := fmt.Sprintf("/api/edge/control/action")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.SystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.SystemResponse), nil
}

func (inst *Client) EdgeServiceMassAction(hostIDName string, body *installer.SystemCtlBody) ([]installer.MassSystemResponse, error) {
	url := fmt.Sprintf("/api/edge/control/action/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]installer.MassSystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.MassSystemResponse)
	return *data, nil
}

func (inst *Client) EdgeSystemCtlStatus(hostIDName string, body *installer.SystemCtlBody) (*installer.AppSystemState, error) {
	url := fmt.Sprintf("/api/edge/control/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.AppSystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.AppSystemState), nil
}

func (inst *Client) EdgeServiceMassStatus(hostIDName string, body *installer.SystemCtlBody) ([]installer.AppSystemState, error) {
	url := fmt.Sprintf("/api/edge/control/status/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]installer.AppSystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.AppSystemState)
	return *data, nil
}
