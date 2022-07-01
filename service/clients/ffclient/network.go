package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// AddNetwork an object
func (inst *FlowClient) AddNetwork(body *model.Point) (*model.Network, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetBody(body).
		Post("/api/networks"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// EditNetwork edit an object
func (inst *FlowClient) EditNetwork(uuid string, device *model.Network) (*model.Network, error) {
	url := fmt.Sprintf("/api/networks/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetBody(device).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// GetNetworkByPluginName an object
func (inst *FlowClient) GetNetworkByPluginName(pluginName string, withPoints ...bool) (*model.Network, error) {
	url := fmt.Sprintf("/api/networks/plugin/%s", pluginName)
	if len(withPoints) > 0 {
		url = fmt.Sprintf("/api/networks/plugin/%s?with_points=true", pluginName)
	}
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// GetNetworks an object
func (inst *FlowClient) GetNetworks(withDevices ...bool) (*[]model.Network, error) {
	url := fmt.Sprintf("/api/networks")
	if len(withDevices) > 0 {
		url = fmt.Sprintf("/api/networks/?with_devices=true")
	}
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]model.Network), nil
}

// GetNetworksWithPoints an object
func (inst *FlowClient) GetNetworksWithPoints() (*[]model.Network, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Network{}).
		Get("/api/networks/?with_points=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]model.Network), nil
}

// GetNetworkWithPoints an object
func (inst *FlowClient) GetNetworkWithPoints(uuid string) (*model.Network, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/networks/{uuid}?with_points=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// GetNetwork an object
func (inst *FlowClient) GetNetwork(uuid string) (*model.Network, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/networks/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// GetFirstNetwork first object
func (inst *FlowClient) GetFirstNetwork(withDevices ...bool) (*model.Network, error) {
	nets, err := inst.GetNetworks(withDevices...)
	if err != nil {
		return nil, err
	}
	for _, net := range *nets {
		return &net, err
	}
	return nil, err
}
