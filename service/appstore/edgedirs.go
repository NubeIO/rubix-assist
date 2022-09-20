package appstore

import (
	"github.com/NubeIO/rubix-assist/model"
)

func (inst *Store) EdgeDirExists(hostUUID, hostName, path string) (bool, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return false, err
	}
	return client.DirExists(path)
}

func (inst *Store) EdgeCreateDir(hostUUID, hostName, path string) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.CreateDir(path)
}

func (inst *Store) EdgeDeleteFolder(hostUUID, hostName, path string, recursively bool) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteDir(path, recursively)
}
