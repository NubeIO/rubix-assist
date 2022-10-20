package ffclient

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

// AsJson return as body as blank interface
func AsJson(res []byte) (interface{}, error) {
	var out interface{}
	err := json.Unmarshal(res, &out)
	if err != nil {
		return nil, err
	}
	return out, err
}

// NetworkSchema get network json schema
func (inst *FlowClient) NetworkSchema(pluginName string) (interface{}, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/schema/json/network", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		Get(url))
	if err != nil {
		return nil, err
	}
	asJson, err := AsJson(resp.Body())
	if err != nil {
		return nil, err
	}
	return asJson, nil
}

// DeviceSchema get device json schema
func (inst *FlowClient) DeviceSchema(pluginName string) (interface{}, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/schema/json/device", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		Get(url))
	if err != nil {
		return nil, err
	}
	asJson, err := AsJson(resp.Body())
	if err != nil {
		return nil, err
	}
	return asJson, nil
}

// PointSchema get point json schema
func (inst *FlowClient) PointSchema(pluginName string) (interface{}, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/schema/json/point", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		Get(url))
	if err != nil {
		return nil, err
	}
	asJson, err := AsJson(resp.Body())
	if err != nil {
		return nil, err
	}
	return asJson, nil
}
