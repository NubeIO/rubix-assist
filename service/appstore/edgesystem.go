package appstore

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
)

// EdgePing ping from the edge device
func (inst *Store) EdgePing(hostUUID, hostName string, body *assistmodel.PingBody) (bool, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return false, err
	}
	return client.Ping(body)
}

func (inst *Store) EdgeGetDeviceInfo(hostUUID, hostName string) (*rubixregistry.DeviceInfo, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeGetDeviceInfo()
}

func (inst *Store) EdgeProductInfo(hostUUID, hostName string) (*installer.Product, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeProductInfo()
}
