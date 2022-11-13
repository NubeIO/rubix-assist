package appstore

import (
	"errors"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/pkg/global"
	log "github.com/sirupsen/logrus"
	"path"
)

func (inst *Store) EdgeWriteConfig(hostUUID, hostName string, body *assistmodel.EdgeConfig) (*model.Message, error) {
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
	appDataConfigPath := global.App.GetAppDataConfigPath(body.AppName)
	exists, err := inst.EdgeDirExists(hostUUID, hostName, appDataConfigPath)
	if err != nil {
		return nil, err
	}
	if !exists {
		dir, err := inst.EdgeCreateDir(hostUUID, hostName, appDataConfigPath)
		if err != nil {
			return nil, err
		}
		log.Infof("made config dir as was not existing: %s", dir.Message)
	}
	absoluteAppDataConfigName := path.Join(appDataConfigPath, configName)

	writeFile := &assistmodel.WriteFile{
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

func (inst *Store) EdgeReadConfig(hostUUID, hostName, appName, configName string) (*assistmodel.EdgeConfigResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	appDataConfigPath := global.App.GetAppDataConfigPath(appName)
	absoluteAppDataConfigName := path.Join(appDataConfigPath, configName)
	file, err := client.ReadFile(absoluteAppDataConfigName)
	if err != nil {
		return nil, err
	}
	return &assistmodel.EdgeConfigResponse{
		Data:     file,
		FilePath: absoluteAppDataConfigName,
	}, nil
}
