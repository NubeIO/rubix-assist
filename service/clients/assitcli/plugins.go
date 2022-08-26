package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"io"
)

// StoreListPlugins upload a plugin to the store
func (inst *Client) StoreListPlugins() ([]installer.BuildDetails, error) {
	url := fmt.Sprintf("/api/store/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]installer.BuildDetails{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]installer.BuildDetails), nil
}

// StoreUploadPlugin upload a plugin to the store
func (inst *Client) StoreUploadPlugin(fileName string, reader io.Reader) (*appstore.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&appstore.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.UploadResponse), nil
}
