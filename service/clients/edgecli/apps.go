package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemd"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"io"
)

// ListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Client) ListApps() ([]model.Apps, error) {
	url := fmt.Sprintf("/api/apps")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.Apps{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]model.Apps)
	return *data, nil
}

// ListAppsStatus get all the apps status
func (inst *Client) ListAppsStatus() ([]model.AppsStatus, error) {
	url := fmt.Sprintf("/api/apps/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.AppsStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]model.AppsStatus)
	return *data, nil
}

// UploadServiceFile add/install a new an app service (service file needs to be needs the build on the edge device)
func (inst *Client) UploadServiceFile(appName, version, fileName string, reader io.Reader) (*model.UploadResponse, error) {
	url := fmt.Sprintf("/api/apps/service/upload?name=%s&version=%s", appName, version)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.UploadResponse), nil
}

// InstallService add/install a new an app service (service file needs to be needs the build on the edge device)
func (inst *Client) InstallService(body *model.Install) (*systemd.InstallResponse, error) {
	url := fmt.Sprintf("/api/apps/service/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&systemd.InstallResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemd.InstallResponse), nil
}
