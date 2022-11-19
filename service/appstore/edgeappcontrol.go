package appstore

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/model"
)

func (inst *Store) EdgeSystemCtlAction(hostUUID, hostName string, body *model.SystemCtlBody) (*model.SystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeSystemCtlAction(body)
}

func (inst *Store) EdgeSystemCtlStatus(hostUUID, hostName string, body *model.SystemCtlBody) (*systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeSystemCtlStatus(body)
}

func (inst *Store) EdgeServiceMassAction(hostUUID, hostName string, body *model.SystemCtlBody) ([]model.MassSystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassAction(body)
}

func (inst *Store) EdgeServiceMassStatus(hostUUID, hostName string, body *model.SystemCtlBody) ([]systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassStatus(body)
}
