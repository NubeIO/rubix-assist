package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// AddStream an object
func (inst *FlowClient) AddStream(body *model.Stream) (*model.Stream, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Stream{}).
		SetBody(body).
		Post("/api/streams"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Stream), nil
}

// EditStream edit an object
func (inst *FlowClient) EditStream(uuid string, body *model.Stream) (*model.Stream, error) {
	url := fmt.Sprintf("/api/streams/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Stream{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Stream), nil
}

// GetStreamClones an object
func (inst *FlowClient) GetStreamClones() ([]model.StreamClone, error) {
	url := fmt.Sprintf("/api/stream_clones")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.StreamClone{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.StreamClone
	out = *resp.Result().(*[]model.StreamClone)
	return out, nil
}

// GetStreams an object
func (inst *FlowClient) GetStreams() ([]model.Stream, error) {
	url := fmt.Sprintf("/api/streams")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Stream{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Stream
	out = *resp.Result().(*[]model.Stream)
	return out, nil
}

// GetStream an object
func (inst *FlowClient) GetStream(uuid string) (*model.Stream, error) {
	url := fmt.Sprintf("/api/streams/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Stream{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Stream), nil
}

// GetStreamsWithChild an object
func (inst *FlowClient) GetStreamsWithChild() ([]model.Stream, error) {
	url := fmt.Sprintf("/api/streams?flow_networks=true&producers=true&consumers=true&command_groups=false&writers=true&tags=true")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Stream{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Stream
	out = *resp.Result().(*[]model.Stream)
	return out, nil
}

// GetStreamWithChild an object
func (inst *FlowClient) GetStreamWithChild(uuid string) (*model.Stream, error) {
	url := fmt.Sprintf("/api/streams/%s?flow_networks=true&producers=true&consumers=true&command_groups=false&writers=true&tags=true", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Stream{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Stream), nil
}

// DeleteStream an object
func (inst *FlowClient) DeleteStream(uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.client.R().
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/api/streams/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
