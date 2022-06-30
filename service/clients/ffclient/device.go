package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// AddDevice an object
func (inst *FlowClient) AddDevice(device *model.Device) (*model.Device, error) {
	url := fmt.Sprintf("/api/devices")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Device{}).
		SetBody(device).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

// GetFirstDevice first object
func (inst *FlowClient) GetFirstDevice() (*model.Device, error) {
	devices, err := inst.GetDevices()
	if err != nil {
		return nil, err
	}

	for _, device := range *devices {
		return &device, err
	}
	return nil, err
}

// GetDevices all objects
func (inst *FlowClient) GetDevices() (*[]model.Device, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.Device{}).
		Get("/api/devices"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]model.Device), nil
}

// GetDevice an object
func (inst *FlowClient) GetDevice(uuid string) (*model.Device, error) {
	url := fmt.Sprintf("/api/devices/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Device{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

// EditDevice edit an object
func (inst *FlowClient) EditDevice(uuid string, device *model.Device) (*model.Device, error) {
	url := fmt.Sprintf("/api/devices/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.Device{}).
		SetBody(device).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}
