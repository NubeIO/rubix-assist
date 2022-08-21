package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

func (inst *Client) RestartNetworking() (string, error) {
	url := fmt.Sprintf("/api/networking/networks/restart")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Post(url))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
