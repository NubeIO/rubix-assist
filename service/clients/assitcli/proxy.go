package assitcli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (inst *Client) ProxyGET(hostID, path string) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	resp, err := inst.Rest.R().
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		Get(path)
	return resp, err
}
