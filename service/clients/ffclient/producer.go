package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *FlowClient) AddProducer(body *model.Producer, remote bool, args Remote) (*model.Producer, error) {
	url := fmt.Sprintf("/api/producers")
	if remote {
		url = fmt.Sprintf("/api/remote/producers/?flow_network_uuid=%s", args.FlowNetworkUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Producer{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *FlowClient) EditProducer(uuid string, body *model.Producer) (*model.Producer, error) {
	url := fmt.Sprintf("/api/producers/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Producer{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *FlowClient) GetProducers() ([]model.Producer, error) {
	url := fmt.Sprintf("/api/producers")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Producer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Producer
	out = *resp.Result().(*[]model.Producer)
	return out, nil
}

func (inst *FlowClient) GetProducer(uuid string) (*model.Producer, error) {
	url := fmt.Sprintf("/api/producers/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Producer{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Producer), nil
}

func (inst *FlowClient) DeleteProducer(uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.client.R().
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/api/producers/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
