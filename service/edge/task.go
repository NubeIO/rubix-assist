package edge

import (
	"errors"
	"fmt"
	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/autocli"
	automodel "github.com/NubeIO/rubix-automater/automater/model"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"github.com/NubeIO/rubix-automater/controller/pipectl"
)

/*
- post a new pipeline
- create
*/

func (inst *Manager) getTokens() (token string, tokens []*model.Token, err error) {
	tokens = []*model.Token{}
	tokens, err = inst.DB.GetTokens()
	if err != nil {
		return "", nil, errors.New("no token provided")
	}
	if len(tokens) == 0 {
		return "", nil, errors.New("no token provided")
	}
	return tokens[0].Token, tokens, nil
}

// getHost returns the host and a GitHub token
func (inst *Manager) getHost(body *App) (*model.Host, error, string) {
	host, err := inst.DB.GetHostByLocationName(body.HostName, body.NetworkName, body.LocationName)
	if err != nil {
		return nil, err, ""
	}
	token, _, err := inst.getTokens()
	if err != nil {
		return nil, err, ""
	}
	return host, err, token
}

// installApp will install the app on the edge device
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

type TaskParams struct {
	LocationName string `json:"locationName"`
	NetworkName  string `json:"network_name"`
	HostName     string `json:"host_name"`
	HostUUID     string `json:"host_uuid"`

	AppName string `json:"app_name"`
	Version string `json:"version"`
}

func (inst *Manager) PipeRunner(app *App) (data *automodel.Pipeline, response *autocli.Response) {
	return inst.pipeRunner(app)
}

func (inst *Manager) pipeRunner(app *App) (data *automodel.Pipeline, response *autocli.Response) {
	response = &autocli.Response{}
	host, err, _ := inst.getHost(app)
	if err != nil {
		response.Message = err
		return nil, response
	}
	pingTask := &jobctl.JobBody{
		Name:       fmt.Sprintf("run %s task on host:%s", PingHost.String(), host.Name),
		TaskName:   PingHost.String(),
		TaskParams: map[string]interface{}{"url": host.IP, "port": host.RubixPort},
	}

	installTask := &jobctl.JobBody{
		Name:     fmt.Sprintf("run %s task on host:%s", InstallApp.String(), host.Name),
		TaskName: InstallApp.String(),
		Options: &automodel.JobOptions{
			EnableInterval: false,
			RunOnInterval:  "",
		},
		TaskParams: map[string]interface{}{"hostName": app.HostName, "networkName": app.NetworkName, "locationName": app.LocationName, "appName": app.AppName, "version": app.Version},
	}

	var jobs []*jobctl.JobBody
	jobs = append(jobs, pingTask)
	jobs = append(jobs, installTask)

	pipeBuilder := &pipectl.PipelineBody{
		Name:       "ping pipeline",
		Jobs:       jobs,
		ScheduleAt: "2 sec",
		PipelineOptions: &automodel.PipelineOptions{
			EnableInterval: false,
			RunOnInterval:  "10 sec",
		},
	}

	client := autocli.New("0.0.0.0", 1663)
	return client.AddPipeline(pipeBuilder)
}

func (inst *Manager) RunAppInstall(body *App) (*installer.InstallResponse, interface{}) {
	host, err, token := inst.getHost(body)
	if err != nil {
		return nil, nil
	}
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
