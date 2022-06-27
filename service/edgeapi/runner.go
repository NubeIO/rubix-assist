package edgeapi

import (
	"errors"
	"github.com/NubeIO/rubix-assist/service/tasks"
)

type TaskResponse struct {
	Message      interface{} `json:"message"`
	ErrorMessage string      `json:"errorMessage"`
	Error        interface{} `json:"error"`
}

func (inst *Manager) TaskRunner(appTask *AppTask) (*TaskResponse, error) {

	resp := &TaskResponse{}
	host, err, _ := inst.getHost(appTask)
	if err != nil {
		resp.Message = "failed to find host"
		resp.Error = err
		return resp, err
	}
	task := appTask.SubTask
	switch task {
	case tasks.PingHost.String():
		_, err := runPingHost(host.IP, host.Port, 2)
		if err != nil {
			resp.Message = "host is offline"
			resp.ErrorMessage = err.Error()
			resp.Error = err
			return resp, err
		} else {
			resp.Message = "host found"
			return resp, nil
		}
	case tasks.InstallApp.String():
		data, r := inst.RunInstall(appTask)
		if r.StatusCode > 299 {
			resp.Message = r.Message
			resp.Error = r.Message
			resp.Error = errors.New("r.Message")
			return resp, err
		} else {
			resp.Message = data.Message
			return resp, nil
		}

	case tasks.InstallPlugin.String():

	case tasks.RemoveApp.String():

	case tasks.StopApp.String():

	case tasks.StartApp.String():

	}
	return resp, errors.New("no valid sub task found")

}
