package autocli

import (
	"fmt"
	automodel "github.com/NubeIO/rubix-automater/automater/model"
	"github.com/NubeIO/rubix-automater/controller/pipectl"
)

func (inst *Client) AddPipeline(body *pipectl.PipelineBody) (data *automodel.Pipeline, response *Response) {
	path := fmt.Sprintf(Paths.Pipeline.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&automodel.Pipeline{}).
		SetError(&Response{}).
		Post(path)
	response = response.buildResponse(resp, err)
	if resp.IsSuccess() {
		data = resp.Result().(*automodel.Pipeline)
		response.Message = data
	}
	return data, response
}
