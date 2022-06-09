package client

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/em"
	"github.com/NubeIO/rubix-cli-app/service/apps/installer"
)

func (inst *Client) InstallApp(body *em.App) (data *installer.InstallResponse, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Apps.Path, "install")
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&installer.InstallResponse{}).
		SetError(&Response{}).
		Post(path)
	response = response.buildResponse(resp, err)
	if resp.IsSuccess() {
		data = resp.Result().(*installer.InstallResponse)
		response.Message = resp.Result().(*installer.InstallResponse)
	}
	return data, response
}
