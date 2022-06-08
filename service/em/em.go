package em

import (
	"errors"
	base "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-cli-app/service/apps/installer"
	"github.com/NubeIO/rubix-cli-app/service/client"
)

type EdgeManager struct {
	em *client.Client
	DB *base.DB
}

func New(apps *EdgeManager) *EdgeManager {
	return apps
}

func (inst *EdgeManager) reset(url string, port int) *client.Client {
	return client.New(url, port)
}

type App struct {
	HostUUID string `json:"host_uuid"`
	HostName string `json:"host_name"`
	AppName  string `json:"app_name"`
	Version  string `json:"version"`
}

func (inst *EdgeManager) response() *client.Response {
	return &client.Response{}
}

func (inst *EdgeManager) InstallApp(body *App) (*installer.InstallResponse, interface{}) {
	host, err := inst.DB.GetHostByName(body.HostName, false)
	if err != nil {
		return nil, nil
	}
	tokens, err := inst.DB.GetTokens()
	if err != nil {
		return nil, err
	}
	if len(tokens) == 0 {
		return nil, errors.New("no token provided")
	}
	app := &installer.App{
		AppName: body.AppName,
		Token:   tokens[0].Token,
		Version: body.Version,
	}
	data, resp := inst.reset(host.IP, host.RubixPort).InstallApp(app)
	if resp.GetStatus() > 299 {
		return data, resp.Message
	}
	return data, nil
}
