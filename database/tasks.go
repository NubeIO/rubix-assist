package base

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ttime"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/service/tasks"
)

const taskName = "task"

func (inst *DB) GetTask(uuid string) (*model.Task, error) {
	m := new(model.Task)
	if err := inst.DB.Where("uuid = ? ", uuid).Preload("Transactions").First(&m).Error; err != nil {
		logger.Errorf("GetTask error: %v", err)
		return nil, err
	}
	return m, nil
}

func (inst *DB) GetTasks() ([]*model.Task, error) {
	var m []*model.Task
	if err := inst.DB.Preload("Transactions").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

// GetTaskByField returns the object for the given field ie name or nil.
func (inst *DB) GetTaskByField(field string, value string) (*model.Task, error) {
	var m *model.Task
	f := fmt.Sprintf("%s = ? ", field)
	query := inst.DB.Where(f, value).First(&m)
	if query.Error != nil {
		return nil, handelNotFound(taskName)
	}
	return m, nil
}

// GetTaskByType get a Task by type and its uuid
func (inst *DB) GetTaskByType(uuid string, TaskType string) (*model.Task, error) {
	var m *model.Task
	f := "host_uuid = ? AND task_type = ?"
	query := inst.DB.Where(f, uuid, TaskType).First(&m)
	if query.Error != nil {
		return nil, handelNotFound(taskName)
	}
	return m, nil
}

func (inst *DB) CreateTask(task *model.Task) (*model.Task, error) {
	host, err := inst.GetHost(task.HostUUID)
	if err != nil {
		return nil, errors.New("no valid host found")
	}
	err = tasks.CheckTask(task.Type)
	if err != nil {
		return nil, err
	}

	if task.IsPipeline {
		task.IsPipeline = true
	}
	if task.IsJob {
		task.IsJob = true
	}

	task.UUID = uuid.ShortUUID("tas")
	task.HostUUID = host.UUID
	if task.HostUUID == "" {
		return nil, errors.New("host uuid can not be empty")
	}
	task.HostName = host.Name
	task.CreatedAt = ttime.Now()
	if err := inst.DB.Create(&task).Error; err != nil {
		return nil, err
	} else {
		return task, nil
	}
}

func (inst *DB) GetTaskByPipelineUUID(uuid string) (*model.Task, error) {
	m := new(model.Task)
	if err := inst.DB.Where("pipeline_uuid = ? ", uuid).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) TaskEntry(task *model.Task) (*model.Task, error) {
	if task.IsPipeline {
		existing, _ := inst.GetTaskByPipelineUUID(task.PipelineUUID)
		if existing != nil { // add entry
			transaction := &model.Transaction{}
			transaction, err := inst.CreateTransaction(transaction)
			if err != nil {
				return nil, err
			}
			return existing, nil
		} else {
			pipeline, err := inst.CreateTaskPipeline(task)
			if err != nil {
				return nil, err
			}
			return pipeline, err
		}
	} else {
		createTask, err := inst.CreateTask(task)
		if err != nil {
			return nil, err
		}
		return createTask, err
	}
}

func (inst *DB) CreateTaskPipeline(task *model.Task) (*model.Task, error) {
	if task.PipelineUUID == "" {
		return nil, errors.New("pipeline uuid cant not be empty")
	}
	task.FromAutomater = true
	task.IsPipeline = true
	task, err := inst.CreateTask(task)
	return task, err
}

func (inst *DB) UpdateTask(uuid string, Task *model.Task) (*model.Task, error) {
	m := new(model.Task)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(Task)
	if query.Error != nil {
		return nil, handelNotFound(taskName)
	} else {
		return Task, query.Error
	}
}

func (inst *DB) DeleteTask(uuid string) (*DeleteMessage, error) {
	m := new(model.Task)
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropTasks() (*DeleteMessage, error) {
	var m *model.Task
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
