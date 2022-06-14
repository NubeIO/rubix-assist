package edgeapi

import (
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/tasks"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"testing"
)

func TestManager_PipeBuilder(t *testing.T) {
	orderedTaskList := []string{"PingHost", "InstallApp"}
	host := &model.Host{
		IP:   "0.0.0.0",
		Port: 1616,
	}
	app := &App{}

	var jobs []*jobctl.JobBody
	for _, task := range orderedTaskList {
		switch task {
		case tasks.PingHost.String():
			jobs = append(jobs, buildPing(app, host))
		case tasks.InstallApp.String():
			jobs = append(jobs, buildInstall(app, host))
		case tasks.InstallPlugin.String():

		case tasks.RemoveApp.String():

		case tasks.StopApp.String():

		case tasks.StartApp.String():

		}

	}

	pprint.PrintJOSN(jobs)

}
