package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"strconv"
)

// AssistPing ping a device from this server
func (inst *Client) AssistPing(hostIDName string) (bool, error) {
	url := fmt.Sprintf("/api/public/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Post(url))
	if err != nil {
		return false, err
	}
	found, err := strconv.ParseBool(resp.String())
	if err != nil {
		return false, err
	}
	return found, nil
}
