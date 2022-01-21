package dbase

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/config"
	"github.com/NubeIO/rubix-updater/pkg/logger"
)

func (d *DB) GetMessage(id string) (*model.Message, error) {
	m := new(model.Message)
	if err := d.DB.Where("id = ? ", id).First(&m).Error; err != nil {
		logger.Errorf("GetMessage error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetMessages() ([]model.Message, error) {
	var m []model.Message
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateMessage(Message *model.Message) (*model.Message, error) {
	Message.ID = config.MakeTopicUUID(model.CommonNaming.Message)
	if err := d.DB.Create(&Message).Error; err != nil {
		return nil, err
	} else {
		return Message, nil
	}
}

func (d *DB) UpdateMessage(id string, Message *model.Message) (*model.Message, error) {
	m := new(model.Message)
	query := d.DB.Where("id = ?", id).Find(&m).Updates(Message)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Message, query.Error
	}
}

func (d *DB) DeleteMessage(id string) (ok bool, err error) {
	m := new(model.Message)
	query := d.DB.Where("id = ? ", id).Delete(&m)
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
