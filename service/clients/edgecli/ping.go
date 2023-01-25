package edgecli

import (
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
)

func (inst *Client) Ping() (globalUUID *string, pingable bool, isValidToken bool) {
	url := "/api/system/device"
	resp, err := inst.Rest.R().
		SetResult(&rubixregistry.DeviceInfo{}).
		Get(url)
	if err != nil {
		return nil, false, false
	}
	if resp.StatusCode() == 401 {
		return nil, true, false
	}
	deviceInfo := resp.Result().(*rubixregistry.DeviceInfo)
	return &deviceInfo.GlobalUUID, true, true
}
