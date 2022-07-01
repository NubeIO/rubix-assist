package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
)

const networkName = "network"

func (d *DB) GetHostNetworks() ([]*model.Network, error) {
	var m []*model.Network
	if err := d.DB.Preload("Hosts").Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) GetHostNetworkByName(name string, isUUID bool) (*model.Network, error) {
	m := new(model.Network)
	switch isUUID {
	case true:
		if err := d.DB.Where("uuid = ? ", name).Preload("Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(networkName)
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", name).Preload("Hosts").First(&m).Error; err != nil {
			return nil, handelNotFound(networkName)
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}

func (d *DB) CreateHostNetwork(body *model.Network) (*model.Network, error) {
	if body.Name == "" {
		body.Name = "net"
	}
	existing, _ := d.GetLocationsByName(body.Name, false)
	if existing != nil {
		return nil, errors.New("a network with this name exists")
	}
	body.UUID = uuid.ShortUUID("net")
	if err := d.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (d *DB) UpdateHostNetworkByName(name string, host *model.Network) (*model.Network, error) {
	m := new(model.Network)
	query := d.DB.Where("name = ?", name).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(networkName)
	} else {
		return m, nil
	}
}

func (d *DB) UpdateHostNetwork(uuid string, host *model.Network) (*model.Network, error) {
	m := new(model.Network)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(networkName)
	} else {
		return host, query.Error
	}
}

func (d *DB) DeleteHostNetwork(uuid string) (*DeleteMessage, error) {
	var m *model.Network
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

// DropHostNetworks delete all.
func (d *DB) DropHostNetworks() (*DeleteMessage, error) {
	var m *model.Network
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
