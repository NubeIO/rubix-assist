package appstore

import (
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
)

// EdgePing ping from the edge device
func (inst *Store) EdgePing(hostUUID, hostName string) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.Ping()
}

func (inst *Store) EdgeGetDeviceInfo(hostUUID, hostName string) (*rubixregistry.DeviceInfo, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeGetDeviceInfo()
}

func (inst *Store) EdgeProductInfo(hostUUID, hostName string) (*model.Product, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeProductInfo()
}
