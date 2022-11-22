package appstore

import (
	"github.com/NubeIO/rubix-assist/model"
)

func (inst *Store) EdgeListAppsStatus(hostUUID, hostName string) ([]model.AppsStatus, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListAppsStatus()
}
