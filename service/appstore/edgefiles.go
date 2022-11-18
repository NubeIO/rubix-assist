package appstore

import (
	"github.com/NubeIO/rubix-assist/model"
)

func (inst *Store) EdgeUploadLocalFile(hostUUID, hostName, file, destination string) (*model.EdgeUploadResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	_, err = client.CreateDir(destination)
	if err != nil {
		return nil, err
	}
	resp, err := client.UploadLocalFile(file, destination)
	if err != nil {
		return nil, err
	}
	return &model.EdgeUploadResponse{
		Destination: resp.Destination,
		File:        resp.File,
		Size:        resp.Size,
		UploadTime:  resp.UploadTime,
	}, nil
}
