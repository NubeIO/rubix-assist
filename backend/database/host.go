package dbase

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/logger"
)

func (d *DB) GetHosts() ([]model.Host, error) {
	var m []model.Host
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) GetHostByName(id string, useID bool) (*model.Host, error) {
	m := new(model.Host)
	fmt.Println(id, useID)
	switch useID {
	case true:
		if err := d.DB.Where("id = ? ", id).First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", id).First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid id or name was provided in the request")
	}
}

func (d *DB) CreateHost(host *model.Host) (*model.Host, error) {
	host.ID, _ = uuid.MakeUUID()
	if err := d.DB.Create(&host).Error; err != nil {
		return nil, err
	} else {
		return host, nil
	}
}

func (d *DB) UpdateHost(id string, host *model.Host) (*model.Host, error) {
	m := new(model.Host)
	query := d.DB.Where("id = ?", id).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return host, query.Error
	}
}

func (d *DB) DeleteHost(id string) (ok bool, err error) {
	m := new(model.Host)
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

// DropHosts delete all.
func (d *DB) DropHosts() (bool, error) {
	var m *model.Host
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
