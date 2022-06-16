package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ttime"
	"github.com/NubeIO/rubix-assist/service/tasks"

	"github.com/NubeIO/rubix-assist/pkg/model"
)

func (d *DB) GetTransaction(uuid string) (*model.Transaction, error) {
	m := new(model.Transaction)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (d *DB) GetTransactions() ([]*model.Transaction, error) {
	var m []*model.Transaction
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	Task, err := d.GetTask(transaction.TaskUUID)
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
	if err := d.DB.Create(&transaction).Error; err != nil {
		return nil, err
	} else {
		return transaction, nil
	}
}

func (d *DB) UpdateTransaction(uuid string, message *model.Transaction) (*model.Transaction, error) {
	m := new(model.Transaction)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(message)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return message, query.Error
	}
}

func (d *DB) DeleteTransaction(uuid string) (*DeleteMessage, error) {
	m := new(model.Transaction)
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

// DropTransactions delete all.
func (d *DB) DropTransactions() (*DeleteMessage, error) {
	var m *model.Transaction
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
