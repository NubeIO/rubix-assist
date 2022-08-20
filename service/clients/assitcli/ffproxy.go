package assitcli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

func (inst *Client) FFProxyGET(hostIDName, path string) (*resty.Response, error) {
	path = fmt.Sprintf("ff/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Get(path)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) FFProxyPOST(hostIDName, path string, body interface{}) (*resty.Response, error) {
	path = fmt.Sprintf("ff/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	fmt.Println(path)
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Post(path)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) FFProxyPATCH(hostIDName, path string, body interface{}) (*resty.Response, error) {
	path = fmt.Sprintf("ff/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Patch(path)
	return resp, err
}

func (inst *Client) FFProxyPUT(hostIDName, path string, body interface{}) (*resty.Response, error) {
	path = fmt.Sprintf("ff/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Put(path)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) FFProxyDELETE(hostIDName, path string) (*resty.Response, error) {
	path = fmt.Sprintf("ff/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Delete(path)
	if err != nil {
		return nil, err
	}
	return resp, err
}
