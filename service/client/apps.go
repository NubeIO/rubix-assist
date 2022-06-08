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
		Post(path)
	return resp.Result().(*installer.InstallResponse), response.buildResponse(resp, err)
}
