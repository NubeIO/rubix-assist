package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemd"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"strconv"
)

// TODO: have this on rubix-ui

// EdgeUploadApp upload an app to the edge device
func (inst *Client) EdgeUploadApp(hostIDName string, app *installer.Upload) (*installer.AppResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/upload")
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

// EdgeUploadService generate a service file and upload it to edge device
func (inst *Client) EdgeUploadService(hostIDName string, app *appstore.ServiceFile) (*appstore.UploadResponse, error) {
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
func (inst *Client) InstallEdgeService(hostIDName string, body *installer.Install) (*systemd.InstallResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/service/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&systemd.InstallResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemd.InstallResponse), nil
}

// EdgeUninstallApp remove/delete an app and its service
func (inst *Client) EdgeUninstallApp(hostIDName, appName string, deleteApp bool) (*systemd.UninstallResponse, error) {
	url := fmt.Sprintf("/api/edge/apps/?name=%s&delete=%s", appName, strconv.FormatBool(deleteApp))
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&systemd.UninstallResponse{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemd.UninstallResponse), nil
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

// EdgeAppStatus get all the apps with its status
// func (inst *Client) EdgeAppStatus(hostIDName string, ) ([]installer.AppsStatus, error) {
//
// }
