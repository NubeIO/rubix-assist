package assistapi

import (
	"fmt"
	"github.com/NubeIO/lib-date/datelib"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

func (inst *Client) EdgeSystemTime(hostIDName string) (*datelib.Time, error) {
	url := fmt.Sprintf("proxy/api/time/")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&datelib.Time{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*datelib.Time), nil
}
