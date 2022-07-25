package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-assist/service/store"
)

// UploadEdgeApp upload an app
func (inst *Client) UploadEdgeApp(app *store.EdgeApp) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&store.UploadResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}

// InstallEdgeApp upload an app
func (inst *Client) InstallEdgeApp(app *store.EdgeApp) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&store.UploadResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}
