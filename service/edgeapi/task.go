package edgeapi

import (
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/tasks"
	automodel "github.com/NubeIO/rubix-automater/automater/model"
)

/*
- post a new pipeline
- create
*/

func (inst *Manager) taskEntry(host *model.Host, data *automodel.Pipeline) error {
	task := &model.Task{
		Type:         tasks.InstallApp.String(),
		HostUUID:     host.UUID,
		HostName:     host.Name,
		IsPipeline:   true,
		PipelineUUID: data.UUID,
		Status:       data.Status.String(),
	}
	_, err := inst.DB.TaskEntry(task)
	if err != nil {
		return err
	}
	return nil

}

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
