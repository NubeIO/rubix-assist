package store

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"

	"os"
)

func checkApp(app *installer.Upload) error {

	return nil
}

type EdgeApp struct {
	AppName   string `json:"app_name"`
	BuildName string `json:"build_name"`
	Version   string `json:"version"`
}

// UploadEdgeApp
// upload the build
func (inst *Store) UploadEdgeApp(hostUUID, hostName string, app *EdgeApp) (*installer.UploadResponse, error) {
	appName := app.AppName
	version := app.Version
	buildName := app.BuildName
	fileName, path, _, err := inst.GetAppZipName(appName, version)
	if err != nil {
		return nil, err
	}
	fileAndPath := filePath(fmt.Sprintf("%s/%s", path, fileName))
	reader, err := os.Open(fileAndPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open file:%s err:%s", fileAndPath, err.Error()))
	}
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	host, err := inst.getHost(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.UploadApp(host.UUID, appName, version, buildName, fileName, reader)
}

// InstallEdgeApp
// make all the dirs and install the uploaded build
func (inst *Store) InstallEdgeApp(hostUUID, hostName string, body *installer.Install) (*installer.AppResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	host, err := inst.getHost(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.InstallApp(host.UUID, body)
}

// UploadEdgeService
// upload the service file
func (inst *Store) UploadEdgeService() {

}

func (inst *Store) InstallEdgeService() {

}
