package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
)

const locationName = "location"

func (d *DB) GetLocations() ([]*model.Location, error) {
	var m []*model.Location
	if err := d.DB.Preload("Networks.Hosts").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateLocationWizard(body *model.Location) (*model.Location, error) {
	location, err := d.CreateLocation(body)
	if err != nil {
		return nil, err
	}

	network, err := d.CreateHostNetwork(&model.Network{LocationUUID: location.UUID})
	if err != nil {
		return nil, err
	}

	_, err = d.CreateHost(&model.Host{NetworkUUID: network.UUID})
	if err != nil {
		return nil, err
	}

	return location, nil

}

func (d *DB) GetLocationsByName(name string, isUUID bool) (*model.Location, error) {
	m := new(model.Location)
	switch isUUID {
	case true:
		if err := d.DB.Where("uuid = ? ", name).Preload("Networks.Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(locationName)
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", name).Preload("Networks.Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(locationName)
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}

func (d *DB) CreateLocation(body *model.Location) (*model.Location, error) {
	if body.Name == "" {
		body.Name = "cloud"
	}
	existingHost, _ := d.GetLocationsByName(body.Name, false)
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
		return nil, handelNotFound(locationName)
	} else {
		return m, nil
	}
}

func (d *DB) UpdateLocation(uuid string, host *model.Location) (*model.Location, error) {
	m := new(model.Location)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(locationName)
	} else {
		return m, nil
	}
}

func (d *DB) DeleteLocation(uuid string) (*DeleteMessage, error) {
	var m *model.Location
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (d *DB) DropLocations() (*DeleteMessage, error) {
	var m *model.Location
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
