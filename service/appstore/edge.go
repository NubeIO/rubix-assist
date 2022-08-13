package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
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

func (inst *Store) setServiceName(appName string) string {
	return fmt.Sprintf("nubeio-%s", appName)
}

func (inst *Store) setServiceFileName(appName string) string {
	return fmt.Sprintf("nubeio-%s.service", appName)
}

func (inst *Store) setServiceWorkingDir(appName, appVersion string) string {
	return inst.App.GetAppInstallPathAndVersion(appName, appVersion)
}

func (inst *Store) setServiceExecStart(appName, appVersion, AppSpecficExecStart string) string {
	workingDir := inst.App.GetAppInstallPathAndVersion(appName, appVersion)
	return fmt.Sprintf("%s/%s", workingDir, AppSpecficExecStart)
}

func (inst *Store) checkServiceExecStart(service, appName, appVersion string) error {
	if strings.Contains(service, inst.App.GetAppInstallPathAndVersion(appName, appVersion)) {
		return nil
	}
	return errors.New("ExecStart command is not matching appBuildName & appVersion")
}

type ServiceFile struct {
	Name                    string `json:"name"`
	Version                 string `json:"version"`
	ServiceDependency       string `json:"service_dependency"` // nodejs
	ServiceDescription      string `json:"service_description"`
	RunAsUser               string `json:"run_as_user"`
	ServiceWorkingDirectory string `json:"service_working_directory"` // /data/rubix-service/apps/install/flow-framework/v0.6.1/
	AppSpecficExecStart     string `json:"app_specfic_exec_start"`    // WORKING-DIR/app -p 1660 -g /data/flow-framework -d data -prod
	CustomServiceExecStart  string `json:"custom_service_exec_start"` // npm run prod:start --prod --datadir /data/rubix-wires/data --envFile /data/rubix-wires/config/.env
}

// InstallEdgeService this assumes that the service file and app already exists on the edge device
func (inst *Store) InstallEdgeService(hostUUID, hostName string, body *installer.Install) (*installer.InstallResp, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.InstallService(body)
}

func (inst *Store) GenerateUploadEdgeService(hostUUID, hostName string, app *ServiceFile) (*installer.UploadResponse, error) {
	resp, err := inst.generateUploadEdgeService(hostUUID, hostName, app)
	if err != nil {
		log.Errorf("generate service file hostUUID:%s, hostName:%s appName:%s", hostUUID, hostName, app.Name)
		log.Errorf("generate service file main err:%s", err.Error())
	}
	return resp, err
}

// GenerateUploadEdgeService this will generate and upload the service file to the edge device
func (inst *Store) generateUploadEdgeService(hostUUID, hostName string, app *ServiceFile) (*installer.UploadResponse, error) {
	tmpFilePath, err := inst.App.MakeTmpDirUpload()
	if err != nil {
		return nil, err
	}
	if app.Name == "" {
		return nil, errors.New("app name can not be empty, try flow-framework")
	}
	serviceName := inst.setServiceName(app.Name)
	serviceFileName := inst.setServiceFileName(app.Name)
	appVersion := app.Version
	if appVersion == "" {
		return nil, errors.New("app version can not be empty, try v0.6.0")
	}
	if err = checkVersion(appVersion); err != nil {
		return nil, err
	}
	appName := app.Name
	if appVersion == "" {
		return nil, errors.New("app build name can not be empty, try wires-builds")
	}
	workingDirectory := app.ServiceWorkingDirectory
	if workingDirectory == "" {
		workingDirectory = inst.setServiceWorkingDir(appName, appVersion)
	}
	user := app.RunAsUser
	if user == "" {
		user = "root"
	}
	execCmd := app.AppSpecficExecStart
	if app.CustomServiceExecStart != "" { // example use would be in wires
		execCmd = app.CustomServiceExecStart
	} else {
		if execCmd == "" {
			return nil, errors.New("app service ExecStart cant not be empty")
		}
		execCmd = inst.setServiceExecStart(app.Name, appVersion, execCmd)
		if err := inst.checkServiceExecStart(execCmd, appName, appVersion); err != nil {
			return nil, err
		}
	}
	log.Infof("generate service working dir: %s", workingDirectory)
	log.Infof("generate service execCmd: %s", execCmd)
	description := app.ServiceDescription
	bld := &builder.SystemDBuilder{
		ServiceName:      serviceName,
		Description:      description,
		User:             user,
		WorkingDirectory: workingDirectory,
		ExecStart:        execCmd,
		SyslogIdentifier: serviceName,
		WriteFile: builder.WriteFile{
			Write:    true,
			FileName: serviceName,
			Path:     tmpFilePath,
		},
	}
	err = bld.Build(os.FileMode(inst.Perm))
	if err != nil {
		log.Errorf("generate service file name:%s, err:%s", serviceName, err.Error())
		return nil, err
	}

	log.Infof("generate service file name:%s", serviceName)

	fileAndPath := filePath(fmt.Sprintf("%s/%s", tmpFilePath, serviceFileName))
	log.Infof("generate service file path:%s", fileAndPath)
	reader, err := os.Open(fileAndPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open file:%s err:%s", fileAndPath, err.Error()))
	}

	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}

	return client.UploadServiceFile(app.Name, appVersion, serviceFileName, reader)
}

func (inst *Store) EdgeUnInstallApp(hostUUID, hostName, appName, serviceName string, deleteApp bool) (*installer.RemoveRes, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.UnInstallApp(appName, serviceName, deleteApp)
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

func (inst *Store) EdgeCtlStatus(hostUUID, hostName string, body *installer.CtlBody) (*systemctl.SystemResponseChecks, error) {
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

func (inst *Store) EdgeServiceMassStatus(hostUUID, hostName string, body *installer.CtlBody) ([]systemctl.MassSystemResponseChecks, error) {
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

type EdgeUploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}

func (inst *Store) EdgeUploadLocalFile(hostUUID, hostName, path, fileName, destination string) (*EdgeUploadResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	resp, err := client.UploadLocalFile(path, fileName, destination)
	if err != nil {
		return nil, err
	}
	return &EdgeUploadResponse{
		Destination: resp.Destination,
		File:        resp.File,
		Size:        resp.Size,
		UploadTime:  resp.UploadTime,
	}, nil
}
