package automate

import (
	"github.com/NubeIO/rubix-automater/automater/model"
	"github.com/NubeIO/rubix-automater/controller/jobctl"
	"github.com/NubeIO/rubix-automater/controller/pipectl"
	"github.com/NubeIO/rubix-automater/service/client"
	"github.com/NubeIO/rubix-automater/service/tasks/install"
)

type App struct {
	HostUUID string `json:"host_uuid"`
	HostName string `json:"host_name"`
	AppName  string `json:"app_name"`
	Version  string `json:"version"`
}

func AddPipeline(app *App) (data *model.Pipeline, response *client.Response) {

	cli := client.New("0.0.0.0", 1663)

	taskBody := install.AppParams{
		HostName: app.HostName,
		AppName:  app.AppName,
		Version:  app.Version,
	}

	jobOne := &jobctl.JobBody{
		Name:       "ping 1",
		TaskName:   "pingHost",
		Options:    &model.JobOptions{},
		TaskParams: map[string]interface{}{"url": "0.0.0.0", "port": 1662},
	}

	jobTwo := &jobctl.JobBody{
		Name:       "install app",
		TaskName:   "installApp",
		Options:    &model.JobOptions{},
		TaskParams: map[string]interface{}{"hostName": taskBody.HostName, "appName": taskBody.AppName, "version": taskBody.Version},
	}

	var jobs []*jobctl.JobBody
	jobs = append(jobs, jobOne)
	jobs = append(jobs, jobTwo)

	body := &pipectl.PipelineBody{
		Name:       "install app pipeline",
		Jobs:       jobs,
		ScheduleAt: "10 sec",
		PipelineOptions: &model.PipelineOptions{
			EnableInterval:   false,
			RunOnInterval:    "10 sec",
			DelayBetweenTask: 1,
			CancelOnFailure:  false,
		},
	}

	return cli.AddPipeline(body)

}
