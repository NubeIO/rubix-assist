package edgeapi

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/autocli"
	"github.com/NubeIO/rubix-assist/service/tasks"
	automodel "github.com/NubeIO/rubix-automater/automater/model"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"github.com/NubeIO/rubix-automater/controller/pipectl"
)

type AppTask struct {
	LocationName    string   `json:"locationName"`
	NetworkName     string   `json:"networkName"`
	HostName        string   `json:"hostName"`
	HostUUID        string   `json:"hostUUID"`
	AppName         string   `json:"appName"`
	SubTask         string   `json:"subTask"`
	Version         string   `json:"version"`
	OrderedTaskList []string `json:"orderedTaskList"`
}

func (inst *Manager) TaskBuilder(appTask *AppTask) (*automodel.Pipeline, *autocli.Response) {
	resp := &autocli.Response{}
	app := &App{
		LocationName: appTask.LocationName,
		NetworkName:  appTask.NetworkName,
		HostName:     appTask.HostName,
		HostUUID:     appTask.HostUUID,
		AppName:      appTask.AppName,
		Version:      appTask.Version,
	}
	host, err, _ := inst.getHost(app)
	if err != nil {
		resp.StatusCode = 404
		resp.Message = err.Error()
		return nil, resp
	}

	var jobs []*jobctl.JobBody
	for _, task := range appTask.OrderedTaskList {
		switch task {
		case tasks.PingHost.String():
			jobs = append(jobs, buildPing(appTask, host))
		case tasks.InstallApp.String():
			jobs = append(jobs, buildInstall(appTask, host))
		case tasks.InstallPlugin.String():

		case tasks.RemoveApp.String():

		case tasks.StopApp.String():

		case tasks.StartApp.String():

		}

	}
	pipeBuilder := &pipectl.PipelineBody{
		Name:       "ping pipeline",
		Jobs:       jobs,
		ScheduleAt: "2 sec",
		PipelineOptions: &automodel.PipelineOptions{
			EnableInterval: false,
			RunOnInterval:  "10 sec",
		},
	}
	cli := autocli.New("0.0.0.0", 1663)
	pipe, resp := cli.AddPipeline(pipeBuilder)
	if resp.StatusCode > 299 {
		return pipe, resp
	}
	err = inst.taskEntry(host, pipe)
	if err != nil {
		resp.Message = err.Error()
		return pipe, resp
	}
	return pipe, resp
}

func buildPing(appTask *AppTask, host *model.Host) *jobctl.JobBody {
	taskName := tasks.SubTask.String()
	subTaskName := tasks.PingHost.String()
	taskParams := map[string]interface{}{
		"locationName": appTask.LocationName,
		"networkName":  appTask.NetworkName,
		"hostName":     host.Name,
		"hostUUID":     host.UUID,
		"appName":      taskName,
		"subTask":      subTaskName,
	}
	task := &jobctl.JobBody{
		Name:        fmt.Sprintf("run %s task on host:%s", taskName, host.Name),
		TaskName:    taskName,
		SubTaskName: subTaskName,
		TaskParams:  taskParams,
	}
	return task
}

func buildInstall(appTask *AppTask, host *model.Host) *jobctl.JobBody {
	taskName := tasks.PingHost.String()
	subTaskName := tasks.PingHost.String()
	taskParams := map[string]interface{}{
		"locationName": appTask.LocationName,
		"networkName":  appTask.NetworkName,
		"hostName":     host.Name,
		"hostUUID":     host.UUID,
		"appName":      taskName,
		"subTask":      subTaskName,
	}
	task := &jobctl.JobBody{
		Name:        fmt.Sprintf("run %s task on host:%s", taskName, host.Name),
		TaskName:    taskName,
		SubTaskName: subTaskName,
		TaskParams:  taskParams,
	}
	return task
}
