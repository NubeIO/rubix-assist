package assitcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

func (inst *Client) AddConsumer(hostIDName string, body *model.Consumer) (*model.Consumer, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Consumer{}).
		SetBody(body).
		Post("proxy/ff/api/consumers"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Consumer), nil
}

func (inst *Client) EditConsumer(hostIDName, uuid string, body *model.Consumer) (*model.Consumer, error) {
	url := fmt.Sprintf("proxy/ff/api/consumers/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Consumer{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Consumer), nil
}

func (inst *Client) GetConsumers(hostIDName string) ([]model.Consumer, error) {
	url := fmt.Sprintf("proxy/ff/api/consumers")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.Consumer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Consumer
	out = *resp.Result().(*[]model.Consumer)
	return out, nil
}

func (inst *Client) GetConsumer(hostIDName, uuid string) (*model.Consumer, error) {
	url := fmt.Sprintf("proxy/ff/api/consumers/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Consumer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Consumer), nil
}

func (inst *Client) DeleteConsumer(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("proxy/ff/api/consumers/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
