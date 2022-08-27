package ffclient

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

func (inst *FlowClient) DeviceInfo() (*model.DeviceInfo, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.DeviceInfo{}).
		Get("/api/system/device_info"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.DeviceInfo), nil
}
