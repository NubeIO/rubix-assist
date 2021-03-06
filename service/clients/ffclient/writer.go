package ffclient

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"strconv"
)

// GetWriter an object
func (inst *FlowClient) GetWriter(uuid string) (*model.Writer, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Writer{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/consumers/writers/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Writer), nil
}

// EditWriter edit an object
func (inst *FlowClient) EditWriter(uuid string, body model.Writer, updateProducer bool) (*model.Writer, error) {
	param := strconv.FormatBool(updateProducer)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Writer{}).
		SetBody(body).
		SetPathParams(map[string]string{"uuid": uuid}).
		SetQueryParam("update_producer", param).
		Patch("/api/consumers/writers/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Writer), nil
}

// CreateWriter edit an object
func (inst *FlowClient) CreateWriter(body model.Writer) (*model.Writer, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("sub_name_%s", name)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Writer{}).
		SetBody(body).
		Post("/api/consumers/writers"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Writer), nil
}
