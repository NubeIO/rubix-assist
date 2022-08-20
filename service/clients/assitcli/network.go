package assitcli

import (
	"encoding/json"
	"fmt"

	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
)

func (inst *Client) GetHostNetworks() (data []model.Network, response *Response) {
	path := fmt.Sprintf(Paths.HostNetwork.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Network{}).
		Get(path)
	return *resp.Result().(*[]model.Network), response.buildResponse(resp, err)
}

func (inst *Client) GetHostNetwork(uuid string) (data *model.Network, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.HostNetwork.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.Network{}).
		Get(path)
	return resp.Result().(*model.Network), response.buildResponse(resp, err)
}

func (inst *Client) AddHostNetwork(body *model.Network) (data *model.Network, response *Response) {
	path := fmt.Sprintf(Paths.HostNetwork.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Network{}).
		Post(path)
	return resp.Result().(*model.Network), response.buildResponse(resp, err)
}

func (inst *Client) UpdateHostNetwork(uuid string, body *model.Network) (data *model.Network, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.HostNetwork.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Network{}).
		Patch(path)
	return resp.Result().(*model.Network), response.buildResponse(resp, err)
}

func (inst *Client) DeleteHostNetwork(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.HostNetwork.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}

func (inst *Client) GetNetworkSchema() (data *model.NetworkSchema, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.HostNetwork.Path, "schema")
	response = &Response{}
	resp, err := inst.Rest.R().
		Get(path)
	var result *model.NetworkSchema
	err = json.Unmarshal(resp.Body(), &result)
	return result, response.buildResponse(resp, err)
}
