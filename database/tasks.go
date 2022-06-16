package base

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ttime"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/tasks"
)

const taskName = "task"

func (d *DB) GetTask(uuid string) (*model.Task, error) {
	m := new(model.Task)
	if err := d.DB.Where("uuid = ? ", uuid).Preload("Transactions").First(&m).Error; err != nil {
		logger.Errorf("GetTask error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetTasks() ([]*model.Task, error) {
	var m []*model.Task
	if err := d.DB.Preload("Transactions").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

// GetTaskByField returns the object for the given field ie name or nil.
func (d *DB) GetTaskByField(field string, value string) (*model.Task, error) {
	var m *model.Task
	f := fmt.Sprintf("%s = ? ", field)
	query := d.DB.Where(f, value).First(&m)
	if query.Error != nil {
		return nil, handelNotFound(taskName)
	}
	return m, nil
}

// GetTaskByType get a Task by type and its uuid
func (d *DB) GetTaskByType(uuid string, TaskType string) (*model.Task, error) {
	var m *model.Task
	f := "host_uuid = ? AND task_type = ?"
	query := d.DB.Where(f, uuid, TaskType).First(&m)
	if query.Error != nil {
		return nil, handelNotFound(taskName)
	}
	return m, nil
}

func (d *DB) CreateTask(task *model.Task) (*model.Task, error) {
	host, err := d.GetHostByName(task.HostUUID, true)
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
	if err := d.DB.Create(&task).Error; err != nil {
		return nil, err
	} else {
		return task, nil
	}
}

func (d *DB) GetTaskByPipelineUUID(uuid string) (*model.Task, error) {
	m := new(model.Task)
	if err := d.DB.Where("pipeline_uuid = ? ", uuid).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (d *DB) TaskEntry(task *model.Task) (*model.Task, error) {
	if task.IsPipeline {
		existing, _ := d.GetTaskByPipelineUUID(task.PipelineUUID)
		if existing != nil { // add entry
			transaction := &model.Transaction{}
			transaction, err := d.CreateTransaction(transaction)
			if err != nil {
				return nil, err
			}
			return existing, nil
		} else {
			pipeline, err := d.CreateTaskPipeline(task)
			if err != nil {
				return nil, err
			}
			return pipeline, err
		}
	} else {
		createTask, err := d.CreateTask(task)
		if err != nil {
			return nil, err
		}
		return createTask, err
	}
}

func (d *DB) CreateTaskPipeline(task *model.Task) (*model.Task, error) {
	if task.PipelineUUID == "" {
		return nil, errors.New("pipeline uuid cant not be empty")
	}
	task.FromAutomater = true
	task.IsPipeline = true
	task, err := d.CreateTask(task)
	return task, err
}

func (d *DB) UpdateTask(uuid string, Task *model.Task) (*model.Task, error) {
	m := new(model.Task)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(Task)
	if query.Error != nil {
		return nil, handelNotFound(taskName)
	} else {
		return Task, query.Error
	}
}

func (d *DB) DeleteTask(uuid string) (*DeleteMessage, error) {
	m := new(model.Task)
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

// DropTasks delete all.
func (d *DB) DropTasks() (*DeleteMessage, error) {
	var m *model.Task
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
