package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
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
		"/api/apps/add/?name=%s&service_name=%s&version=%s&product=%s&arch=%s&do_not_validate_arch=%v",
		app.Name, app.ServiceName, app.Version, app.Product, app.Arch, app.DoNotValidateArch)
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
	url := fmt.Sprintf("/api/apps/service/upload/?name=%s&version=%s", appName, version)
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
func (inst *Client) EdgeUninstallApp(appName, serviceName string, deleteApp bool) (*installer.UninstallResponse, error) {
	url := fmt.Sprintf("/api/apps/?name=%s&service_name%s&delete=%v", appName, serviceName, deleteApp)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.UninstallResponse{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.UninstallResponse), nil
}

func (inst *Client) EdgeCtlAction(body *installer.CtlBody) (*systemctl.SystemResponse, error) {
	url := fmt.Sprintf("/api/apps/control/action")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&systemctl.SystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemResponse), nil
}

func (inst *Client) EdgeServiceMassAction(body *installer.CtlBody) ([]systemctl.MassSystemResponse, error) {
	url := fmt.Sprintf("/api/apps/control/action/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]systemctl.MassSystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]systemctl.MassSystemResponse)
	return *data, nil
}

func (inst *Client) EdgeCtlStatus(body *installer.CtlBody) (*systemctl.SystemState, error) {
	url := fmt.Sprintf("/api/apps/control/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&systemctl.SystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemState), nil
}

func (inst *Client) EdgeServiceMassStatus(body *installer.CtlBody) ([]systemctl.SystemState, error) {
	url := fmt.Sprintf("/api/apps/control/status/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]systemctl.SystemState{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]systemctl.SystemState)
	return *data, nil
}
