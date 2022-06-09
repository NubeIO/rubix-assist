package em

import (
	"errors"
	base "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-automater/automater/model"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"github.com/NubeIO/rubix-automater/controller/pipectl"
	autocli "github.com/NubeIO/rubix-automater/service/client"
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

// get api call to install an app
/*
- get the host id and token
- send a pipeline message to install the required apps
- automater will try to install the apps by calling RA install app api
*/

func (inst *EdgeManager) InstallAppPipeline(body *App) (data *model.Pipeline, response *autocli.Response) {
	_, err := inst.DB.GetHostByName(body.HostName, false)
	if err != nil {
		return nil, nil
	}
	autoData, autoResp := inst.AddPipeline(body)
	return autoData, autoResp
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

func (inst *EdgeManager) AddPipeline(app *App) (data *model.Pipeline, response *autocli.Response) {

	cli := autocli.New("0.0.0.0", 1663)

	jobOne := &jobctl.JobBody{
		Name:       "ping 1",
		TaskName:   "pingHost",
		Options:    &model.JobOptions{},
		TaskParams: map[string]interface{}{"url": "0.0.0.0", "port": 1662},
	}

	jobTwo := &jobctl.JobBody{
		Name:       "install app",
		TaskName:   "installApp",
		Options:    &model.JobOptions{},
		TaskParams: map[string]interface{}{"hostName": app.HostName, "appName": app.AppName, "version": app.Version},
	}

	var jobs []*jobctl.JobBody
	jobs = append(jobs, jobOne)
	jobs = append(jobs, jobTwo)

	body := &pipectl.PipelineBody{
		Name:       "install app pipeline",
		Jobs:       jobs,
		ScheduleAt: "10 sec",
		PipelineOptions: &model.PipelineOptions{
			EnableInterval:   false,
			RunOnInterval:    "10 sec",
			DelayBetweenTask: 1,
			CancelOnFailure:  false,
		},
	}

	return cli.AddPipeline(body)

}

func (inst *EdgeManager) InstallApp2(body *App) (*installer.InstallResponse, interface{}) {
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
		Token:   "tokens[0].Token",
		Version: body.Version,
	}
	data, resp := inst.reset(host.IP, host.RubixPort).InstallApp(app)
	if resp.GetStatus() > 299 {
		return data, resp.Message
	}
	return data, nil
}
