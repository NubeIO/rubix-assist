package appstore

import (
	"github.com/NubeIO/rubix-assist/amodel"
)

func (inst *Store) EdgeUploadLocalFile(hostUUID, hostName, file, destination string) (*amodel.EdgeUploadResponse, error) {
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
	return &amodel.EdgeUploadResponse{
		Destination: resp.Destination,
		File:        resp.File,
		Size:        resp.Size,
		UploadTime:  resp.UploadTime,
	}, nil
}
