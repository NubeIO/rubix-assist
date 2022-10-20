package assistcli

import (
	"fmt"
	"github.com/NubeIO/lib-date/datelib"
)

type Message struct {
	Message interface{} `json:"message"`
}

type Time struct {
	*datelib.Time
}

func (inst *Client) GetTime() (*Time, *Message, error) {
	path := fmt.Sprintf("%s/%s", Paths.System.Path, "time")
	resp, err := inst.Rest.R().
		SetResult(&Time{}).
		SetError(&Message{}).
		Get(path)
	if resp.IsError() || err != nil {
		return nil, resp.Error().(*Message), err
	} else {
		return resp.Result().(*Time), nil, nil
	}
}

type NetworkingInfo struct {
	Interface     string `json:"interface"`
	Ip            string `json:"ip"`
	IpAndMask     string `json:"ip_and_mask"`
	Netmask       string `json:"netmask"`
	NetMaskLength int    `json:"net_mask_length"`
	Gateway       string `json:"gateway"`
	MacAddress    string `json:"mac_address"`
}

func (inst *Client) GetNetworking() (*[]NetworkingInfo, *Message, error) {
	path := fmt.Sprintf("%s/%s", Paths.Networking.Path, "networks")
	resp, err := inst.Rest.R().
		SetResult(&[]NetworkingInfo{}).
		SetError(&Message{}).
		Get(path)
	if resp.IsError() || err != nil {
		return nil, resp.Error().(*Message), err
	} else {
		return resp.Result().(*[]NetworkingInfo), nil, nil
	}
}

type InterfaceNames struct {
	InterfaceNames []string `json:"interface_names"`
}

func (inst *Client) GetInterfaces() (*InterfaceNames, *Message, error) {
	path := fmt.Sprintf("%s/%s", Paths.Networking.Path, "interfaces")
	resp, err := inst.Rest.R().
		SetResult(&InterfaceNames{}).
		SetError(&Message{}).
		Get(path)
	if resp.IsError() || err != nil {
		return nil, resp.Error().(*Message), err
	} else {
		return resp.Result().(*InterfaceNames), nil, nil
	}
}

type InternetIP struct {
	IpAddress string `json:"ip_address"`
	Ok        bool   `json:"ok"`
}

func (inst *Client) GetInternetIP() (*InternetIP, *Message, error) {
	path := fmt.Sprintf("%s/%s", Paths.Networking.Path, "internet")
	resp, err := inst.Rest.R().
		SetResult(&InternetIP{}).
		SetError(&Message{}).
		Get(path)
	if resp.IsError() || err != nil {
		return nil, resp.Error().(*Message), err
	} else {
		return resp.Result().(*InternetIP), nil, nil
	}
}
