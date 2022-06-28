package ffclient

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

// ClientAddPoint an object
func (a *FlowClient) ClientAddPoint(deviceUUID string) (*model.Point, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("pnt_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&model.Point{}).
		SetBody(map[string]string{"name": name, "device_uuid": deviceUUID}).
		Post("/api/points"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

// ClientGetPoint an object
func (a *FlowClient) ClientGetPoint(uuid string) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/points/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

// ClientEditPoint an object
func (a *FlowClient) ClientEditPoint(uuid string, body model.Point) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetBody(body).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Patch("/api/points/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}
