package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) EdgeSystemCtlAction(body *model.SystemCtlBody) (*model.SystemResponse, error) {
	url := fmt.Sprintf("/api/apps/control/action")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.SystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.SystemResponse), nil
}

func (inst *Client) EdgeSystemCtlStatus(body *model.SystemCtlBody) (*systemctl.SystemState, error) {
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

func (inst *Client) EdgeServiceMassAction(body *model.SystemCtlBody) ([]model.MassSystemResponse, error) {
	url := fmt.Sprintf("/api/apps/control/action/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.MassSystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]model.MassSystemResponse)
	return *data, nil
}

func (inst *Client) EdgeServiceMassStatus(body *model.SystemCtlBody) ([]systemctl.SystemState, error) {
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
