package edgeapi

import (
	"fmt"
	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/NubeIO/edge/service/client"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/tasks"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
)

func buildInstall(appTask *AppTask, host *model.Host) *jobctl.JobBody {
	taskName := tasks.SubTask.String()
	subTaskName := tasks.InstallApp.String()
	taskParams := map[string]interface{}{
		"locationName":       appTask.LocationName,
		"networkName":        appTask.NetworkName,
		"hostName":           host.Name,
		"hostUUID":           host.UUID,
		"appName":            appTask.AppName,
		"version":            appTask.Version,
		"manualInstall":      appTask.ManualInstall,
		"manualAssetZipName": appTask.ManualAssetZipName,
		"manualAssetTag":     appTask.ManualAssetTag,
		"cleanup":            appTask.Cleanup,
	}
	task := &jobctl.JobBody{
		Name:        fmt.Sprintf("run %s task on host:%s", subTaskName, host.Name),
		TaskName:    taskName,
		SubTaskName: subTaskName,
		TaskParams:  taskParams,
	}
	return task
}

func (inst *Manager) RunInstall(body *AppTask) (*installer.InstallResponse, *client.Response) {
	response := &client.Response{}
	host, err, token := inst.getHost(body)
	if err != nil {
		response.Message = err.Error()
		return nil, response
	}
	app := &installer.App{
		AppName: body.AppName,
		Token:   token,
		Version: body.Version,
	}
	data, resp := inst.reset(host.IP, host.Port).InstallApp(app)
	if resp.StatusCode > 299 {
		return data, resp
	}
	return data, resp
}
