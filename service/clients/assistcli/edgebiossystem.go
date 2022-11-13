package assistcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/assistcli/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) EdgeBiosPing(hostIDName string) (*model.Message, error) {
	url := fmt.Sprintf("/proxy/eb/api/system/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

func (inst *Client) EdgeBiosDeviceType(hostIDName string) (*amodel.DeviceType, error) {
	url := fmt.Sprintf("/proxy/eb/api/system/device-type")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&amodel.DeviceType{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.DeviceType), nil
}
