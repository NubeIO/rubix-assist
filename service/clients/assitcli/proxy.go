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
