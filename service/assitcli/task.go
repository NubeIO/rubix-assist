package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/model"
)

func (inst *Client) GetTasks() (data []model.Task, response *Response) {
	path := fmt.Sprintf(Paths.Tasks.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Task{}).
		Get(path)
	return *resp.Result().(*[]model.Task), response.buildResponse(resp, err)
}

func (inst *Client) GetTask(uuid string) (data *model.Task, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Tasks.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.Task{}).
		Get(path)
	return resp.Result().(*model.Task), response.buildResponse(resp, err)
}

func (inst *Client) AddTask(body *model.Task) (data *model.Task, response *Response) {
	path := fmt.Sprintf(Paths.Tasks.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Task{}).
		Post(path)
	return resp.Result().(*model.Task), response.buildResponse(resp, err)
}

func (inst *Client) UpdateTask(uuid string, body *model.Task) (data *model.Task, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Tasks.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Task{}).
		Patch(path)
	return resp.Result().(*model.Task), response.buildResponse(resp, err)
}

func (inst *Client) DeleteTask(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Tasks.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}
