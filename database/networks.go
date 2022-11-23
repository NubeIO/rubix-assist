package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
)

const networkName = "network"

func (inst *DB) GetHostNetworks() ([]*amodel.Network, error) {
	var m []*amodel.Network
	if err := inst.DB.Preload("Hosts").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) GetHostNetworkByName(name string, isUUID bool) (*amodel.Network, error) {
	m := new(amodel.Network)
	switch isUUID {
	case true:
		if err := inst.DB.Where("uuid = ? ", name).Preload("Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(networkName)
		}
		return m, nil
	case false:
		if err := inst.DB.Where("name = ? ", name).Preload("Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(networkName)
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}

func (inst *DB) CreateHostNetwork(body *amodel.Network) (*amodel.Network, error) {
	if body.Name == "" {
		body.Name = "net"
	}
	existing, _ := inst.GetLocationsByName(body.Name, false)
	if existing != nil {
		return nil, errors.New("a network with this name exists")
	}
	body.UUID = uuid.ShortUUID("net")
	if err := inst.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (inst *DB) UpdateHostNetworkByName(name string, host *amodel.Network) (*amodel.Network, error) {
	m := new(amodel.Network)
	query := inst.DB.Where("name = ?", name).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(networkName)
	} else {
		return m, nil
	}
}

func (inst *DB) UpdateHostNetwork(uuid string, host *amodel.Network) (*amodel.Network, error) {
	m := new(amodel.Network)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(networkName)
	} else {
		return host, query.Error
	}
}

func (inst *DB) DeleteHostNetwork(uuid string) (*DeleteMessage, error) {
	var m *amodel.Network
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropHostNetworks() (*DeleteMessage, error) {
	var m *amodel.Network
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
