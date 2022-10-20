package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) EdgeSystemCtlAction(body *installer.SystemCtlBody) (*installer.SystemResponse, error) {
	url := fmt.Sprintf("/api/apps/control/action")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.SystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.SystemResponse), nil
}

func (inst *Client) EdgeSystemCtlStatus(body *installer.SystemCtlBody) (*systemctl.SystemState, error) {
	url := fmt.Sprintf("/api/apps/control/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&systemctl.SystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemState), nil
}

func (inst *Client) EdgeServiceMassAction(body *installer.SystemCtlBody) ([]installer.MassSystemResponse, error) {
	url := fmt.Sprintf("/api/apps/control/action/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]installer.MassSystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.MassSystemResponse)
	return *data, nil
}

func (inst *Client) EdgeServiceMassStatus(body *installer.SystemCtlBody) ([]systemctl.SystemState, error) {
	url := fmt.Sprintf("/api/apps/control/status/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]systemctl.SystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]systemctl.SystemState)
	return *data, nil
}
