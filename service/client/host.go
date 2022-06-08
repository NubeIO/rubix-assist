package client

import (
	"encoding/json"
	"fmt"

	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/go-resty/resty/v2"
)

type Path struct {
	Path string
}

var Paths = struct {
	Hosts       Path
	Ping        Path
	HostNetwork Path
	Location    Path
	Users       Path
	Edge        Path
	Apps        Path
}{
	Hosts:       Path{Path: "/api/hosts"},
	Ping:        Path{Path: "/api/system/ping"},
	HostNetwork: Path{Path: "/api/networks"},
	Location:    Path{Path: "/api/locations"},
	Users:       Path{Path: "/api/locations"},
	Edge:        Path{Path: "/api/edge"},
	Apps:        Path{Path: "/api/edge/apps"},
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	resty      *resty.Response
}

func (response *Response) buildResponse(restyResp *resty.Response, err error) *Response {
	response.resty = restyResp
	var msg interface{}
	if err != nil {
		msg = err.Error()
	} else {
		asJson, err := response.AsJson()
		if err != nil {
			msg = restyResp.String()
		} else {
			msg = asJson
		}
	}
	response.StatusCode = restyResp.StatusCode()
	response.Message = msg
	return response
}

func (response *Response) AsString() string {
	return response.resty.String()
}

func (response *Response) GetError() interface{} {
	return response.resty.Error()
}

func (response *Response) GetMessage() interface{} {
	return response.Message
}

func (response *Response) GetStatus() int {
	return response.resty.StatusCode()
}

// AsJson return as body as blank interface
func (response *Response) AsJson() (interface{}, error) {
	var out interface{}
	err := json.Unmarshal(response.resty.Body(), &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (inst *Client) GetHosts() (data []model.Host, response *Response) {
	path := fmt.Sprintf(Paths.Hosts.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Host{}).
		Get(path)
	return *resp.Result().(*[]model.Host), response.buildResponse(resp, err)
}

func (inst *Client) GetHost(uuid string) (data *model.Host, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.Host{}).
		Get(path)
	return resp.Result().(*model.Host), response.buildResponse(resp, err)
}

func (inst *Client) AddHost(body *model.Host) (data *model.Host, response *Response) {
	path := fmt.Sprintf(Paths.Hosts.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Host{}).
		Post(path)
	return resp.Result().(*model.Host), response.buildResponse(resp, err)
}

func (inst *Client) UpdateHost(uuid string, body *model.Host) (data *model.Host, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Host{}).
		Patch(path)
	return resp.Result().(*model.Host), response.buildResponse(resp, err)
}

func (inst *Client) DeleteHost(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}
func (inst *Client) GetHostSchema() (data *model.HostSchema, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, "schema")
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.HostSchema{}).
		Get(path)
	return resp.Result().(*model.HostSchema), response.buildResponse(resp, err)
}
