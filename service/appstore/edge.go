package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
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
	path := inst.getAppStorePathAndVersion(app.Name, app.Version)
	buildDetails, err := inst.App.GetBuildZipNameByArch(path, app.Arch, app.DoNotValidateArch)
	if buildDetails == nil {
		return nil, errors.New(fmt.Sprintf("failed to match build zip name app:%s version:%s arch:%s", app.Name, app.Version, app.Arch))
	}
	fileAndPath := filePath(fmt.Sprintf("%s/%s", path, buildDetails.ZipName))
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

func (inst *Store) EdgeUninstallApp(hostUUID, hostName, appName, serviceName string, deleteApp bool) (*installer.UninstallResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeUninstallApp(appName, serviceName, deleteApp)
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

func (inst *Store) EdgeListNubeServices(hostUUID, hostName string) ([]installer.InstalledServices, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.ListAppsStatus()
}

func (inst *Store) EdgeCtlAction(hostUUID, hostName string, body *installer.SystemCtlBody) (*systemctl.SystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeCtlAction(body)
}

func (inst *Store) EdgeCtlStatus(hostUUID, hostName string, body *installer.SystemCtlBody) (*systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeCtlStatus(body)
}

func (inst *Store) EdgeServiceMassAction(hostUUID, hostName string, body *installer.SystemCtlBody) ([]systemctl.MassSystemResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassAction(body)
}

func (inst *Store) EdgeServiceMassStatus(hostUUID, hostName string, body *installer.SystemCtlBody) ([]systemctl.SystemState, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeServiceMassStatus(body)
}
