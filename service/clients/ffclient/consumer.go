package ffclient

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// ClientAddConsumer an object
func (a *FlowClient) ClientAddConsumer(body Consumer) (*ResponseBody, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("sub_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetBody(body).
		Post("/api/consumers"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// ClientGetConsumer an object
func (a *FlowClient) ClientGetConsumer(uuid string) (*ResponseBody, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/consumers/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// ClientEditConsumer edit an object
func (a *FlowClient) ClientEditConsumer(uuid_ string) (*ResponseBody, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("sub_new_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetBody(map[string]string{"name": name}).
		SetPathParams(map[string]string{"uuid": uuid_}).
		Post("/api/consumers/{}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}
