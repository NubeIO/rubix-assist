package base

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/rubix-assist/pkg/model"
)

const hostName = "host"

func (d *DB) GetHostByLocationName(hostName, networkName, locationName string) (*model.Host, error) {
	location, err := d.GetLocationsByName(locationName, false)
	if err != nil {
		return nil, err
	}
	for _, network := range location.Networks {
		if network.Name == networkName {
			for _, host := range network.Hosts {
				if host.Name == hostName {
					return host, err
				}
			}
		}
	}
	return nil, errors.New("no host was found")
}

func (d *DB) GetHostByName(name string, isUUID bool) (*model.Host, error) {
	m := new(model.Host)
	switch isUUID {
	case true:
		if err := d.DB.Where("uuid = ? ", name).First(&m).Error; err != nil {
			return nil, handelNotFound(hostName)
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", name).First(&m).Error; err != nil {
			return nil, handelNotFound(hostName)
		}
		return m, nil
	default:
		return nil, errors.New("no valid uuid or name was provided in the request")
	}
}

func (d *DB) GetHosts() ([]*model.Host, error) {
	var m []*model.Host
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateHost(host *model.Host) (*model.Host, error) {
	if host.Name == "" {
		host.Name = "rc"
	}
	existingHost, _ := d.GetHostByName(host.Name, false)
	if existingHost != nil {
		return nil, errors.New("an existing host with this name exists")
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
	if host.RubixUsername == "" {
		host.RubixUsername = "admin"
	}
	if host.Username == "" {
		host.Username = "pi"
	}
	if host.Port == 0 {
		host.Port = 22
	}
	if host.IP == "" {
		host.IP = "0.0.0.0"
	}
	if host.RubixPort == 0 {
		host.RubixPort = 1661
	}
	if host.WiresPort == 0 {
		host.WiresPort = 1313
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
		return nil, handelNotFound(hostName)
	} else {
		return m, nil
	}
}

func (d *DB) UpdateHost(uuid string, host *model.Host) (*model.Host, error) {
	m := new(model.Host)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(hostName)
	} else {
		return m, nil
	}
}

func (d *DB) DeleteHost(uuid string) (*DeleteMessage, error) {
	var m *model.Host
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

// DropHosts delete all.
func (d *DB) DropHosts() (*DeleteMessage, error) {
	var m *model.Host
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
