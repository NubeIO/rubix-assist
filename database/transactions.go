package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ttime"
	"github.com/NubeIO/rubix-assist/service/tasks"
)

func (inst *DB) GetTransaction(uuid string) (*amodel.Transaction, error) {
	m := new(amodel.Transaction)
	if err := inst.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) GetTransactions() ([]*amodel.Transaction, error) {
	var m []*amodel.Transaction
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) CreateTransaction(transaction *amodel.Transaction) (*amodel.Transaction, error) {
	Task, err := inst.GetTask(transaction.TaskUUID)
	if err != nil {
		return nil, errors.New("no task uuid was found, a translation needs to be added to an existing task")
	}
	err = tasks.CheckTask(transaction.TaskType)
	if err != nil {
		return nil, err
	}
	transaction.UUID = uuid.ShortUUID("trn")
	transaction.TaskUUID = Task.UUID
	t := ttime.Now()
	transaction.CreatedAt = &t
	if err := inst.DB.Create(&transaction).Error; err != nil {
		return nil, err
	} else {
		return transaction, nil
	}
}

func (inst *DB) UpdateTransaction(uuid string, message *amodel.Transaction) (*amodel.Transaction, error) {
	m := new(amodel.Transaction)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(message)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return message, query.Error
	}
}

func (inst *DB) DeleteTransaction(uuid string) (*DeleteMessage, error) {
	m := new(amodel.Transaction)
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropTransactions() (*DeleteMessage, error) {
	var m *amodel.Transaction
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
