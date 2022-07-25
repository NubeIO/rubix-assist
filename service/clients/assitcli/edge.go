package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-assist/service/store"
)

// UploadEdgeApp upload an app
func (inst *Client) UploadEdgeApp(hostIDName string, app *store.EdgeApp) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&store.UploadResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}

// InstallEdgeApp upload an app
func (inst *Client) InstallEdgeApp(hostIDName string, app *store.EdgeApp) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&store.UploadResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}

// UploadEdgeService gener a
func (inst *Client) UploadEdgeService(hostIDName string, app *store.EdgeApp) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&store.UploadResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}
