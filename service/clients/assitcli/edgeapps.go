package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"strconv"
)

// AddUploadEdgeApp upload an app to the edge device
func (inst *Client) AddUploadEdgeApp(hostIDName string, app *appstore.EdgeApp) (*installer.AppResponse, error) {
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
func (inst *Client) UploadEdgeService(hostIDName string, app *appstore.ServiceFile) (*appstore.UploadResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/service/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&appstore.UploadResponse{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.UploadResponse), nil
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

// EdgeUnInstallApp remove/delete an app and its service
func (inst *Client) EdgeUnInstallApp(hostIDName, appName string, deleteApp bool) (*installer.RemoveRes, error) {
	url := fmt.Sprintf("/api/edge/apps/?name=%s&delete=%s", appName, strconv.FormatBool(deleteApp))
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.RemoveRes{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.RemoveRes), nil
}

// EdgeListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Client) EdgeListApps(hostIDName string) ([]installer.Apps, error) {
	url := fmt.Sprintf("/api/edge/apps")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]installer.Apps{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.Apps)
	return *data, nil
}

// EdgeListAppsAndService get all the apps by listed in the installation (/data/rubix-service/apps/install) dir and then check the service
func (inst *Client) EdgeListAppsAndService(hostIDName string) ([]installer.InstalledServices, error) {
	url := fmt.Sprintf("/api/edge/apps/services")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]installer.InstalledServices{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.InstalledServices)
	return *data, nil
}

// EdgeListNubeServices list all the services by filtering all the service files with name nubeio
func (inst *Client) EdgeListNubeServices(hostIDName string) ([]installer.InstalledServices, error) {
	url := fmt.Sprintf("/api/edge/apps/services/nube")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]installer.InstalledServices{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]installer.InstalledServices)
	return *data, nil
}
