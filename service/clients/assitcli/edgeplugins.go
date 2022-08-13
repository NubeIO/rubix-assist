package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

type EdgeUploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}

// EdgeUploadPlugin upload a plugin to the edge device
func (inst *Client) EdgeUploadPlugin(hostIDName string, body *appstore.Plugin) (*EdgeUploadResponse, error) {
	url := fmt.Sprintf("/api/edge/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&EdgeUploadResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*EdgeUploadResponse), nil
}

// EdgeListPlugins list all the plugin in the dir /flow-framework/data/plugins
func (inst *Client) EdgeListPlugins(hostIDName string) ([]appstore.Plugin, error) {
	url := fmt.Sprintf("/api/edge/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]appstore.Plugin{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]appstore.Plugin)
	return *data, nil
}

// EdgeDeletePlugin delete one
func (inst *Client) EdgeDeletePlugin(hostIDName string, body *appstore.Plugin) (*edgecli.Message, error) {
	url := fmt.Sprintf("/api/edge/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&edgecli.Message{}).
		SetBody(body).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*edgecli.Message), nil
}

// EdgeDeleteAllPlugins delete all
func (inst *Client) EdgeDeleteAllPlugins(hostIDName string) (*edgecli.Message, error) {
	url := fmt.Sprintf("/api/edge/plugins/all")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&edgecli.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*edgecli.Message), nil
}
