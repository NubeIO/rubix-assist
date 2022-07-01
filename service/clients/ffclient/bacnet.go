package ffclient

import (
	"fmt"
	"github.com/NubeDev/bacnet"
	"github.com/NubeDev/bacnet/btypes"
	"github.com/NubeDev/bacnet/network"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

const bacnetMaster = "bacnetmaster"

// BacnetWhoIs do a whois on an existing network
func (inst *FlowClient) BacnetWhoIs(body *bacnet.WhoIsOpts, networkUUID string, addDevices bool) (*[]btypes.Device, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/whois/%s?add_devices=%t", bacnetMaster, networkUUID, addDevices)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetBody(body).
		SetResult(&[]btypes.Device{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]btypes.Device), nil
}

// BacnetDevicePoints get points from an added device
func (inst *FlowClient) BacnetDevicePoints(deviceUUID string, addPoints bool) (*[]network.PointDetails, error) {
	url := fmt.Sprintf("/api/plugins/api/%s/device/points/%s?add_points=%t", bacnetMaster, deviceUUID, addPoints)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]network.PointDetails{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]network.PointDetails), nil
}