package edgeapi

import (
	"context"
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/model"
	automodel "github.com/NubeIO/rubix-automater/automater/model"
	"github.com/mustafaturan/bus/v3"
	log "github.com/sirupsen/logrus"
)

const allTopic = "*"
const automater = "automater-transaction"

func (inst *Manager) transactionEntry(data *automodel.PublishTransaction) error {
	task, err := inst.DB.GetTaskByPipelineUUID(data.PipelineID)
	if err != nil {
		return err
	}
	transaction := &model.Transaction{
		Status:        data.Status,
		TaskType:      data.TaskType,
		SubTaskType:   data.SubTaskType,
		FailureReason: data.FailureReason,
		Data:          nil,
		TaskUUID:      task.UUID,
		JobID:         data.JobID,
		IsPipeLine:    data.IsPipeLine,
		PipelineID:    data.PipelineID,
		CreatedAt:     data.CreatedAt,
		StartedAt:     data.StartedAt,
		CompletedAt:   data.CompletedAt,
		Duration:      0,
	}

	_, err = inst.DB.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil

}

func (inst *Manager) registerAutomater() {
	handler := bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			go func() {
				switch e.Topic {
				case automater:
					payload, ok := e.Data.(*automodel.PublishTransaction)
					if ok {
						msg := fmt.Sprintf("automater redis message topic: %s msg: %s %s", e.Topic, payload.Status, payload.Status)
						log.Info(msg)
						err := inst.transactionEntry(payload)
						if err != nil {
							return
						}
					}
				}
			}()
		},
		Matcher: automater,
	}
	inst.Events.EventBus.Bus.RegisterHandler(automater, handler)
}
