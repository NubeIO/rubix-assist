package appstore

import (
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
)

func (inst *Store) EdgeFileExists(hostUUID, hostName, path string) (bool, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return false, err
	}
	return client.FileExists(path)
}

func (inst *Store) EdgeListFiles(hostUUID, hostName, path string) ([]string, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListFiles(path)
}

func (inst *Store) EdgeWalkFiles(hostUUID, hostName, path string) ([]string, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.Walk(path)
}

func (inst *Store) EdgeReadFile(hostUUID, hostName, path string) ([]byte, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ReadFile(path)
}

func (inst *Store) EdgeCreateFile(hostUUID, hostName string, body *assistmodel.WriteFile) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.CreateFile(body)
}

func (inst *Store) EdgeWriteString(hostUUID, hostName string, body *assistmodel.WriteFile) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.WriteString(body)
}

func (inst *Store) EdgeWriteFileJson(hostUUID, hostName string, body *assistmodel.WriteFile) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.WriteFileJson(body)
}

func (inst *Store) EdgeWriteFileYml(hostUUID, hostName string, body *assistmodel.WriteFile) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.WriteFileYml(body)
}

func (inst *Store) EdgeRenameFile(hostUUID, hostName, oldNameAndPath, newNameAndPath string) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.RenameFile(oldNameAndPath, newNameAndPath)
}

func (inst *Store) EdgeCopyFile(hostUUID, hostName, from, to string) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.CopyFile(from, to)
}

func (inst *Store) EdgeMoveFile(hostUUID, hostName, from, to string) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.MoveFile(from, to)
}

func (inst *Store) EdgeUploadLocalFile(hostUUID, hostName, file, destination string) (*assistmodel.EdgeUploadResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	resp, err := client.UploadLocalFile(file, destination)
	if err != nil {
		return nil, err
	}
	return &assistmodel.EdgeUploadResponse{
		Destination: resp.Destination,
		File:        resp.File,
		Size:        resp.Size,
		UploadTime:  resp.UploadTime,
	}, nil
}

func (inst *Store) EdgeDownloadFile(hostUUID, hostName, path, file, destination string) (*assistmodel.EdgeDownloadResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DownloadFile(path, file, destination)
}

func (inst *Store) EdgeDeleteFile(hostUUID, hostName, path string) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteFile(path)
}

func (inst *Store) EdgeDeleteAllFiles(hostUUID, hostName, path string) (*model.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteAllFiles(path)
}
