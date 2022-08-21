package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-ufw/ufw"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-edge/service/system"
)

func (inst *Client) UWFStatusList() ([]ufw.UFWStatus, error) {
	url := fmt.Sprintf("/api/networking/firewall")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]ufw.UFWStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]ufw.UFWStatus)
	return *data, nil
}

func (inst *Client) UWFEnable() (*system.Message, error) {
	url := fmt.Sprintf("/api/networking/firewall/enable")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) UWFDisable() (*system.Message, error) {
	url := fmt.Sprintf("/api/networking/firewall/disable")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

//func (inst *System) UWFOpenPort(body UFWBody) (*ufw.Message, error) {
//	return inst.ufw.UWFOpenPort(body.Port)
//}

func (inst *Client) UWFOpenPort(body system.UFWBody) (*ufw.Message, error) {
	url := fmt.Sprintf("/api/networking/firewall/port/open")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&ufw.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ufw.Message), nil
}

func (inst *Client) UWFClosePort(body system.UFWBody) (*ufw.Message, error) {
	url := fmt.Sprintf("/api/networking/firewall/port/close")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&ufw.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ufw.Message), nil
}

//networkFirewall := apiRoutes.Group("/networking/firewall")
//{
//networkFirewall.GET("/", api.UWFStatusList)
//networkFirewall.GET("/status", api.UWFStatus)
//networkFirewall.GET("/active", api.UWFActive)
//networkFirewall.GET("/enable", api.UWFEnable)
//networkFirewall.GET("/disable", api.UWFDisable)
//networkFirewall.GET("/port/open", api.UWFOpenPort)
//networkFirewall.GET("/port/close", api.UWFClosePort)
//}
//
