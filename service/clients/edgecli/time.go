package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-date/datelib"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/NubeIO/rubix-edge/service/system"
)

func (inst *Client) SystemTime() (*datelib.Time, error) {
	url := fmt.Sprintf("/api/time")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&datelib.Time{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*datelib.Time), nil
}

func (inst *Client) SetSystemTime(body system.DateBody) (*datelib.Time, error) {
	url := fmt.Sprintf("/api/time")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&datelib.Time{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*datelib.Time), nil
}

func (inst *Client) GetHardwareTZ() (string, error) {
	url := fmt.Sprintf("/api/timezone")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Get(url))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (inst *Client) GetTimeZoneList() ([]string, error) {
	url := fmt.Sprintf("/api/timezone/list")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]string{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]string)
	return *data, nil
}

func (inst *Client) UpdateTimezone(body system.DateBody) (*system.Message, error) {
	url := fmt.Sprintf("/api/timezone")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&system.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

//systemTimeZone := apiRoutes.Group("/timezone")
//{
//systemTimeZone.GET("/", api.GetHardwareTZ)
//systemTimeZone.POST("/", api.UpdateTimezone)
//systemTimeZone.GET("/list", api.GetTimeZoneList)
//systemTimeZone.POST("/config", api.GenerateTimeSyncConfig)
//}
