package edgeapi

import (
	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/NubeIO/rubix-assist/pkg/model"
)

// installApp will install the app on the edgeapi device
func (inst *Manager) installApp(body *App, host *model.Host, token string) (*installer.InstallResponse, interface{}) {
	app := &installer.App{
		AppName: body.AppName,
		Token:   token,
		Version: body.Version,
	}
	data, resp := inst.reset(host.IP, host.RubixPort).InstallApp(app)
	if resp.StatusCode > 299 {
		return data, resp.Message
	}
	return data, nil
}
