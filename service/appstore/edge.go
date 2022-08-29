package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
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
	var dontCheckArch bool
	if appName == rubixWires {
		dontCheckArch = true
	}
	path := inst.getAppStorePathAndVersion(appName, version)
	buildDetails, err := inst.App.GetBuildZipNameByArch(path, archType, dontCheckArch)
	if buildDetails == nil {
		return nil, errors.New(fmt.Sprintf("failed to match build zip name app:%s version:%s arch:%s", appName, version, archType))
	}
	var fileName = buildDetails.ZipName
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

func (inst *Store) EdgeProductInfo(hostUUID, hostName string) (*installer.Product, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgeProductInfo()
}

func (inst *Store) EdgePublicInfo(hostUUID, hostName string) (*edgecli.DeviceProduct, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.EdgePublicInfo()
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
