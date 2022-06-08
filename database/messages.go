package dbase

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"

	"github.com/NubeIO/rubix-assist-model/model"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (d *DB) GetMessage(uuid string) (*model.Message, error) {
	m := new(model.Message)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetMessage error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetMessages() ([]*model.Message, error) {
	var m []*model.Message
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateMessage(message *model.Message) (*model.Message, error) {
	alert, err := d.GetAlert(message.AlertUUID)
	if err != nil {
		return nil, errors.New("no alert found")
	}
	message.UUID = uuid.ShortUUID("msg")
	message.AlertUUID = alert.UUID
	if err := d.DB.Create(&message).Error; err != nil {
		return nil, err
	} else {
		return message, nil
	}
}

func (d *DB) UpdateMessage(uuid string, message *model.Message) (*model.Message, error) {
	m := new(model.Message)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(message)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return message, query.Error
	}
}

func (d *DB) DeleteMessage(uuid string) (ok bool, err error) {
	m := new(model.Message)
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

// DropMessages delete all.
func (d *DB) DropMessages() (bool, error) {
	var m *model.Message
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
