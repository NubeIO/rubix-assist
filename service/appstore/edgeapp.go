package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"os"
)

// EdgeUploadApp uploads the build
func (inst *Store) EdgeUploadApp(hostUUID, hostName string, app *installer.Upload) (*installer.AppResponse, error) {
	if app.Name == "" {
		return nil, errors.New("upload app to edge: app name can not be empty")
	}
	if app.Version == "" {
		return nil, errors.New("upload app to edge: app version can not be empty")
	}
	if app.Product == "" {
		return nil, errors.New("upload app to edge: product type can not be empty, try RubixCompute, RubixComputeIO, RubixCompute5, Server, Edge28, Nuc")
	}
	if app.Arch == "" {
		return nil, errors.New("upload app to edge arch type can not be empty, try armv7 amd64")
	}
	path := inst.getAppsStoreAppWithVersionPath(app.Name, app.Version)
	buildDetails, err := inst.App.GetBuildZipNameByArch(path, app.Arch, app.DoNotValidateArch)
	if buildDetails == nil {
		return nil, errors.New(fmt.Sprintf("failed to match build zip name app:%s version:%s arch:%s", app.Name, app.Version, app.Arch))
	}
	fileAndPath := fmt.Sprintf("%s/%s", path, buildDetails.ZipName)
	reader, err := os.Open(fileAndPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open build for app:%s zip file_name:%s  err:%s", app.Name, buildDetails.ZipName, err.Error()))
	}
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.UploadApp(app, buildDetails.ZipName, reader)
}

func (inst *Store) EdgeUninstallApp(hostUUID, hostName, appName string, deleteApp bool) (*installer.UninstallResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeUninstallApp(appName, deleteApp)
}

func (inst *Store) EdgeListApps(hostUUID, hostName string) ([]installer.Apps, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListApps()
}

func (inst *Store) EdgeListAppsStatus(hostUUID, hostName string) ([]installer.AppsStatus, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListAppsStatus()
}
