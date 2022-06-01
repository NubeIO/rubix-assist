package dbase

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/rubix-assist-model/model"
	"gorm.io/gorm"

	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (d *DB) GetHosts() ([]*model.Host, error) {
	var m []*model.Host
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) GetHostByName(name string, isUUID bool) (*model.Host, error) {
	m := new(model.Host)
	switch isUUID {
	case true:
		if err := d.DB.Where("uuid = ? ", name).First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", name).First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}

func (d *DB) CreateHost(host *model.Host) (*model.Host, error) {
	if host.Name == "" {
		host.Name = uuid.ShortUUID("host")
	}
	existingHost, _ := d.GetHostByName(host.Name, false)
	if existingHost != nil {
		return nil, errors.New("a host with this name exists")
	}
	host.UUID = uuid.ShortUUID("hos")
	if host.PingEnable == nil {
		host.PingEnable = nils.NewFalse()
	}
	if host.IsLocalhost == nil {
		host.IsLocalhost = nils.NewFalse()
	}
	if host.HTTPS == nil {
		host.HTTPS = nils.NewFalse()
	}
	if host.Port == 0 {
		host.Port = 22
	}
	if host.RubixPort == 0 {
		host.RubixPort = 1660
	}
	if err := d.DB.Create(&host).Error; err != nil {
		return nil, err
	} else {
		return host, nil
	}
}

func (d *DB) UpdateHostByName(name string, host *model.Host) (*model.Host, error) {
	m := new(model.Host)
	query := d.DB.Where("name = ?", name).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return m, query.Error
	}
}

func (d *DB) UpdateHost(uuid string, host *model.Host) (*model.Host, error) {
	m := new(model.Host)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return host, query.Error
	}
}

func (d *DB) DeleteHost(uuid string) (*DeleteMessage, error) {
	var m *model.Host
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

type DeleteMessage struct {
	Message string `json:"message"`
}

func deleteResponse(query *gorm.DB) (*DeleteMessage, error) {
	msg := &DeleteMessage{
		Message: "delete failed",
	}
	if query.Error != nil {
		return msg, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return msg, errors.New("not found")
	}
	msg.Message = "deleted ok"
	return msg, nil
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
