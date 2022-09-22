package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

type EdgeUploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}

// UploadPlugin upload a plugin to the edge device
func (inst *Client) UploadPlugin(hostIDName string, body *appstore.Plugin) (*EdgeUploadResponse, error) {
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

// ListPlugins list all the plugin in the dir /flow-framework/data/plugins
func (inst *Client) ListPlugins(hostIDName string) ([]appstore.Plugin, error) {
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

func (inst *Client) DeletePlugin(hostIDName string, body *appstore.Plugin) (*model.Message, error) {
	url := fmt.Sprintf("/api/edge/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Message{}).
		SetBody(body).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

func (inst *Client) DeleteAllPlugins(hostIDName string) (*model.Message, error) {
	url := fmt.Sprintf("/api/edge/plugins/all")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}
