package appstore

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

func (inst *Store) EdgeSystemCtlAction(hostUUID, hostName string, body *installer.SystemCtlBody) (*installer.SystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeSystemCtlAction(body)
}

func (inst *Store) EdgeSystemCtlStatus(hostUUID, hostName string, body *installer.SystemCtlBody) (*systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeSystemCtlStatus(body)
}

func (inst *Store) EdgeServiceMassAction(hostUUID, hostName string, body *installer.SystemCtlBody) ([]installer.MassSystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassAction(body)
}

func (inst *Store) EdgeServiceMassStatus(hostUUID, hostName string, body *installer.SystemCtlBody) ([]systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassStatus(body)
}
