package appstore

import (
	"github.com/NubeIO/rubix-assist/model"
)

func (inst *Store) EdgeListApps(hostUUID, hostName string) ([]model.Apps, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListApps()
}

func (inst *Store) EdgeListAppsStatus(hostUUID, hostName string) ([]model.AppsStatus, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListAppsStatus()
}
