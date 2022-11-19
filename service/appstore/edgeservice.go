package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/lib-systemctl-go/systemd"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/NubeIO/rubix-assist/service/systemctl"
	log "github.com/sirupsen/logrus"
	"os"
)

func (inst *Store) EdgeInstallService(hostUUID, hostName string, body *model.Install) (*systemd.InstallResponse, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.InstallService(body)
}

func (inst *Store) GenerateServiceFileAndEdgeUpload(hostUUID, hostName string, app *systemctl.ServiceFile) (*model.UploadResponse, error) {
	tmpDir, absoluteServiceFileName, err := systemctl.GenerateServiceFile(app, global.Installer)
	if err != nil {
		return nil, err
	}
	serviceName := global.Installer.GetServiceNameFromAppName(app.Name)
	reader, err := os.Open(absoluteServiceFileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open service_file: %s err: %s", absoluteServiceFileName, err.Error()))
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
