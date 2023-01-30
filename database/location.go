package base

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
)

const locationName = "location"

func (inst *DB) GetLocations() ([]*amodel.Location, error) {
	var m []*amodel.Location
	if err := inst.DB.Preload("Networks.Hosts").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) GetLocation(uuid string) (*amodel.Location, error) {
	var m *amodel.Location
	if err := inst.DB.Where("uuid = ? ", uuid).Preload("Networks.Hosts").First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) CreateLocation(body *amodel.Location) (*amodel.Location, error) {
	body.UUID = uuid.ShortUUID("loc")
	if err := inst.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (inst *DB) UpdateLocation(uuid string, host *amodel.Location) (*amodel.Location, error) {
	var m *amodel.Location
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(locationName)
	} else {
		return m, nil
	}
}

func (inst *DB) DeleteLocation(uuid string) (*DeleteMessage, error) {
	var m *amodel.Location
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropLocations() (*DeleteMessage, error) {
	var m *amodel.Location
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
