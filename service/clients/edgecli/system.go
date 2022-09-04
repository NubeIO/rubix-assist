package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"strconv"
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

// EdgeGetDeviceInfo get edge device info
func (inst *Client) EdgeGetDeviceInfo() (*rubixregistry.DeviceInfo, error) {
	url := fmt.Sprintf("/api/system/device")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&rubixregistry.DeviceInfo{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*rubixregistry.DeviceInfo), nil
}

// Ping ping a edge device
func (inst *Client) Ping(body *assistmodel.PingBody) (bool, error) {
	url := fmt.Sprintf("/api/public/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
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
