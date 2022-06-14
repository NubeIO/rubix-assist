package edgeapi

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/service/tasks"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

type TaskResponse struct {
	Message      interface{} `json:"message"`
	ErrorMessage string      `json:"errorMessage"`
	Error        interface{} `json:"error"`
}

func (inst *Manager) TaskRunner(appTask *AppTask) (*TaskResponse, error) {

	resp := &TaskResponse{}
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
		resp.Message = "failed to find host"
		resp.Error = err
		return resp, err
	}
	task := appTask.SubTask

	switch task {
	case tasks.PingHost.String():
		_, err := runPingHost(host.IP, host.RubixPort, 2)
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

	case tasks.InstallPlugin.String():

	case tasks.RemoveApp.String():

	case tasks.StopApp.String():

	case tasks.StartApp.String():

	}
	return resp, errors.New("no valid sub task found")

}

func runPingHost(url string, port int, countSetting int) (bool, error) {
	failCount := 0
	for i := 1; i <= 1; i++ { //ping 3 times
		if ping(url, port) {
			logrus.Infoln("run task ping host ok:", fmt.Sprintf("%s:%d", url, port))
		} else {
			failCount++
			logrus.Infoln("run task ping host:", fmt.Sprintf("%s:%d", url, port), " fail count:", failCount)
		}
	}
	if failCount >= 1 {
		return false, errors.New(fmt.Sprintf("ping fail count:%d was grater then the allowable ping fail count %d", failCount, countSetting))
	}
	return true, nil
}

func ping(url string, port int) (found bool) {
	conn, err := net.DialTimeout("tcp",
		fmt.Sprintf("%s:%d", url, port),
		300*time.Millisecond)
	if err == nil {
		conn.Close()
		return true
	}
	logrus.Errorln("run task ping error:", err)
	return false

}
