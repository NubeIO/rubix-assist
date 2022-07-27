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

// AddUploadStoreApp upload an app
func (inst *Client) AddUploadStoreApp(appName, version, product, arch, fileName string, reader io.Reader) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/add/?name=%s&version=%s&product=%s&arch=%s", appName, version, product, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&store.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}
