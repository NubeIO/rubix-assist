package dbase

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist-model/model"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (d *DB) GetLocations() ([]*model.Location, error) {
	var m []*model.Location
	if err := d.DB.Preload("Networks.Hosts").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) GetLocationsByName(name string, isUUID bool) (*model.Location, error) {
	m := new(model.Location)
	switch isUUID {
	case true:
		if err := d.DB.Where("uuid = ? ", name).Preload("Networks.Hosts").First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", name).Preload("Networks").First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}

func (d *DB) CreateLocation(body *model.Location) (*model.Location, error) {
	if body.Name == "" {
		body.Name = uuid.ShortUUID("location")
	}
	existingHost, _ := d.GetLocationsByName(body.Name, false)
	fmt.Println(1111, existingHost)
	if existingHost != nil {
		return nil, errors.New("a location with this name exists")
	}
	body.UUID = uuid.ShortUUID("loc")
	if err := d.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (d *DB) UpdateLocationsByName(name string, host *model.Location) (*model.Location, error) {
	m := new(model.Location)
	query := d.DB.Where("name = ?", name).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return m, query.Error
	}
}

func (d *DB) UpdateLocation(uuid string, host *model.Location) (*model.Location, error) {
	m := new(model.Location)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return host, query.Error
	}
}

func (d *DB) DeleteLocation(uuid string) (*DeleteMessage, error) {
	var m *model.Location
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (d *DB) DropLocations() (bool, error) {
	var m *model.Location
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
