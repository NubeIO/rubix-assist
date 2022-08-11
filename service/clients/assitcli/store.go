package assitcli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"io"
)

// ListAppsWithVersions list apps with versions
func (inst *Client) ListAppsWithVersions() ([]appstore.ListApps, error) {
	url := fmt.Sprintf("/api/appstore/apps")
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
	url := fmt.Sprintf("/api/appstore/apps/details")
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
	url := fmt.Sprintf("/api/appstore/add/?name=%s&version=%s&product=%s&arch=%s", appName, version, product, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&appstore.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.UploadResponse), nil
}

// CheckStoreApp list apps and appstore
func (inst *Client) CheckStoreApp(appName, version string) (*[]appstore.App, error) {
	url := fmt.Sprintf("/api/appstore/check/app/?name=%s&version=%s", appName, version)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]appstore.App{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]appstore.App), nil
}

//func (inst *Controller) CheckStoreApp(c *gin.Context) {
//	m := &appstore.App{}
//	err = c.ShouldBindJSON(&m)
//	data, err := inst.Store.CheckApp(m)
//	if err != nil {
//		reposeHandler(data, err, c)
//		return
//	}
//	reposeHandler(data, err, c)
//}
//
//func (inst *Controller) CheckStoreApps(c *gin.Context) {
//	var m []appstore.App
//	err = c.ShouldBindJSON(&m)
//	data, err := inst.Store.CheckApps(m)
//	if err != nil {
//		reposeHandler(data, err, c)
//		return
//	}
//	reposeHandler(data, err, c)
//}
