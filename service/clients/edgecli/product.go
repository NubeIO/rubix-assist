package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"github.com/NubeIO/rubix-edge/pkg/model"
	"strconv"
	"time"
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

type DeviceProduct struct {
	Device     *model.DeviceInfo  `json:"device"`
	Product    *installer.Product `json:"product"`
	Networking []networking.NetworkInterfaces
}

// EdgePublicInfo get edge product info
func (inst *Client) EdgePublicInfo() (*DeviceProduct, error) {
	url := fmt.Sprintf("/api/public/device")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&DeviceProduct{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*DeviceProduct), nil
}

type PingBody struct {
	Ip      string        `json:"ip"`
	Port    int           `json:"port"`
	TimeOut time.Duration `json:"time_out"`
}

// Ping ping a device
func (inst *Client) Ping(body *PingBody) (bool, error) {
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
