package dbase

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/config"
	"github.com/NubeIO/rubix-updater/pkg/logger"
)

func (d *DB) GetAlert(id string) (*model.Alert, error) {
	m := new(model.Alert)
	if err := d.DB.Where("id = ? ", id).First(&m).Error; err != nil {
		logger.Errorf("GetAlert error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetAlerts() ([]model.Alert, error) {
	var m []model.Alert
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateAlert(Alert *model.Alert) (*model.Alert, error) {
	Alert.ID = config.MakeTopicUUID(model.CommonNaming.Alert)
	if err := d.DB.Create(&Alert).Error; err != nil {
		return nil, err
	} else {
		return Alert, nil
	}
}

func (d *DB) UpdateAlert(id string, Alert *model.Alert) (*model.Alert, error) {
	m := new(model.Alert)
	query := d.DB.Where("id = ?", id).Find(&m).Updates(Alert)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Alert, query.Error
	}
}

func (d *DB) DeleteAlert(id string) (ok bool, err error) {
	m := new(model.Alert)
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

// DropAlerts delete all.
func (d *DB) DropAlerts() (bool, error) {
	var m *model.Alert
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
