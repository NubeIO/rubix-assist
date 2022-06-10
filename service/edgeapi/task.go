package edgeapi

import (
	"fmt"
	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/autocli"
	"github.com/NubeIO/rubix-assist/service/tasks"
	automodel "github.com/NubeIO/rubix-automater/automater/model"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"github.com/NubeIO/rubix-automater/controller/pipectl"
)

/*
- post a new pipeline
- create
*/

type TaskParams struct {
	LocationName string `json:"locationName"`
	NetworkName  string `json:"network_name"`
	HostName     string `json:"host_name"`
	HostUUID     string `json:"host_uuid"`

	AppName string `json:"app_name"`
	Version string `json:"version"`
}

func (inst *Manager) PipeRunner(app *App) (*automodel.Pipeline, *autocli.Response) {
	return inst.pipeRunner(app)
}

func (inst *Manager) pipeRunner(app *App) (*automodel.Pipeline, *autocli.Response) {
	resp := &autocli.Response{}
	host, err, _ := inst.getHost(app)
	if err != nil {
		resp.Message = err
		return nil, resp
	}
	pingTask := &jobctl.JobBody{
		Name:       fmt.Sprintf("run %s task on host:%s", tasks.PingHost.String(), host.Name),
		TaskName:   tasks.PingHost.String(),
		TaskParams: map[string]interface{}{"url": host.IP, "port": host.RubixPort},
	}

	installTask := &jobctl.JobBody{
		Name:     fmt.Sprintf("run %s task on host:%s", tasks.InstallApp.String(), host.Name),
		TaskName: tasks.InstallApp.String(),
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
	pipe, resp := client.AddPipeline(pipeBuilder)

	if resp.StatusCode > 299 {
		return pipe, resp
	}
	pipe, resp = inst.taskEntry(pipe, resp)
	return pipe, resp
}

func (inst *Manager) taskEntry(data *automodel.Pipeline, response *autocli.Response) (*automodel.Pipeline, *autocli.Response) {
	task := &model.Task{
		IsPipeline:   true,
		PipelineUUID: data.UUID,
		Status:       data.Status.String(),
	}
	_, err := inst.DB.TaskEntry(task)
	if err != nil {
		response.Message = err
		return data, response
	}
	return data, response

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
