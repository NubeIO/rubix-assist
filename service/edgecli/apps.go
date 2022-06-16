package edgecli

import (
	"fmt"
	"github.com/NubeIO/edge/service/apps"
	"github.com/NubeIO/edge/service/apps/installer"
)

type AppsResp struct {
	Code    int        `json:"code"`
	Message string     `json:"msg"`
	Data    []apps.App `json:"data"`
}

type AppResp struct {
	Code    int       `json:"code"`
	Message string    `json:"msg"`
	Data    *apps.App `json:"data"`
}

func (inst *Client) GetApps() (data []apps.App, response *Response) {
	path := fmt.Sprintf(paths.Apps.path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]apps.App{}).
		Get(path)
	return *resp.Result().(*[]apps.App), response.buildResponse(resp, err)
}

func (inst *Client) InstallApp(body *installer.App) (*installer.InstallResponse, *Response) {
	path := fmt.Sprintf(paths.Apps.path)
	response := &Response{}
	resp, _ := inst.Rest.R().
		SetBody(body).
		SetResult(&installer.InstallResponse{}).
		SetError(&installer.InstallResponse{}).
		Post(path)
	response.StatusCode = resp.StatusCode()
	if resp.IsSuccess() {
		response.Message = resp.Result().(*installer.InstallResponse)
		return resp.Result().(*installer.InstallResponse), response
	}
	return resp.Error().(*installer.InstallResponse), response
}

func (inst *Client) GetApp(uuid string) (data *AppResp, response *Response) {
	path := fmt.Sprintf("%s/%s", paths.Apps.path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&AppResp{}).
		Get(path)
	return resp.Result().(*AppResp), response.buildResponse(resp, err)
}

func (inst *Client) UpdateApp(uuid string, body *apps.App) (data *apps.App, response *Response) {
	path := fmt.Sprintf("%s/%s", paths.Apps.path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&apps.App{}).
		Patch(path)
	return resp.Result().(*apps.App), response.buildResponse(resp, err)
}

func (inst *Client) DeleteApp(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", paths.Apps.path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}
