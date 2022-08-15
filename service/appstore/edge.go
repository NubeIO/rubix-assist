package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"os"
)

type EdgeApp struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	Product           string `json:"product"`
	Arch              string `json:"arch"`
	ServiceDependency string `json:"service_dependency"` // nodejs
}

// AddUploadEdgeApp
// upload the build
func (inst *Store) AddUploadEdgeApp(hostUUID, hostName string, app *EdgeApp) (*installer.AppResponse, error) {
	appName := app.Name
	version := app.Version
	archType := app.Arch
	productType := app.Product
	if appName == "" {
		return nil, errors.New("upload app to edge app name can not be empty")
	}
	if version == "" {
		return nil, errors.New("upload app to edge  app version can not be empty")
	}
	if productType == "" {
		return nil, errors.New("upload app to edge  product type can not be empty, try RubixCompute, RubixComputeIO, RubixCompute5, Server, Edge28, Nuc")
	}
	if archType == "" {
		return nil, errors.New("upload app to edge arch type can not be empty, try armv7 amd64")
	}
	var fileName string
	path := inst.getAppStorePathAndVersion(appName, version)
	fileNames, err := inst.App.GetBuildZipNames(path)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed to find zip build err:%s", err.Error()))
		return nil, err
	}
	if len(fileNames) > 0 {
		fileName = fileNames[0].ZipName
	} else {
		err := errors.New(fmt.Sprintf("no zip builds found in path:%s", path))
		return nil, err
	}
	fileAndPath := filePath(fmt.Sprintf("%s/%s", path, fileName))
	reader, err := os.Open(fileAndPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open build for app:%s fileName:%s  err:%s", appName, fileName, err.Error()))
	}
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.UploadApp(appName, version, productType, archType, fileName, reader)
}

func (inst *Store) EdgeUnInstallApp(hostUUID, hostName, appName string, deleteApp bool) (*installer.RemoveRes, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeUnInstallApp(appName, deleteApp)
}

func (inst *Store) EdgeListApps(hostUUID, hostName string) ([]installer.Apps, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListApps()
}

func (inst *Store) EdgeListAppsAndService(hostUUID, hostName string) ([]installer.InstalledServices, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListAppsAndService()
}

func (inst *Store) EdgeListNubeServices(hostUUID, hostName string) ([]installer.InstalledServices, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListAppsAndService()
}

func (inst *Store) EdgeCtlAction(hostUUID, hostName string, body *installer.CtlBody) (*systemctl.SystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeCtlAction(body)
}

func (inst *Store) EdgeCtlStatus(hostUUID, hostName string, body *installer.CtlBody) (*systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeCtlStatus(body)
}

func (inst *Store) EdgeServiceMassAction(hostUUID, hostName string, body *installer.CtlBody) ([]systemctl.MassSystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassAction(body)
}

func (inst *Store) EdgeServiceMassStatus(hostUUID, hostName string, body *installer.CtlBody) ([]systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassStatus(body)
}

func (inst *Store) EdgeProductInfo(hostUUID, hostName string) (*installer.Product, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeProductInfo()
}
