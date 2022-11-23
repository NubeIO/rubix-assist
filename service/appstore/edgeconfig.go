package appstore

import (
	"errors"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/pkg/global"
	log "github.com/sirupsen/logrus"
	"path"
)

func (inst *Store) EdgeWriteConfig(hostUUID, hostName string, body *amodel.EdgeConfig) (*amodel.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	if body.AppName == "" {
		return nil, errors.New("app_name can not be empty")
	}
	configName := body.ConfigName
	if configName == "" {
		configName = "config.yml"
	}
	appDataConfigPath := global.Installer.GetAppDataConfigPath(body.AppName)
	dirExistence, err := client.DirExists(appDataConfigPath)
	if err != nil {
		return nil, err
	}
	if !dirExistence.Exists {
		dir, err := client.CreateDir(appDataConfigPath)
		if err != nil {
			return nil, err
		}
		log.Infof("made config dir as was not existing: %s", dir.Message)
	}
	absoluteAppDataConfigName := path.Join(appDataConfigPath, configName)

	writeFile := &amodel.WriteFile{
		FilePath:     absoluteAppDataConfigName,
		Body:         body.Body,
		BodyAsString: body.BodyAsString,
	}
	if configName == "config.yml" {
		return client.WriteFileYml(writeFile)
	} else if configName == ".env" {
		return client.WriteString(writeFile)
	} else if configName == "config.json" {
		return client.WriteFileJson(writeFile)
	}

	return nil, errors.New("no valid config_name, config.yml or .env or config.json")
}

func (inst *Store) EdgeReadConfig(hostUUID, hostName, appName, configName string) (*amodel.EdgeConfigResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	appDataConfigPath := global.Installer.GetAppDataConfigPath(appName)
	absoluteAppDataConfigName := path.Join(appDataConfigPath, configName)
	file, err := client.ReadFile(absoluteAppDataConfigName)
	if err != nil {
		return nil, err
	}
	return &amodel.EdgeConfigResponse{
		Data:     file,
		FilePath: absoluteAppDataConfigName,
	}, nil
}
