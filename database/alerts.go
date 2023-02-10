package base

import (
	"errors"
	"fmt"

	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ttime"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/service/alerts"
	"gorm.io/gorm"
)

const alertName = "alert"

func (inst *DB) GetAlert(uuid string) (*amodel.Alert, error) {
	m := new(amodel.Alert)
	if err := inst.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetAlert error: %v", err)
		return nil, err
	}
	return m, nil
}

func (inst *DB) GetAlerts() ([]*amodel.Alert, error) {
	var m []*amodel.Alert
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

// GetAlertByField returns the object for the given field ie name or nil.
func (inst *DB) GetAlertByField(field string, value string) (*amodel.Alert, error) {
	var m *amodel.Alert
	f := fmt.Sprintf("%s = ? ", field)
	query := inst.DB.Where(f, value).First(&m)
	if query.Error != nil {
		return nil, handelNotFound(alertName)
	}
	return m, nil
}

func (inst *DB) CreateAlert(alert *amodel.Alert) (*amodel.Alert, error) {
	var err error
	if alert.Status == "" {
		alert.Status = string(alerts.Active)
	} else {
		if err = alerts.CheckStatus(alert.Status); err != nil {
			return nil, err
		}
	}
	if err = alerts.CheckAlertType(alert.Type); err != nil {
		return nil, err
	}
	if err = alerts.CheckEntityType(alert.EntityType); err != nil {
		return nil, err
	}
	alert.UUID = uuid.ShortUUID("alr")
	t := ttime.Now()
	alert.CreatedAt = &t
	if err := inst.DB.Create(&alert).Error; err != nil {
		return nil, err
	}
	return alert, nil
}

func (inst *DB) UpdateAlertStatus(uuid string, status string) (alert *amodel.Alert, err error) {
	var query *gorm.DB
	if err = alerts.CheckStatus(status); err != nil {
		return nil, err
	}
	if alerts.CheckStatusClosed(status) { // Move alert to alertClosed table
		a := amodel.Alert{}
		query = inst.DB.Where("uuid = ?", uuid).First(&a)
		if query.Error != nil {
			return nil, query.Error
		}
		ac := amodel.AlertClosed{
			Alert: a,
		}
		ac.Status = status
		query = inst.DB.Create(&ac)
		if query.Error != nil {
			return nil, query.Error
		}
		query = inst.DB.Delete(&a)
		if query.Error != nil {
			return nil, query.Error
		}
	} else { // else update alert status
		alert = &amodel.Alert{}
		query = inst.DB.Model(&alert).Where("uuid = ?", uuid).Update("status", status)
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			// look for alert in closed table and re-open
			ac := amodel.AlertClosed{}
			query = inst.DB.Where("uuid = ?", uuid).Find(&ac)
			if query.Error != nil {
				return nil, query.Error
			}
			alert = &ac.Alert
			alert.Status = status
			query = inst.DB.Create(&alert)
			if query.Error != nil {
				return nil, query.Error
			}
			query = inst.DB.Delete(&ac)
			if query.Error != nil {
				return nil, query.Error
			}
		}
	}
	return alert, query.Error
}

func (inst *DB) DeleteAlert(uuid string) (*DeleteMessage, error) {
	m := new(amodel.Alert)
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropAlerts() (*DeleteMessage, error) {
	var m *amodel.Alert
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
