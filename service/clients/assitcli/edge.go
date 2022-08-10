package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-assist/service/store"
)

// AddUploadEdgeApp upload an app to the edge device
func (inst *Client) AddUploadEdgeApp(hostIDName string, app *store.EdgeApp) (*installer.AppResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/add")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.AppResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.AppResponse), nil
}

// UploadEdgeService generate a service file and upload it to edge device
func (inst *Client) UploadEdgeService(hostIDName string, app *store.ServiceFile) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/service/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&store.UploadResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}

// InstallEdgeService this assumes that the service file and app already exists on the edge device
func (inst *Client) InstallEdgeService(hostIDName string, body *installer.Install) (*installer.InstallResp, error) {
	url := fmt.Sprintf("/api/edge/apps/service/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.InstallResp{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.InstallResp), nil
}

// EdgeListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Client) EdgeListApps() (*installer.Apps, error) {
	url := fmt.Sprintf("/edge/api/apps")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.Apps{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.Apps), nil
}

// EdgeListAppsAndService get all the apps by listed in the installation (/data/rubix-service/apps/install) dir and then check the service
func (inst *Client) EdgeListAppsAndService() (*installer.InstalledServices, error) {
	url := fmt.Sprintf("/api/edge/apps/services")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.InstalledServices{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.InstalledServices), nil
}

// EdgeListNubeServices list all the services by filtering all the service files with name nubeio
func (inst *Client) EdgeListNubeServices() (*installer.InstalledServices, error) {
	url := fmt.Sprintf("/api/edge/services/nube")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.InstalledServices{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.InstalledServices), nil
}

func (inst *Client) EdgeCtlAction(body *installer.CtlBody) (*systemctl.SystemResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/control/action")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&systemctl.SystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemResponse), nil
}

func (inst *Client) EdgeServiceMassAction(body *installer.CtlBody) (*[]systemctl.MassSystemResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/control/action/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]systemctl.MassSystemResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]systemctl.MassSystemResponse), nil
}

func (inst *Client) EdgeCtlStatus(body *installer.CtlBody) (*systemctl.SystemResponseChecks, error) {
	url := fmt.Sprintf("/api/edge/apps/control/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&systemctl.SystemResponseChecks{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemResponseChecks), nil
}

func (inst *Client) EdgeServiceMassStatus(body *installer.CtlBody) ([]systemctl.MassSystemResponseChecks, error) {
	url := fmt.Sprintf("/api/edge/apps/control/status/mass")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]systemctl.MassSystemResponseChecks{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().([]systemctl.MassSystemResponseChecks), nil
}
