package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	log "github.com/sirupsen/logrus"
)

func (inst *Store) EdgeWriteConfig(hostUUID, hostName string, body *assistmodel.EdgeConfig) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	if body.AppName == "" {
		return nil, errors.New("app_name can not be empty")
	}
	configName := body.ConfigType
	if configName == "" {
		configName = "config.yml"
	}
	appConfigPath := inst.App.GetAppConfigPath(body.AppName)
	exists, err := inst.EdgeDirExists(hostUUID, hostName, appConfigPath)
	if err != nil {
		return nil, err
	}
	if !exists {
		dir, err := inst.EdgeCreateDir(hostUUID, hostName, appConfigPath)
		if err != nil {
			return nil, err
		}
		log.Infof("made config dir as was not existing:%s", dir.Message)
	}
	absoluteConfigName := fmt.Sprintf("%s/%s", appConfigPath, configName)

	writeFile := &edgecli.WriteFile{
		FilePath:     absoluteConfigName,
		Body:         body.Body,
		BodyAsString: body.BodyAsString,
	}
	if configName == "config.yml" {
		return client.WriteFileYml(writeFile)
	} else if configName == ".env" {
		return client.WriteFile(writeFile)
	} else if configName == "config.json" {
		return client.WriteFileJson(writeFile)
	}

	return nil, errors.New("no valid config type, config.yml or .env or config.json")
}

func (inst *Store) EdgeReadConfig(hostUUID, hostName, appName, configName string) (*assistmodel.EdgeConfigResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	appConfigPath := inst.App.GetAppConfigPath(appName)
	absoluteConfigName := fmt.Sprintf("%s/%s", appConfigPath, configName)
	file, err := client.ReadFile(absoluteConfigName)
	if err != nil {
		return nil, err
	}
	return &assistmodel.EdgeConfigResponse{
		Data:     file,
		FilePath: absoluteConfigName,
	}, nil
}
