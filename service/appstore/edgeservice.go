package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/sergeymakinen/go-systemdconf/v2"
	"github.com/sergeymakinen/go-systemdconf/v2/unit"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

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
	return errors.New(fmt.Sprintf("ExecStart command is not matching appName:%sappName & appVersion:%s", appName, appVersion))
}

type ServiceFile struct {
	Name                    string   `json:"name"`
	Version                 string   `json:"version"`
	ServiceDependency       string   `json:"service_dependency"` // nodejs
	ServiceDescription      string   `json:"service_description"`
	RunAsUser               string   `json:"run_as_user"`
	ServiceWorkingDirectory string   `json:"service_working_directory"` // /data/rubix-service/apps/install/flow-framework/v0.6.1/
	AppSpecficExecStart     string   `json:"app_specfic_exec_start"`    // WORKING-DIR/app -p 1660 -g /data/flow-framework -d data -prod
	CustomServiceExecStart  string   `json:"custom_service_exec_start"` // npm run prod:start --prod --datadir /data/rubix-wires/data --envFile /data/rubix-wires/config/.env
	EnvironmentVars         []string `json:"environment_vars"`          // Environment="g=/data/bacnet-server-c"
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

func (inst *Store) generateServiceFile(app *ServiceFile) (tmpDir, serviceFile, fileAndPath string, err error) {
	tmpFilePath, err := inst.App.MakeTmpDirUpload()
	if err != nil {
		return "", "", "", err
	}
	if app.Name == "" {
		return "", "", "", errors.New("app name can not be empty, try flow-framework")
	}
	serviceName := inst.setServiceName(app.Name)
	serviceFileName := inst.setServiceFileName(app.Name)
	appVersion := app.Version
	if appVersion == "" {
		return "", "", "", errors.New("app version can not be empty, try v0.6.0")
	}
	if err = checkVersion(appVersion); err != nil {
		return "", "", "", err
	}
	appName := app.Name
	if appVersion == "" {
		return "", "", "", errors.New("app build name can not be empty, try wires-builds")
	}
	workingDirectory := app.ServiceWorkingDirectory
	if workingDirectory == "" {
		workingDirectory = inst.setServiceWorkingDir(appName, appVersion)
	}
	log.Infof("generate service working dir: %s", workingDirectory)
	user := app.RunAsUser
	if user == "" {
		user = "root"
	}
	execCmd := app.AppSpecficExecStart
	if app.CustomServiceExecStart != "" { // example use would be in wires
		execCmd = app.CustomServiceExecStart
	} else {
		if execCmd == "" {
			return "", "", "", errors.New("app service ExecStart cant not be empty")
		}
		execCmd = inst.setServiceExecStart(app.Name, appVersion, execCmd)
		if err := inst.checkServiceExecStart(execCmd, appName, appVersion); err != nil {
			return "", "", "", err
		}
	}
	log.Infof("generate service execCmd: %s", execCmd)
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
	servicePath := fmt.Sprintf("%s/%s", tmpFilePath, serviceFileName)
	err = fileutils.WriteFile(servicePath, string(b), os.FileMode(FilePerm))
	if err != nil {
		log.Errorf("write service file error %s", err.Error())
	}
	log.Infof("generate service file name:%s", serviceName)
	return tmpFilePath, serviceFileName, servicePath, nil
}

// GenerateUploadEdgeService this will generate and upload the service file to the edge device
func (inst *Store) generateUploadEdgeService(hostUUID, hostName string, app *ServiceFile) (*installer.UploadResponse, error) {
	tmpDir, serviceFile, fileAndPath, err := inst.generateServiceFile(app)
	if err != nil {
		return nil, err
	}
	reader, err := os.Open(fileAndPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open service-file:%s err:%s", serviceFile, err.Error()))
	}

	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	err = fileutils.RmRF(tmpDir)
	if err != nil {
		log.Errorf("assist: delete tmp dir after generating service file%s", fileAndPath)
	}
	return client.UploadServiceFile(app.Name, app.Version, serviceFile, reader)
}
