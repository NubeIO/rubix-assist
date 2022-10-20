package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"strconv"
)

func (inst *FlowClient) GetWriterClones() ([]model.WriterClone, error) {
	url := fmt.Sprintf("/api/producers/writer_clones")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.WriterClone{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.WriterClone
	out = *resp.Result().(*[]model.WriterClone)
	return out, nil
}

func (inst *FlowClient) GetWriterClone(uuid string) (*model.WriterClone, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.WriterClone{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/producers/writer_clones/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.WriterClone), nil
}

func (inst *FlowClient) EditWriterClone(uuid string, body model.WriterClone, updateProducer bool) (*model.WriterClone, error) {
	param := strconv.FormatBool(updateProducer)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.WriterClone{}).
		SetBody(body).
		SetPathParams(map[string]string{"uuid": uuid}).
		SetQueryParam("update_producer", param).
		Patch("/api/producers/writer_clones/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.WriterClone), nil
}

func (inst *FlowClient) CreateWriterClone(body model.WriterClone) (*model.WriterClone, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.WriterClone{}).
		SetBody(body).
		Post("/api/producers/writer_clones"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.WriterClone), nil
}

func (inst *FlowClient) DeleteWriterClone(uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.client.R().
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/api/producers/writer_clones/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
