package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/namings"
	"github.com/NubeIO/rubix-assist/pkg/constants"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	log "github.com/sirupsen/logrus"
	"path"
)

func (inst *Client) AppUninstall(appName string) (*amodel.Message, error) {
	installLocation := global.Installer.GetAppInstallPath(appName)
	url := fmt.Sprintf("/api/files/delete-all?path=%s", installLocation)
	if _, err := nresty.FormatRestyResponse(inst.Rest.R().Delete(url)); err != nil {
		log.Error(err)
	}
	_, _ = inst.uninstallServiceFile(appName)
	return &amodel.Message{Message: "successfully uninstalled the app"}, nil
}

func (inst *Client) uninstallServiceFile(appName string) (*amodel.Message, error) {
	serviceFileName := namings.GetServiceNameFromAppName(appName)
	serviceFile := path.Join(constants.ServiceDir, serviceFileName)
	symlinkServiceFile := path.Join(constants.ServiceDirSoftLink, serviceFileName)

	var err error
	url := fmt.Sprintf("/api/systemctl/disable?unit=%s", serviceFileName)
	if _, err := nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("disabled service %s", serviceFileName)

	url = fmt.Sprintf("/api/systemctl/stop?unit=%s", serviceFileName)
	if _, err := nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("stopped service %s", serviceFileName)

	url = fmt.Sprintf("/api/syscall/unlink?path=%s", symlinkServiceFile)
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("soft un-linked %s", symlinkServiceFile)

	url = "/api/systemctl/daemon-reload"
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("daemon reloaded")

	url = fmt.Sprintf("/api/files/delete-all?path=%s", serviceFile)
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Delete(url)); err != nil {
		log.Error(err)
	}
	log.Infof("deleted file %s", serviceFile)
	return nil, nil
}
