package ffclient

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"

	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// ClientAddNetwork an object
func (a *FlowClient) ClientAddNetwork(pluginUUID string) (*ResponseBody, error) {
	name := uuid.ShortUUID()
	name = fmt.Sprintf("net_name_%s", name)
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetBody(map[string]string{"name": name, "plugin_conf_id": pluginUUID}).
		Post("/api/networks"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// ClientGetNetwork an object
func (a *FlowClient) ClientGetNetwork(uuid string) (*ResponseBody, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&ResponseBody{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/networks/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ResponseBody), nil
}

// GetNetworksWithPoints an object
func (a *FlowClient) GetNetworksWithPoints() (*[]model.Network, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&[]model.Network{}).
		Get("/api/networks/?with_points=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]model.Network), nil
}

// GetNetworkWithPoints an object
func (a *FlowClient) GetNetworkWithPoints(uuid string) (*model.Network, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&model.Network{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/api/networks/{uuid}?with_points=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}
