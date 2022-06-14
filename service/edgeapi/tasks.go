package edgeapi

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/tasks"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
)

type AppTask struct {
	LocationName string `json:"locationName"`
	NetworkName  string `json:"networkName"`
	HostName     string `json:"hostName"`
	HostUUID     string `json:"hostUUID"`
	AppName      string `json:"appName"`
	SubTask      string `json:"subTask"`
	Version      string `json:"version"`
}

func buildPing(app *App, host *model.Host) *jobctl.JobBody {
	taskName := tasks.SubTask.String()
	subTaskName := tasks.PingHost.String()
	//params :=[]AppTask{}
	taskParams := map[string]interface{}{
		"locationName": host.IP,
		"networkName":  host.RubixPort,
		"hostName":     host.RubixPort,
		"hostUUID":     host.RubixPort,
		"appName":      host.RubixPort,
		"subTask":      host.RubixPort,
		"version":      host.RubixPort,
	}
	task := &jobctl.JobBody{
		Name:       fmt.Sprintf("run %s task on host:%s", taskName, host.Name),
		TaskName:   subTaskName,
		TaskParams: taskParams,
	}
	return task
}

func buildInstall(app *App, host *model.Host) *jobctl.JobBody {
	taskName := tasks.PingHost.String()
	subTaskName := tasks.InstallApp.String()
	taskParams := map[string]interface{}{
		"url":  host.IP,
		"port": host.RubixPort,
	}
	task := &jobctl.JobBody{
		Name:       fmt.Sprintf("run %s task on host:%s", taskName, host.Name),
		TaskName:   subTaskName,
		TaskParams: taskParams,
	}
	return task
}
