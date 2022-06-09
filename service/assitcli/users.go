package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/model"
)

func (inst *Client) GetUsers() (data []model.User, response *Response) {
	path := fmt.Sprintf(Paths.Users.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.User{}).
		Get(path)
	return *resp.Result().(*[]model.User), response.buildResponse(resp, err)
}

func (inst *Client) AddUser(body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf(Paths.Users.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Post(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}

func (inst *Client) UpdateUser(uuid string, body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Users.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Patch(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}

func (inst *Client) DeleteUser(uuid string, body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Users.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Patch(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}
