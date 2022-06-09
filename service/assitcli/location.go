package assitcli

import (
	"fmt"

	"github.com/NubeIO/rubix-assist/pkg/model"
)

func (inst *Client) GetLocations() (data []model.Location, response *Response) {
	path := fmt.Sprintf(Paths.Location.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Location{}).
		Get(path)
	return *resp.Result().(*[]model.Location), response.buildResponse(resp, err)
}

func (inst *Client) GetLocation(uuid string) (data *model.Location, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Location.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.Location{}).
		Get(path)
	return resp.Result().(*model.Location), response.buildResponse(resp, err)
}

func (inst *Client) AddLocation(body *model.Location) (data *model.Location, response *Response) {
	path := fmt.Sprintf(Paths.Location.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Location{}).
		Post(path)
	return resp.Result().(*model.Location), response.buildResponse(resp, err)
}

func (inst *Client) UpdateLocation(uuid string, body *model.Location) (data *model.Location, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Location.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Location{}).
		Patch(path)
	return resp.Result().(*model.Location), response.buildResponse(resp, err)
}

func (inst *Client) DeleteLocation(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Location.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}

func (inst *Client) GetLocationSchema() (data *model.LocationSchema, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Location.Path, "schema")
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.LocationSchema{}).
		Get(path)
	return resp.Result().(*model.LocationSchema), response.buildResponse(resp, err)
}
