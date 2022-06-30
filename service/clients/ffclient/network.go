package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"

	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// GetNetworkByPluginName an object
func (inst *FlowClient) GetNetworkByPluginName(pluginName string) (*model.Network, error) {
	url := fmt.Sprintf("/api/networks/plugin/%s", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// GetNetworkByPluginNameWithPoints an object
func (inst *FlowClient) GetNetworkByPluginNameWithPoints(pluginName string) (*model.Network, error) {
	url := fmt.Sprintf("/api/networks/plugin/%s/?with_points=true", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
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
