package base

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/model"
)

func (d *DB) GetTask(uuid string) (*model.Task, error) {
	m := new(model.Task)
	if err := d.DB.Where("uuid = ? ", uuid).Preload("Messages").First(&m).Error; err != nil {
		logger.Errorf("GetTask error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetTasks() ([]*model.Task, error) {
	var m []*model.Task
	if err := d.DB.Preload("Messages").Find(&m).Error; err != nil {
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
		return nil, query.Error
	}
	return m, nil
}

// GetTaskByType get an Task by type and its uuid
func (d *DB) GetTaskByType(uuid string, TaskType string) (*model.Task, error) {
	var m *model.Task
	f := "host_uuid = ? AND Task_type = ?"
	query := d.DB.Where(f, uuid, TaskType).First(&m)
	if query.Error != nil {
		return nil, query.Error
	}
	return m, nil
}

func (d *DB) CreateTask(Task *model.Task) (*model.Task, error) {
	host, err := d.GetHostByName(Task.HostUUID, true)
	if err != nil {
		return nil, errors.New("no valid host found")
	}
	//items, err := reflections.Items("a")
	//typeExist := false
	//for _, a := range items {
	//	if Task.TaskType == a {
	//		typeExist = true
	//	}
	//}
	//if !typeExist {
	//	return nil, errors.New("incorrect TaskType provided")
	//}
	Task.UUID = uuid.ShortUUID("alt")
	Task.HostUUID = host.UUID
	if err := d.DB.Create(&Task).Error; err != nil {
		return nil, err
	} else {
		return Task, nil
	}
}

func (d *DB) UpdateTask(uuid string, Task *model.Task) (*model.Task, error) {
	m := new(model.Task)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(Task)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Task, query.Error
	}
}

func (d *DB) DeleteTask(uuid string) (ok bool, err error) {
	m := new(model.Task)
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}

// DropTasks delete all.
func (d *DB) DropTasks() (bool, error) {
	var m *model.Task
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	var msg *model.Task
	query = d.DB.Where("1 = 1")
	query.Delete(&msg)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}
