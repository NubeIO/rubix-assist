package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"

	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/model"
)

func (d *DB) GetTransaction(uuid string) (*model.Transaction, error) {
	m := new(model.Transaction)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetTransaction error: %v", err)
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

func (d *DB) CreateTransaction(message *model.Transaction) (*model.Transaction, error) {
	Task, err := d.GetTask(message.TaskUUID)
	if err != nil {
		return nil, errors.New("no Task found")
	}
	message.UUID = uuid.ShortUUID("msg")
	message.TaskUUID = Task.UUID
	if err := d.DB.Create(&message).Error; err != nil {
		return nil, err
	} else {
		return message, nil
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

func (d *DB) DeleteTransaction(uuid string) (ok bool, err error) {
	m := new(model.Transaction)
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

// DropTransactions delete all.
func (d *DB) DropTransactions() (bool, error) {
	var m *model.Transaction
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}
