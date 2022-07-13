package edgeapi

import (
	"errors"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/tasks"
	log "github.com/sirupsen/logrus"
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
		resp.Message = err.Error()
		resp.Error = err
		return resp, err
	}
	log.Infof("task runner SUB-TASK resp from automater:%s", appTask.SubTask)
	task := appTask.SubTask
	pprint.PrintJOSN(appTask)
	switch task {
	case tasks.PingHost.String():
		_, err := runPingHost(host.IP, host.Port, 2)
		if err != nil {
			resp.Message = "host is offline"
			resp.ErrorMessage = err.Error()
			resp.Error = err
			log.Errorf("task runner PingHost resp from automater:%s", resp.Error)
			return resp, err
		} else {
			resp.Message = "host found ok!"
			log.Infof("task runner PingHost resp from automater:%s", resp.Message)
			return resp, nil
		}
	case tasks.InstallApp.String():
		data, r := inst.RunInstall(appTask)
		if r.StatusCode > 299 {
			resp.Message = r.Message
			resp.Error = r.Message
			resp.Error = errors.New("r.Message")
			log.Errorf("task runner RunInstall resp from automater:%s", resp.Error)
			pprint.PrintJOSN(data)
			return resp, err
		} else {
			resp.Message = data.Message
			log.Infof("task runner InstallApp resp from automater:%s", resp.Message)
			pprint.PrintJOSN(data)
			return resp, nil
		}

	case tasks.InstallPlugin.String():

	case tasks.RemoveApp.String():

	case tasks.StopApp.String():

	case tasks.StartApp.String():

	}
	return resp, errors.New("no valid sub task found")

}
