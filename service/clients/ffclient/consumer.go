package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// AddConsumer an object
func (inst *FlowClient) AddConsumer(body *model.Consumer) (*model.Consumer, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Consumer{}).
		SetBody(body).
		Post("/api/consumers"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Consumer), nil
}

// EditConsumer edit an object
func (inst *FlowClient) EditConsumer(uuid string, body *model.Consumer) (*model.Consumer, error) {
	url := fmt.Sprintf("/api/consumers/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Consumer{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Consumer), nil
}

// GetConsumers an object
func (inst *FlowClient) GetConsumers() ([]model.Consumer, error) {
	url := fmt.Sprintf("/api/consumers")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Consumer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Consumer
	out = *resp.Result().(*[]model.Consumer)
	return out, nil
}

// GetConsumer an object
func (inst *FlowClient) GetConsumer(uuid string) (*model.Consumer, error) {
	url := fmt.Sprintf("/api/consumers/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Consumer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Consumer), nil
}

// DeleteConsumer an object
func (inst *FlowClient) DeleteConsumer(uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.client.R().
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/api/consumers/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
