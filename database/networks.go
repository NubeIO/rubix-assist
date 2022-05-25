package dbase

import (
	"errors"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

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
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", name).Preload("Hosts").First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}

func (d *DB) CreateHostNetwork(host *model.Network) (*model.Network, error) {
	if host.Name == "" {
		host.Name = "RC"
	}
	//existingHost, _ := d.GetHostByName(host.Name, false)
	//if existingHost != nil {
	//	return nil, errors.New("a host with this name exists")
	//}
	host.UUID = config.MakeTopicUUID(model.CommonNaming.Host)
	if err := d.DB.Create(&host).Error; err != nil {
		return nil, err
	} else {
		return host, nil
	}
}

func (d *DB) UpdateHostNetworkByName(name string, host *model.Network) (*model.Network, error) {
	m := new(model.Network)
	query := d.DB.Where("name = ?", name).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return m, query.Error
	}
}

func (d *DB) UpdateHostNetwork(uuid string, host *model.Network) (*model.Network, error) {
	m := new(model.Network)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return host, query.Error
	}
}

func (d *DB) DeleteHostNetwork(uuid string) (ok bool, err error) {
	var m *model.Network
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

// DropHostNetworks delete all.
func (d *DB) DropHostNetworks() (bool, error) {
	var m *model.Network
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
