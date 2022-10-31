package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *FlowClient) AddNetwork(body *model.Network) (*model.Network, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetBody(body).
		Post("/api/networks"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *FlowClient) EditNetwork(uuid string, device *model.Network) (*model.Network, error) {
	url := fmt.Sprintf("/api/networks/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetBody(device).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

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

func (inst *FlowClient) GetNetworks(remote bool, args Remote, withDevices bool) ([]model.Network, error) {
	url := fmt.Sprintf("/api/networks")
	if withDevices {
		url = fmt.Sprintf("/api/networks?with_devices=true")
	}
	if remote {
		url = fmt.Sprintf("/api/remote/networks/?flow_network_uuid=%s", args.FlowNetworkUUID)
		if withDevices {
			url = fmt.Sprintf("/api/remote/networks/?with_devices=true&flow_network_uuid=%s", args.FlowNetworkUUID)
		}
	}

	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Network
	out = *resp.Result().(*[]model.Network)
	return out, nil
}

func (inst *FlowClient) GetNetworksWithPoints() ([]model.Network, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Network{}).
		Get("/api/networks?with_points=true"))
	if err != nil {
		return nil, err
	}
	var out []model.Network
	out = *resp.Result().(*[]model.Network)
	return out, nil
}

func (inst *FlowClient) GetNetworkWithPoints(uuid string, remote bool, args Remote) (*model.Network, error) {
	url := fmt.Sprintf("/api/networks/%s/?with_points=true", uuid)
	if remote {
		url = fmt.Sprintf("/api/remote/networks/%s/?with_points=true&flow_network_uuid=%s", uuid, args.FlowNetworkUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}
