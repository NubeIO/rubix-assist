package ffclient

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// ClientAddDevice an object
func (a *FlowClient) ClientAddDevice(networkUUID string) (*ResponseBody, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("dev_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetBody(map[string]string{"name": name, "network_uuid": networkUUID}).
		Post("/api/devices"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// ClientGetDevice an object
func (a *FlowClient) ClientGetDevice(uuid string) (*ResponseBody, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/devices/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// ClientEditDevice edit an object
func (a *FlowClient) ClientEditDevice(uuid_ string) (*ResponseBody, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("dev_new_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetBody(map[string]string{"name": name}).
		SetPathParams(map[string]string{"uuid": uuid_}).
		Post("/api/devices/{}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}
