package assistcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assistcli/nresty"
)

func (inst *Client) AddFlowNetwork(hostIDName string, body *model.FlowNetwork) (*model.FlowNetwork, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.FlowNetwork{}).
		SetBody(body).
		Post("proxy/ff/api/flow_networks"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetwork), nil
}

func (inst *Client) EditFlowNetwork(hostIDName, uuid string, body *model.FlowNetwork) (*model.FlowNetwork, error) {
	url := fmt.Sprintf("proxy/ff/api/flow_networks/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.FlowNetwork{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetwork), nil
}

func (inst *Client) GetFlowNetwork(hostIDName, uuid string, withStreams bool, overrideUrl ...string) (*model.FlowNetwork, error) {
	url := fmt.Sprintf("proxy/ff/api/flow_networks/%s", uuid)
	if withStreams == true {
		url = fmt.Sprintf("proxy/ff/api/flow_networks/%s?with_streams=true", uuid)
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.FlowNetwork{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetwork), nil
}

func (inst *Client) GetFlowNetworks(hostIDName string, withStreams bool, overrideUrl ...string) ([]model.FlowNetwork, error) {
	url := fmt.Sprintf("proxy/ff/api/flow_networks")
	if withStreams == true {
		url = fmt.Sprintf("proxy/ff/api/flow_networks?with_streams=true")
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.FlowNetwork{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.FlowNetwork
	out = *resp.Result().(*[]model.FlowNetwork)
	return out, nil
}

func (inst *Client) GetFlowNetworksWithChild(hostIDName string) ([]model.FlowNetwork, error) {
	url := fmt.Sprintf("proxy/ff/api/flow_networks?with_streams=true&with_producers=true&with_consumers=true&with_writers=true&with_tags=true")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.FlowNetwork{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.FlowNetwork
	out = *resp.Result().(*[]model.FlowNetwork)
	return out, nil
}

func (inst *Client) GetFlowNetworkWithChild(hostIDName, uuid string) (*model.FlowNetwork, error) {
	url := fmt.Sprintf("proxy/ff/api/flow_networks%s?with_streams=true&with_producers=true&with_consumers=true&with_writers=true&with_tags=true", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.FlowNetwork{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetwork), nil
}

func (inst *Client) DeleteFlowNetwork(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("proxy/ff/api/flow_networks/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
