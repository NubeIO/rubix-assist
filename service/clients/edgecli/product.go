package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// EdgeProductInfo get edge product info
func (inst *Client) EdgeProductInfo() (*installer.Product, error) {
	url := fmt.Sprintf("/api/system/product")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&installer.Product{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*installer.Product), nil
}
