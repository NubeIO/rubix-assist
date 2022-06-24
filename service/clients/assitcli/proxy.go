package assitcli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

func (inst *Client) ProxyGET(hostID, path string) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		Get(path)
	return resp, err
}

func (inst *Client) ProxyPOST(hostID, path string, body interface{}) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		Post(path)
	return resp, err
}

func (inst *Client) ProxyPATCH(hostID, path string, body interface{}) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		Patch(path)
	return resp, err
}

func (inst *Client) ProxyPUT(hostID, path string, body interface{}) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		Put(path)
	return resp, err
}

func (inst *Client) ProxyDELETE(hostID, path string) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		Delete(path)
	return resp, err
}
