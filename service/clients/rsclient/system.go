package rsclient

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// Ping Ping server
func (inst *Client) Ping() (*Ping, error) {
	url := fmt.Sprintf("/api/system/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Ping{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Ping), nil
}
