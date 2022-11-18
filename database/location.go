package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	model2 "github.com/NubeIO/rubix-assist/model"
)

const locationName = "location"

func (inst *DB) GetLocations() ([]*model2.Location, error) {
	var m []*model2.Location
	if err := inst.DB.Preload("Networks.Hosts").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) CreateLocationWizard(body *model2.Location) (*model2.Location, error) {
	location, err := inst.CreateLocation(body)
	if err != nil {
		return nil, err
	}

	network, err := inst.CreateHostNetwork(&model2.Network{LocationUUID: location.UUID})
	if err != nil {
		return nil, err
	}

	_, err = inst.CreateHost(&model2.Host{NetworkUUID: network.UUID})
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (inst *DB) GetLocationsByName(name string, isUUID bool) (*model2.Location, error) {
	m := new(model2.Location)
	switch isUUID {
	case true:
		if err := inst.DB.Where("uuid = ? ", name).Preload("Networks.Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(locationName)
		}
		return m, nil
	case false:
		if err := inst.DB.Where("name = ? ", name).Preload("Networks.Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(locationName)
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}

func (inst *DB) CreateLocation(body *model2.Location) (*model2.Location, error) {
	if body.Name == "" {
		body.Name = "cloud"
	}
	existingHost, _ := inst.GetLocationsByName(body.Name, false)
	if existingHost != nil {
		return nil, errors.New("a location with this name exists")
	}
	body.UUID = uuid.ShortUUID("loc")
	if err := inst.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (inst *DB) UpdateLocationsByName(name string, host *model2.Location) (*model2.Location, error) {
	m := new(model2.Location)
	query := inst.DB.Where("name = ?", name).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(locationName)
	} else {
		return m, nil
	}
}

func (inst *DB) UpdateLocation(uuid string, host *model2.Location) (*model2.Location, error) {
	m := new(model2.Location)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(locationName)
	} else {
		return m, nil
	}
}

func (inst *DB) DeleteLocation(uuid string) (*DeleteMessage, error) {
	var m *model2.Location
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropLocations() (*DeleteMessage, error) {
	var m *model2.Location
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
