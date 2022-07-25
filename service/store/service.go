package store

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"os"
	"strings"
)

func (inst *Store) setServiceName(appName string) string {
	return fmt.Sprintf("nubeio-%s", appName)
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
	AppName                 string `json:"app_name"`
	AppVersion              string `json:"app_version"`
	AppBuildName            string `json:"app_build_name"`
	ServiceDescription      string `json:"service_description"`
	RunAsUser               string `json:"run_as_user"`
	ServiceWorkingDirectory string `json:"service_working_directory"` // /data/rubix-service/apps/install/flow-framework/v0.6.1/
	AppSpecficExecStart     string `json:"app_specfic_exec_start"`    // WORKING-DIR/app -p 1660 -g /data/flow-framework -d data -prod
	CustomServiceExecStart  string `json:"custom_service_exec_start"` // npm run prod:start --prod --datadir /data/rubix-wires/data --envFile /data/rubix-wires/config/.env
}

// EdgeServiceInstall this assumes that the service file and app already exists on the edge device
func (inst *Store) EdgeServiceInstall(hostUUID, hostName string, body *installer.Install) (*installer.InstallResp, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.InstallService(body)
}

// EdgeServiceUpload this will generate and upload the service file to the edge device
func (inst *Store) EdgeServiceUpload(hostUUID, hostName string, app *ServiceFile) (*installer.UploadResponse, error) {
	tmpFilePath, err := inst.App.MakeTmpDirUpload()
	if err != nil {
		return nil, err
	}
	if app.AppName == "" {
		return nil, errors.New("app name can not be empty, try flow-framework")
	}
	serviceName := inst.setServiceName(app.AppName)
	appVersion := app.AppVersion
	if appVersion == "" {
		return nil, errors.New("app version can not be empty, try v0.6.0")
	}
	if err = checkVersion(appVersion); err != nil {
		return nil, err
	}
	appBuildName := app.AppBuildName
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
		execCmd = inst.setServiceExecStart(app.AppName, appVersion, execCmd)
		if err := inst.checkServiceExecStart(execCmd, appBuildName, appVersion); err != nil {
			return nil, err
		}
	}

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
		return nil, err
	}

	fileName := fmt.Sprintf("%s.service", serviceName)

	fileAndPath := filePath(fmt.Sprintf("%s/%s", tmpFilePath, fileName))
	reader, err := os.Open(fileAndPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open file:%s err:%s", fileAndPath, err.Error()))
	}

	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}

	return client.UploadServiceFile(app.AppName, appVersion, appBuildName, serviceName, reader)
}
