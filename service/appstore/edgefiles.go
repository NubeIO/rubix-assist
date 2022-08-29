package appstore

import (
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
)

func (inst *Store) EdgeReadFile(hostUUID, hostName, path string) ([]byte, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ReadFile(path)
}

func (inst *Store) EdgeWriteFile(hostUUID, hostName string, body *edgecli.WriteFile) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.WriteFile(body)
}

func (inst *Store) EdgeWriteFileJson(hostUUID, hostName string, body *edgecli.WriteFile) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.WriteFileJson(body)
}

func (inst *Store) EdgeWriteFileYml(hostUUID, hostName string, body *edgecli.WriteFile) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.WriteFileYml(body)
}

func (inst *Store) EdgeFileExists(hostUUID, hostName, path string) (bool, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return false, err
	}
	return client.FileExists(path)
}

func (inst *Store) EdgeDirExists(hostUUID, hostName, path string) (bool, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return false, err
	}
	return client.DirExists(path)
}

func (inst *Store) EdgeCreateFile(hostUUID, hostName string, body *edgecli.WriteFile) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.CreateFile(body)
}
func (inst *Store) EdgeCreateDir(hostUUID, hostName, path string) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.CreateDir(path)
}

type EdgeUploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}

func (inst *Store) EdgeUploadLocalFile(hostUUID, hostName, path, fileName, destination string) (*EdgeUploadResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	resp, err := client.UploadLocalFile(path, fileName, destination)
	if err != nil {
		return nil, err
	}
	return &EdgeUploadResponse{
		Destination: resp.Destination,
		File:        resp.File,
		Size:        resp.Size,
		UploadTime:  resp.UploadTime,
	}, nil
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

func (inst *Store) EdgeRenameFile(hostUUID, hostName, oldNameAndPath, newNameAndPath string) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.RenameFile(oldNameAndPath, newNameAndPath)
}

func (inst *Store) EdgeCopyFile(hostUUID, hostName, from, to string) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.CopyFile(from, to)
}

func (inst *Store) EdgeMoveFile(hostUUID, hostName, from, to string) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.MoveFile(from, to)
}

func (inst *Store) EdgeDeleteFile(hostUUID, hostName, path string) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteFile(path)
}

func (inst *Store) EdgeDeleteAllFiles(hostUUID, hostName, path string) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteAllFiles(path)
}

func (inst *Store) EdgeDeleteFolder(hostUUID, hostName, path string, recursively bool) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteDir(path, recursively)
}

func (inst *Store) EdgeDownloadFile(hostUUID, hostName, path, file, destination string) (*edgecli.DownloadResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DownloadFile(path, file, destination)

}
