package base

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/rubix-assist/amodel"
)

const hostName = "host"

func (inst *DB) GetHostByLocationName(hostName, networkName, locationName string) (*amodel.Host, error) {
	location, err := inst.GetLocationsByName(locationName, false)
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

func (inst *DB) getHost(uuid string) (*amodel.Host, error) {
	m := new(amodel.Host)
	if err := inst.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no host was found with uuid: %s", uuid))
	}
	return m, nil
}

func matchUUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "hos_" {
			return true
		}
	}
	return false
}

func (inst *DB) GetHost(uuid string) (*amodel.Host, error) {
	match := matchUUID(uuid)
	var host *amodel.Host
	if match {
		host, _ = inst.getHost(uuid)
		if host != nil {
			return host, nil
		}
	} else {
		host, _ = inst.GetHostByName(uuid)
		if host != nil {
			return host, nil
		}
	}
	return host, nil
}

func (inst *DB) GetHostByName(name string) (*amodel.Host, error) {
	m := new(amodel.Host)
	if err := inst.DB.Where("name = ? ", name).First(&m).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no host was found with name: %s", name))
	}
	return m, nil
}

func (inst *DB) GetHosts() ([]*amodel.Host, error) {
	var m []*amodel.Host
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) CreateHost(host *amodel.Host) (*amodel.Host, error) {
	if host.Name == "" {
		host.Name = "rc"
	}
	if len(host.Name) < 1 {
		return nil, errors.New("host name length must be grater then two")
	}
	existingHost, _ := inst.GetHostByName(host.Name)
	if existingHost != nil {
		return nil, errors.New("an existing host with this name exists")
	}
	host.UUID = uuid.ShortUUID("hos")
	if host.PingEnable == nil {
		host.PingEnable = nils.NewTrue()
	}
	if host.HTTPS == nil {
		host.HTTPS = nils.NewFalse()
	}
	if host.IP == "" {
		host.IP = "0.0.0.0"
	}
	if host.Port == 0 {
		host.Port = 1661
	}
	if err := inst.DB.Create(&host).Error; err != nil {
		return nil, err
	}
	return host, nil
}

func (inst *DB) UpdateHostByName(name string, host *amodel.Host) (*amodel.Host, error) {
	m := new(amodel.Host)
	query := inst.DB.Where("name = ?", name).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(hostName)
	}
	return m, nil
}

func (inst *DB) UpdateHost(uuid string, host *amodel.Host) (*amodel.Host, error) {
	m := new(amodel.Host)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(hostName)
	}
	return m, nil
}

func (inst *DB) DeleteHost(uuid string) (*DeleteMessage, error) {
	var m *amodel.Host
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropHosts() (*DeleteMessage, error) {
	var m *amodel.Host
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
