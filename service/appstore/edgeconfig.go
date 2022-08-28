package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	log "github.com/sirupsen/logrus"
)

func (inst *Store) EdgeWriteConfig(hostUUID, hostName string, body *EdgeConfig) (*edgecli.Message, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	appName := body.AppName
	if appName == "" {
		return nil, errors.New("app-name can not be empty")
	}
	configName := body.ConfigType
	if configName == "" {
		configName = "config.yml"
	}
	path := inst.App.GetAppConfigPath(appName)
	if path == "" {
		return nil, errors.New(fmt.Sprintf("not config path for for app:%s", appName))
	}
	exists, err := inst.EdgeDirExists(hostUUID, hostName, path)
	if err != nil {
		return nil, err
	}
	if !exists {
		dir, err := inst.EdgeCreateDir(hostUUID, hostName, path)
		if err != nil {
			return nil, err
		}
		log.Infof("made config dir as was not existing:%s", dir.Message)
	}
	fileNamePath := fmt.Sprintf("%s/%s", path, configName)

	writeBody := &edgecli.WriteFile{
		FilePath: fileNamePath,
		Body:     body.Body,
	}
	if configName == "config.yml" {
		return client.WriteFileYml(writeBody)
	} else if configName == ".env" {
		writeBody.BodyAsString = body.BodyAsString
		return client.WriteFile(writeBody)
	} else if configName == "config.json" {
		return client.WriteFileJson(writeBody)
	}

	return nil, errors.New("no valid config type, config.yml or .env or config.json")
}

type EdgeConfig struct {
	AppName      string      `json:"app_name,omitempty"`
	Body         interface{} `json:"body"` // used when writing data
	BodyAsString string      `json:"body_as_string"`
	Data         []byte      `json:"data"` // used for reading data
	Path         string      `json:"path,omitempty"`
	ConfigType   string      `json:"config_type,omitempty"` // config.yml
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
