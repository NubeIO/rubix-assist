package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemd"
	"github.com/sergeymakinen/go-systemdconf/v2"
	"github.com/sergeymakinen/go-systemdconf/v2/unit"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

func (inst *Store) getServiceExecStart(appName, appVersion, appSpecificExecStart string) string {
	workingDir := inst.App.GetAppInstallPathWithVersionPath(appName, appVersion)
	return path.Join(workingDir, appSpecificExecStart)
}

func (inst *Store) checkServiceExecStart(service, appName, appVersion string) error {
	if strings.Contains(service, inst.App.GetAppInstallPathWithVersionPath(appName, appVersion)) {
		return nil
	}
	return errors.New(fmt.Sprintf("ExecStart command is not matching app_name: %s & app_version: %s", appName, appVersion))
}

type ServiceFile struct {
	Name                        string   `json:"name"`
	Version                     string   `json:"version"`
	ServiceDescription          string   `json:"service_description"`
	RunAsUser                   string   `json:"run_as_user"`
	ServiceWorkingDirectory     string   `json:"service_working_directory"`        // /data/rubix-service/apps/install/flow-framework/v0.6.1/
	ExecStart                   string   `json:"exec_start"`                       // app -p 1660 -g <data_dir> -d data -prod
	AttachWorkingDirOnExecStart bool     `json:"attach_working_dir_on_exec_start"` // true, false
	EnvironmentVars             []string `json:"environment_vars"`                 // Environment="g=/data/bacnet-server-c"
}

// EdgeInstallService this assumes that the service file and app already exists on the edge device
func (inst *Store) EdgeInstallService(hostUUID, hostName string, body *installer.Install) (*systemd.InstallResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.InstallService(body)
}

func (inst *Store) GenerateServiceFileAndEdgeUpload(hostUUID, hostName string, app *ServiceFile) (*installer.UploadResponse, error) {
	tmpDir, absoluteServiceFileName, err := inst.generateServiceFile(app)
	serviceName := inst.App.GetServiceNameFromAppName(app.Name)
	if err != nil {
		return nil, err
	}
	reader, err := os.Open(absoluteServiceFileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open service_file: %s err: %s", serviceName, err.Error()))
	}

	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	err = fileutils.RmRF(tmpDir)
	if err != nil {
		log.Errorf("delete tmp dir after generating service file %s", absoluteServiceFileName)
	}
	return client.UploadServiceFile(app.Name, app.Version, serviceName, reader)
}

func (inst *Store) generateServiceFile(app *ServiceFile) (tmpDir, absoluteServiceFileName string, err error) {
	tmpFilePath, err := inst.App.MakeTmpDirUpload()
	if err != nil {
		return "", "", err
	}
	if app.Name == "" {
		return "", "", errors.New("app name can not be empty, try flow-framework")
	}
	if app.Version == "" {
		return "", "", errors.New("app version can not be empty, try v0.6.0")
	}
	if err = installer.CheckVersion(app.Version); err != nil {
		return "", "", err
	}
	workingDirectory := app.ServiceWorkingDirectory
	if workingDirectory == "" {
		workingDirectory = inst.getAppWorkingDir(app.Name, app.Version)
	}
	log.Infof("generate service working dir: %s", workingDirectory)
	user := app.RunAsUser
	if user == "" {
		user = "root"
	}
	execCmd := app.ExecStart
	if app.AttachWorkingDirOnExecStart {
		execCmd = inst.getServiceExecStart(app.Name, app.Version, execCmd)
		if err := inst.checkServiceExecStart(execCmd, app.Name, app.Version); err != nil {
			return "", "", err
		}
	}
	if strings.Contains(execCmd, "<data_dir>") {
		execCmd = strings.ReplaceAll(execCmd, "<data_dir>", inst.App.GetAppDataPath(app.Name))
	}
	if strings.Contains(execCmd, "<data_dir_name>") {
		execCmd = strings.ReplaceAll(execCmd, "<data_dir_name>", inst.App.GetDataDirNameFromAppName(app.Name))
	}
	log.Infof("generate service exec_cmd: %s", execCmd)
	description := app.ServiceDescription
	if description == "" {
		description = fmt.Sprintf("NubeIO %s", app.Name)
	}
	var env systemdconf.Value
	for _, s := range app.EnvironmentVars {
		env = append(env, s)
	}
	service := unit.ServiceFile{
		Unit: unit.UnitSection{ // [Unit]
			Description: systemdconf.Value{description},
			After:       systemdconf.Value{"network.target"},
		},
		Service: unit.ServiceSection{ // [Service]
			ExecStartPre: nil,
			Type:         systemdconf.Value{"simple"},
			ExecOptions: unit.ExecOptions{
				User:             systemdconf.Value{user},
				WorkingDirectory: systemdconf.Value{workingDirectory},
				Environment:      env,
				StandardOutput:   systemdconf.Value{"syslog"},
				StandardError:    systemdconf.Value{"syslog"},
				SyslogIdentifier: systemdconf.Value{app.Name},
			},
			ExecStart: systemdconf.Value{
				execCmd,
			},
			Restart:    systemdconf.Value{"always"},
			RestartSec: systemdconf.Value{"10"},
		},
		Install: unit.InstallSection{ // [Install]
			WantedBy: systemdconf.Value{"multi-user.target"},
		},
	}
	b, _ := systemdconf.Marshal(service)
	serviceName := inst.App.GetServiceNameFromAppName(app.Name)
	absoluteServiceFileName = fmt.Sprintf("%s/%s", tmpFilePath, serviceName)
	err = fileutils.WriteFile(absoluteServiceFileName, string(b), os.FileMode(inst.App.FileMode))
	if err != nil {
		log.Errorf("write service file error %s", err.Error())
	}
	log.Infof("generate service file name: %s", serviceName)
	return tmpFilePath, absoluteServiceFileName, nil
}
