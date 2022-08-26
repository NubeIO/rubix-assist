package assitcli

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

// AddPoint an object
func (inst *Client) AddPoint(hostIDName string, body *model.Point) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Point{}).
		SetBody(body).
		Post("proxy/ff/api/points"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

// GetPoints an objects
func (inst *Client) GetPoints(hostIDName string) ([]model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.Point{}).
		Get("proxy/ff/api/points"))
	if err != nil {
		return nil, err
	}
	var out []model.Point
	out = *resp.Result().(*[]model.Point)
	return out, nil
}

// GetPoint an object
func (inst *Client) GetPoint(hostIDName, uuid string) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("proxy/ff/api/points/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

// DeletePoint an object
func (inst *Client) DeletePoint(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("proxy/ff/api/points/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

// EditPoint an object
func (inst *Client) EditPoint(hostIDName, uuid string, body *model.Point) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetBody(body).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Patch("proxy/ff/api/points/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}
