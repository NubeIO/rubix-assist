package edgeapi

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/tasks"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func buildPing(appTask *AppTask, host *model.Host) *jobctl.JobBody {
	taskName := tasks.SubTask.String()
	subTaskName := tasks.PingHost.String()
	taskParams := map[string]interface{}{
		"locationName": appTask.LocationName,
		"networkName":  appTask.NetworkName,
		"hostName":     host.Name,
		"hostUUID":     host.UUID,
		"appName":      appTask.AppName,
		"subTask":      subTaskName,
	}
	task := &jobctl.JobBody{
		Name:        fmt.Sprintf("run %s task on host:%s", subTaskName, host.Name),
		TaskName:    taskName,
		SubTaskName: subTaskName,
		TaskParams:  taskParams,
	}
	return task
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
