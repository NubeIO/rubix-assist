package assitcli

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

func (inst *Client) GetPlugins(hostIDName string) ([]model.PluginConf, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.PluginConf{}).
		Get("proxy/ff/api/plugins"))
	if err != nil {
		return nil, err
	}
	var out []model.PluginConf
	out = *resp.Result().(*[]model.PluginConf)
	return out, nil
}

func (inst *Client) GetPlugin(hostIDName, uuid string) (*model.PluginConf, error) {
	url := fmt.Sprintf("proxy/ff/api/plugins/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
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

func (inst *Client) DisablePlugin(hostIDName, uuid string) (interface{}, error) {
	url := fmt.Sprintf("proxy/ff/api/plugins/enable/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
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

func (inst *Client) EnablePlugin(hostIDName, uuid string) (interface{}, error) {
	url := fmt.Sprintf("proxy/ff/api/plugins/enable/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
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

func (inst *Client) RestartPlugin(hostIDName, uuid string) (*Message, error) {
	url := fmt.Sprintf("proxy/ff/api/plugins/restart/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Post(url))
	if err != nil {
		return nil, err
	}
	return &Message{Message: resp.String()}, nil
}
