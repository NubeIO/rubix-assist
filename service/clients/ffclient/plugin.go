package ffclient

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

func (inst *FlowClient) GetPlugins() ([]model.PluginConf, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.PluginConf{}).
		Get("/api/plugins"))
	if err != nil {
		return nil, err
	}
	var out []model.PluginConf
	out = *resp.Result().(*[]model.PluginConf)
	return out, nil
}

func (inst *FlowClient) GetPlugin(uuid string) (*model.PluginConf, error) {
	url := fmt.Sprintf("/api/plugins/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.PluginConf{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.PluginConf), nil
}

type enablePlugin struct {
	Enabled bool `json:"enabled"`
}

func (inst *FlowClient) DisablePlugin(uuid string) (interface{}, error) {
	url := fmt.Sprintf("/api/plugins/enable/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&ResponseBody{}).
		SetBody(enablePlugin{Enabled: false}).
		Post(url))
	if err != nil {
		return nil, err
	}
	var out interface{}
	err = json.Unmarshal(resp.Body(), &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (inst *FlowClient) EnablePlugin(uuid string) (interface{}, error) {
	url := fmt.Sprintf("/api/plugins/enable/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&ResponseBody{}).
		SetBody(enablePlugin{Enabled: true}).
		Post(url))
	if err != nil {
		return nil, err
	}
	var out interface{}
	err = json.Unmarshal(resp.Body(), &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (inst *FlowClient) CreateNetworkPlugin(body *model.Network, pluginName string) (*model.Network, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/networks", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *FlowClient) DeleteNetworkPlugin(body *model.Network, pluginName string) (ok bool, err error) {
	url := fmt.Sprintf("/api/plugins/api/%s/networks", pluginName)
	_, err = nresty.FormatRestyResponse(inst.client.R().
		SetBody(body).
		Delete(url))
	if err != nil {
		return false, err
	}
	return true, err
}

func (inst *FlowClient) DeleteDevicePlugin(body *model.Device, pluginName string) (ok bool, err error) {
	url := fmt.Sprintf("/api/plugins/api/%s/devices", pluginName)
	_, err = nresty.FormatRestyResponse(inst.client.R().
		SetBody(body).
		Delete(url))
	if err != nil {
		return false, err
	}
	return true, err
}

func (inst *FlowClient) DeletePointPlugin(body *model.Point, pluginName string) (ok bool, err error) {
	url := fmt.Sprintf("/api/plugins/api/%s/points", pluginName)
	_, err = nresty.FormatRestyResponse(inst.client.R().
		SetBody(body).
		Delete(url))
	if err != nil {
		return false, err
	}
	return true, err
}

func (inst *FlowClient) CreateDevicePlugin(body *model.Device, pluginName string) (*model.Device, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/devices", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Device{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *FlowClient) CreatePointPlugin(body *model.Point, pluginName string) (*model.Point, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/points", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Point{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *FlowClient) UpdateNetworkPlugin(body *model.Network, pluginName string) (*model.Network, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/networks", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Network{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *FlowClient) UpdateDevicePlugin(body *model.Device, pluginName string) (*model.Device, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/devices", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Device{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *FlowClient) UpdatePointPlugin(body *model.Point, pluginName string) (*model.Point, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/points", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Point{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *FlowClient) WritePointPlugin(pointUUID string, body *model.PointWriter, pluginName string) (*model.Point, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/points/write/{uuid}", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Point{}).
		SetBody(body).
		SetPathParams(map[string]string{"uuid": pointUUID}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}
