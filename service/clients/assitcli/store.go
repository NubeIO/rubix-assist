package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-assist/service/store"
	"io"
)

// ListStore list apps and store
func (inst *Client) ListStore() ([]store.App, error) {
	url := fmt.Sprintf("/api/store")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]store.App{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]store.App), nil
}

// AddUploadStoreApp upload an app
func (inst *Client) AddUploadStoreApp(appName, version, product, arch, fileName string, reader io.Reader) (*store.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/add/?name=%s&version=%s&product=%s&arch=%s", appName, version, product, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&store.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*store.UploadResponse), nil
}

// CheckStoreApp list apps and store
func (inst *Client) CheckStoreApp(appName, version string) (*[]store.App, error) {
	url := fmt.Sprintf("/api/store/check/app/?name=%s&version=%s", appName, version)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]store.App{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]store.App), nil
}

//func (inst *Controller) CheckStoreApp(c *gin.Context) {
//	m := &store.App{}
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
//	var m []store.App
//	err = c.ShouldBindJSON(&m)
//	data, err := inst.Store.CheckApps(m)
//	if err != nil {
//		reposeHandler(data, err, c)
//		return
//	}
//	reposeHandler(data, err, c)
//}
