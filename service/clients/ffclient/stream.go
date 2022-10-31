package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

// AddStreamToExistingFlow add a stream to an existing flow
func (inst *FlowClient) AddStreamToExistingFlow(flowNetworkUUID string, body *model.Stream, remote bool, args Remote) (*model.Stream, error) {
	url := fmt.Sprintf("/api/streams")
	if remote {
		url = fmt.Sprintf("/api/remote/streams/?flow_network_uuid=%s", args.FlowNetworkUUID)
	}
	flowNetwork := &model.FlowNetwork{
		CommonFlowNetwork: model.CommonFlowNetwork{
			CommonUUID: model.CommonUUID{
				UUID: flowNetworkUUID,
			},
		},
	}
	body.FlowNetworks = append(body.FlowNetworks, flowNetwork)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Stream{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Stream), nil
}

func (inst *FlowClient) AddStream(body *model.Stream, remote bool, args Remote) (*model.Stream, error) {
	url := fmt.Sprintf("/api/streams")
	if remote {
		url = fmt.Sprintf("/api/remote/streams/?flow_network_uuid=%s", args.FlowNetworkUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Stream{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Stream), nil
}

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

func (inst *FlowClient) GetStreamClones(remote bool, args Remote) ([]model.StreamClone, error) {
	url := fmt.Sprintf("/api/stream_clones")
	if remote {
		url = fmt.Sprintf("/api/remote/stream_clones/?flow_network_uuid=%s", args.FlowNetworkUUID)
	}
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

func (inst *FlowClient) GetStreams(remote bool, args Remote) ([]model.Stream, error) {
	url := fmt.Sprintf("/api/streams")
	if remote {
		url = fmt.Sprintf("/api/remote/streams/?flow_network_uuid=%s", args.FlowNetworkUUID)
	}
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

func (inst *FlowClient) GetStreamsWithChild(remote bool, args Remote) ([]model.Stream, error) {
	url := fmt.Sprintf("/api/streams?flow_networks=true&producers=true&consumers=true&command_groups=false&writers=true&tags=true")
	if remote {
		url = fmt.Sprintf("/api/remote/streams/?flow_networks=true&producers=true&consumers=true&command_groups=false&writers=true&tags=true&flow_network_uuid=%s", args.FlowNetworkUUID)
	}
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

func (inst *FlowClient) DeleteStream(uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.client.R().
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/api/streams/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *FlowClient) DeleteStreamClone(uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.client.R().
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/api/stream_clones/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
