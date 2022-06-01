package dbase

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"

	"github.com/NubeIO/rubix-assist-model/model"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/oleiade/reflections"
)

func (d *DB) GetAlert(uuid string) (*model.Alert, error) {
	m := new(model.Alert)
	if err := d.DB.Where("uuid = ? ", uuid).Preload("Messages").First(&m).Error; err != nil {
		logger.Errorf("GetAlert error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetAlerts() ([]model.Alert, error) {
	var m []model.Alert
	if err := d.DB.Preload("Messages").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

// GetAlertByField returns the object for the given field ie name or nil.
func (d *DB) GetAlertByField(field string, value string) (*model.Alert, error) {
	var m *model.Alert
	f := fmt.Sprintf("%s = ? ", field)
	query := d.DB.Where(f, value).First(&m)
	if query.Error != nil {
		return nil, query.Error
	}
	return m, nil
}

// GetAlertByType get an alert by type and its uuid
func (d *DB) GetAlertByType(uuid string, alertType string) (*model.Alert, error) {
	var m *model.Alert
	f := "host_uuid = ? AND alert_type = ?"
	query := d.DB.Where(f, uuid, alertType).First(&m)
	if query.Error != nil {
		return nil, query.Error
	}
	return m, nil
}

func (d *DB) CreateAlert(alert *model.Alert) (*model.Alert, error) {
	host, err := d.GetHostByName(alert.HostUUID, true)
	if err != nil {
		return nil, errors.New("no valid host found")
	}
	items, err := reflections.Items(model.CommonAlertTypes)
	typeExist := false
	for _, a := range items {
		if alert.AlertType == a {
			typeExist = true
		}
	}
	if !typeExist {
		return nil, errors.New("incorrect AlertType provided")
	}
	alert.UUID = uuid.ShortUUID("alt")
	alert.HostUUID = host.UUID
	if err := d.DB.Create(&alert).Error; err != nil {
		return nil, err
	} else {
		return alert, nil
	}
}

func (d *DB) UpdateAlert(uuid string, Alert *model.Alert) (*model.Alert, error) {
	m := new(model.Alert)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(Alert)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Alert, query.Error
	}
}

func (d *DB) DeleteAlert(uuid string) (ok bool, err error) {
	m := new(model.Alert)
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

// DropAlerts delete all.
func (d *DB) DropAlerts() (bool, error) {
	var m *model.Alert
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	var msg *model.Message
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
