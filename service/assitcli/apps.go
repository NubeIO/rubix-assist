package assitcli

import (
	"fmt"
	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/NubeIO/rubix-assist/service/edgeapi"
)

func (inst *Client) InstallApp(body *edgeapi.App) (data *installer.InstallResponse, response *Response) {
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
