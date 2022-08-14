package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"io"
	"strconv"
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

// ListAppsAndService get all the apps by listed in the installation (/data/rubix-service/apps/install) dir and then check the service
func (inst *Client) ListAppsAndService() ([]installer.InstalledServices, error) {
	url := fmt.Sprintf("/api/apps/services")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]installer.InstalledServices{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.InstalledServices)
	return *data, nil
}

// ListNubeServices list all the services by filtering all the service files with name nubeio
func (inst *Client) ListNubeServices() ([]installer.InstalledServices, error) {
	url := fmt.Sprintf("/api/apps/services/nube")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]installer.InstalledServices{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.InstalledServices)
	return *data, nil
}

// UploadApp upload an app
func (inst *Client) UploadApp(appName, version, productType, archType, fileName string, reader io.Reader) (*installer.AppResponse, error) {
	url := fmt.Sprintf("/api/apps/add/?name=%s&version=%s&product=%s&arch=%s", appName, version, productType, archType)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.AppResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.AppResponse), nil
}

// InstallApp add/install a new an app (needs the build on the edge device)
func (inst *Client) InstallApp(body *installer.Install) (*installer.AppResponse, error) {
	url := fmt.Sprintf("/api/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.AppResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.AppResponse), nil
}

// EdgeUnInstallApp remove/delete an app and its service
func (inst *Client) EdgeUnInstallApp(appName string, deleteApp bool) (*installer.RemoveRes, error) {
	url := fmt.Sprintf("/api/apps/?name=%s&delete=%s", appName, strconv.FormatBool(deleteApp))
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.RemoveRes{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.RemoveRes), nil
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
func (inst *Client) InstallService(body *installer.Install) (*installer.InstallResp, error) {
	url := fmt.Sprintf("/api/apps/service/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.InstallResp{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.InstallResp), nil

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
