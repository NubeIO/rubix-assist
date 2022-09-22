package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"io"
)

// ListAppsWithVersions list apps with versions
func (inst *Client) ListAppsWithVersions() ([]appstore.ListApps, error) {
	url := fmt.Sprintf("/api/store/apps")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]appstore.ListApps{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]appstore.ListApps), nil
}

// UploadAddOnAppStore upload an app
func (inst *Client) UploadAddOnAppStore(appName, version, arch, fileName string, reader io.Reader) (*appstore.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/apps/?name=%s&version=%s&arch=%s", appName, version, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&appstore.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.UploadResponse), nil
}
