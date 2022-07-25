package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"io"
)

// UploadApp upload an app
func (inst *Client) UploadApp(hostIDName, appName, version, buildName, fileName string, reader io.Reader) (*installer.UploadResponse, error) {
	url := fmt.Sprintf("/api/apps/upload/?name=%s&buildName=%s&version=%s", appName, buildName, version)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.UploadResponse), nil
}

// InstallApp add/install a new an app (needs the build on the edge device)
/*
{
    "name": "flow-framework",
    "build_name": "flow-framework",
    "version": "v1.1",
    "source": "/data/tmp/tmp_DB18FE83463A/flow-framework-0.6.0-8655148f.amd64.zip"
}
*/
func (inst *Client) InstallApp(hostIDName string, body *installer.Install) (*installer.AppResponse, error) {
	url := fmt.Sprintf("/api/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&installer.AppResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.AppResponse), nil
}

// UploadServiceFile add/install a new an app service (service file needs to be needs the build on the edge device)
func (inst *Client) UploadServiceFile(appName, version, buildName, fileName string, reader io.Reader) (*installer.UploadResponse, error) {
	url := fmt.Sprintf("/api/apps/service/?name=%s&buildName=%s&version=%s", appName, buildName, version)
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
/*
{
    "name": "flow-framework",
    "build_name": "flow-framework",
    "version": "v0.6.0",
    "service_name": "nubeio-flow-framework.service",
    "source": "/data/tmp/tmp_F7CFBE2FA1E3/nubeio-flow-framework.service"
}
*/
func (inst *Client) InstallService(body *installer.Install) (*installer.InstallResp, error) {
	url := fmt.Sprintf("/api/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.InstallResp{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.InstallResp), nil

}
