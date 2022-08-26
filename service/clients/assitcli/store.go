package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
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

// ListAppsBuildDetails list apps with arch
func (inst *Client) ListAppsBuildDetails() ([]installer.BuildDetails, error) {
	url := fmt.Sprintf("/api/store/apps/details")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]installer.BuildDetails{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]installer.BuildDetails), nil
}

// AddUploadStoreApp upload an app
func (inst *Client) AddUploadStoreApp(appName, version, product, arch, fileName string, reader io.Reader) (*appstore.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/apps/?name=%s&version=%s&product=%s&arch=%s", appName, version, product, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&appstore.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.UploadResponse), nil
}
