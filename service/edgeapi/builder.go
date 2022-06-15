package edgeapi

import (
	"github.com/NubeIO/rubix-assist/service/autocli"
	"github.com/NubeIO/rubix-assist/service/tasks"
	automodel "github.com/NubeIO/rubix-automater/automater/model"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"github.com/NubeIO/rubix-automater/controller/pipectl"
)

type AppTask struct {
	Description        string   `json:"description"`
	LocationName       string   `json:"locationName"`
	NetworkName        string   `json:"networkName"`
	HostName           string   `json:"hostName"`
	HostUUID           string   `json:"hostUUID"`
	AppName            string   `json:"appName"`
	SubTask            string   `json:"subTask"`
	Version            string   `json:"version"`
	ManualInstall      bool     `json:"manualInstall"`      // will not download from GitHub, and will use the app-store download path
	ManualAssetZipName string   `json:"manualAssetZipName"` // flow-framework-0.5.5-1575cf89.amd64.zip
	ManualAssetTag     string   `json:"manualAssetTag"`     // this is the release tag as in v0.0.1
	Cleanup            bool     `json:"cleanup"`
	OrderedTaskList    []string `json:"orderedTaskList"`
}

func (inst *Manager) TaskBuilder(appTask *AppTask) (*automodel.Pipeline, *autocli.Response) {
	resp := &autocli.Response{}
	host, err, _ := inst.getHost(appTask)
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
	description := "run assist sub-task"
	if appTask.Description != "" {
		description = appTask.Description
	}
	pipeBuilder := &pipectl.PipelineBody{
		Name:       description,
		Jobs:       jobs,
		ScheduleAt: "1 sec",
		PipelineOptions: &automodel.PipelineOptions{
			EnableInterval: false,
			RunOnInterval:  "1 sec",
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
