package store

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/builder"
	log "github.com/sirupsen/logrus"
	"strings"

	"os"
)

type EdgeApp struct {
	Name              string `json:"name"`
	BuildName         string `json:"build_name"`
	Version           string `json:"version"`
	ServiceDependency string `json:"service_dependency"` // nodejs
}

// AddUploadEdgeApp
// upload the build
func (inst *Store) AddUploadEdgeApp(hostUUID, hostName string, app *EdgeApp) (*installer.AppResponse, error) {
	appName := app.Name
	version := app.Version
	buildName := app.BuildName
	if appName == "" {
		return nil, errors.New("app name can not be empty")
	}
	if buildName == "" {
		return nil, errors.New("app build name can not be empty")
	}
	if version == "" {
		return nil, errors.New("app version can not be empty")
	}
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
	return client.UploadApp(appName, version, buildName, fileName, reader)
}

func (inst *Store) setServiceName(appName string) string {
	return fmt.Sprintf("nubeio-%s", appName)
}

func (inst *Store) setServiceFileName(appName string) string {
	return fmt.Sprintf("nubeio-%s.service", appName)
}

func (inst *Store) setServiceWorkingDir(appBuildName, appVersion string) string {
	return inst.App.GetAppInstallPathAndVersion(appBuildName, appVersion)
}

func (inst *Store) setServiceExecStart(appBuildName, appVersion, AppSpecficExecStart string) string {
	workingDir := inst.App.GetAppInstallPathAndVersion(appBuildName, appVersion)
	return fmt.Sprintf("%s/%s", workingDir, AppSpecficExecStart)
}

func (inst *Store) checkServiceExecStart(service, appBuildName, appVersion string) error {
	if strings.Contains(service, inst.App.GetAppInstallPathAndVersion(appBuildName, appVersion)) {
		return nil
	}
	return errors.New("ExecStart command is not matching appBuildName & appVersion")
}

type ServiceFile struct {
	Name                    string `json:"name"`
	Version                 string `json:"version"`
	BuildName               string `json:"build_name"`
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
	appBuildName := app.BuildName
	if appVersion == "" {
		return nil, errors.New("app build name can not be empty, try wires-builds")
	}
	workingDirectory := app.ServiceWorkingDirectory
	if workingDirectory == "" {
		workingDirectory = inst.setServiceWorkingDir(appBuildName, appVersion)
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
		if err := inst.checkServiceExecStart(execCmd, appBuildName, appVersion); err != nil {
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

	return client.UploadServiceFile(app.Name, appVersion, appBuildName, serviceFileName, reader)
}
