package appstore

import (
	"fmt"
)

type EdgeConfig struct {
	Data []byte `json:"data"`
	Path string `json:"path"`
}

func (inst *Store) EdgeReadConfig(hostUUID, hostName, appName, configName string) (*EdgeConfig, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	path := inst.App.GetAppConfigPath(appName)
	fileNamePath := fmt.Sprintf("%s/%s", path, configName)
	file, err := client.ReadFile(fileNamePath)
	if err != nil {
		return nil, err
	}
	return &EdgeConfig{
		Data: file,
		Path: fileNamePath,
	}, nil

}
