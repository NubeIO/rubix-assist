package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemd"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"io"
)

// ListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Client) ListApps() ([]installer.Apps, error) {
	url := fmt.Sprintf("/api/apps")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]installer.Apps{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.Apps)
	return *data, nil
}

// ListAppsStatus get all the apps status
func (inst *Client) ListAppsStatus() ([]installer.AppsStatus, error) {
	url := fmt.Sprintf("/api/apps/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]installer.AppsStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.AppsStatus)
	return *data, nil
}

// UploadApp uploads an app
func (inst *Client) UploadApp(app *installer.Upload, zipFileName string, reader io.Reader) (*installer.AppResponse, error) {
	url := fmt.Sprintf(
		"/api/apps/upload?name=%s&version=%s&product=%s&arch=%s&do_not_validate_arch=%v&move_extracted_file_to_name_app=%v&move_one_level_inside_file_to_outside=%v",
		app.Name, app.Version, app.Product, app.Arch, app.DoNotValidateArch, app.MoveExtractedFileToNameApp, app.MoveOneLevelInsideFileToOutside)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.AppResponse{}).
		SetFileReader("file", zipFileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.AppResponse), nil
}

// UploadServiceFile add/install a new an app service (service file needs to be needs the build on the edge device)
func (inst *Client) UploadServiceFile(appName, version, fileName string, reader io.Reader) (*installer.UploadResponse, error) {
	url := fmt.Sprintf("/api/apps/service/upload?name=%s&version=%s", appName, version)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.UploadResponse), nil
}

// InstallService add/install a new an app service (service file needs to be needs the build on the edge device)
func (inst *Client) InstallService(body *installer.Install) (*systemd.InstallResponse, error) {
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

// EdgeUninstallApp remove/delete an app and its service
func (inst *Client) EdgeUninstallApp(appName string, deleteApp bool) (*installer.UninstallResponse, error) {
	url := fmt.Sprintf("/api/apps?name=%s&delete=%v", appName, deleteApp)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.UninstallResponse{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.UninstallResponse), nil
}
