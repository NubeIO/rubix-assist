package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"io"
)

// UploadApp upload an app
func (inst *Client) UploadApp(appName, version, buildName, productType, archType, fileName string, reader io.Reader) (*installer.AppResponse, error) {
	url := fmt.Sprintf("/api/apps/add/?name=%s&buildName=%s&version=%s&product=%s&arch=%s", appName, buildName, version, productType, archType)
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

// UploadServiceFile add/install a new an app service (service file needs to be needs the build on the edge device)
func (inst *Client) UploadServiceFile(appName, version, buildName, fileName string, reader io.Reader) (*installer.UploadResponse, error) {
	url := fmt.Sprintf("/api/apps/service/upload/?name=%s&buildName=%s&version=%s", appName, buildName, version)
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
