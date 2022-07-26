package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-assist/service/store"
)

// AddUploadEdgeApp upload an app to the edge device
func (inst *Client) AddUploadEdgeApp(hostIDName string, app *store.EdgeApp) (*installer.AppResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/add")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.AppResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.AppResponse), nil
}

// UploadEdgeService generate a service file and upload it to edge device
func (inst *Client) UploadEdgeService(hostIDName string, app *store.ServiceFile) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/service/upload")
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

// InstallEdgeService this assumes that the service file and app already exists on the edge device
func (inst *Client) InstallEdgeService(hostIDName string, body *installer.Install) (*installer.InstallResp, error) {
	url := fmt.Sprintf("/api/edge/apps/service/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.InstallResp{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.InstallResp), nil
}
