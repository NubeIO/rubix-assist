package base

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ip"
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

func (d *DB) GetHost(uuid string) (*model.Host, error) {
	m := new(model.Host)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no host was found with uuid:%s", uuid))
	}
	return m, nil
}

func (d *DB) GetHostByName(name string) (*model.Host, error) {
	m := new(model.Host)
	if err := d.DB.Where("name = ? ", name).First(&m).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no host was found with name:%s", name))
	}
	return m, nil
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
	if len(host.Name) < 1 {
		return nil, errors.New("host name length must be grater then two")
	}
	existingHost, _ := d.GetHostByName(host.Name)
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
	if host.Username == "" {
		host.Username = "admin"
	}
	if host.Password == "" {
		host.Password = "admin"
	}
	if host.IP == "" {
		host.IP = "0.0.0.0"
	}
	if host.Port == 0 {
		host.Port = 1661
	}
	if host.WiresPort == 0 {
		host.WiresPort = 1313
	}
	err := ip.CheckURL(host.IP, host.Port)
	if err != nil {
		return nil, fmt.Errorf("invaild ssh ip:port:%s", err.Error())
	}
	err = ip.CheckURL(host.IP, host.Port)
	if err != nil {
		return nil, fmt.Errorf("invaild rubix ip:port:%s", err.Error())
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
