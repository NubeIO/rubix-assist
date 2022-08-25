package assistapi

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// EdgeReadFile read a files content
func (inst *Client) EdgeReadFile(path string) ([]byte, error) {
	url := fmt.Sprintf("proxy/api/files/read/?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
