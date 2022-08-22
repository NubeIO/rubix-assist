package assistapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

func (inst *Client) FFProxyGET(hostIDName, path string) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/ff/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Get(path)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) ProxyGET(hostIDName, path string) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Get(path)
	if err != nil {
		return nil, err
	}
	return resp, err
}
