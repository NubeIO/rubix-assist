package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-assist/service/store"
	"io"
)

// ListStore list apps and store
func (inst *Client) ListStore() (*[]store.App, error) {
	url := fmt.Sprintf("/api/store")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]store.App{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]store.App), nil
}

// UploadApp upload an app
func (inst *Client) UploadApp(appName, version, fileName string, reader io.Reader) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/upload/?name=%s&version=%s", appName, version)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&store.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}

// AddApp add a new an app
func (inst *Client) AddApp(app *store.App) (*store.App, error) {
	url := fmt.Sprintf("/api/store")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&store.App{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.App), nil
}
