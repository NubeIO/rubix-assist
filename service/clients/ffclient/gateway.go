package ffclient

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// ClientAddGateway an object
func (a *FlowClient) ClientAddGateway(body *model.Stream) (*ResponseBody, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("gte_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetBody(body).
		Post("/api/streams"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// ClientGetGateway an object
func (a *FlowClient) ClientGetGateway(uuid string) (*ResponseBody, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/streams/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// ClientEditGateway edit an object
func (a *FlowClient) ClientEditGateway(uuid_ string) (*ResponseBody, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("dev_new_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetBody(map[string]string{"name": name}).
		SetPathParams(map[string]string{"uuid": uuid_}).
		Post("/api/streams/{}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}
