package store

import (
	"errors"
	"fmt"
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
	if strings.Contains(service, inst.setServiceWorkingDir(appBuildName, appVersion)) {
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

func (inst *Store) generateServiceFile(app *ServiceFile, tmpFilePath string) error {
	if tmpFilePath == "" {
		return errors.New("tmp path for service file can not be empty, try /data/tmp")
	}
	if app.AppName == "" {
		return errors.New("app name can not be empty, try flow-framework")
	}
	appName := inst.setServiceName(app.AppName)
	appVersion := app.AppVersion
	if appVersion == "" {
		return errors.New("app version can not be empty, try v0.6.0")
	}
	if err := checkVersion(appVersion); err != nil {
		return err
	}
	appBuildName := app.AppBuildName
	if appVersion == "" {
		return errors.New("app build name can not be empty, try wires-builds")
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
	if execCmd == "" {
		return errors.New("app service ExecStart cant not be empty")
	}
	execCmd = inst.setServiceExecStart(appName, appVersion, execCmd)
	if err := inst.checkServiceExecStart(execCmd, appBuildName, appVersion); err != nil {
		return err
	}
	description := app.ServiceDescription
	bld := &builder.SystemDBuilder{
		ServiceName:      appName,
		Description:      description,
		User:             user,
		WorkingDirectory: workingDirectory,
		ExecStart:        execCmd,
		SyslogIdentifier: appName,
		WriteFile: builder.WriteFile{
			Write:    true,
			FileName: appName,
			Path:     tmpFilePath,
		},
	}
	err := bld.Build(os.FileMode(inst.Perm))
	if err != nil {
		return err
	}
	return nil

}
