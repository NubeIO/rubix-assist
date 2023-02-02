package base

import (
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

func (inst *DB) GetHostNetwork(uuid string) (*amodel.Network, error) {
	m := amodel.Network{}
	if err := inst.DB.Where("uuid = ? ", uuid).Preload("Hosts").First(&m).Error; err != nil {
		return nil, handelNotFound(networkName)
	}
	inst.attachOpenVPN(m.Hosts)
	return &m, nil
}

func (inst *DB) CreateHostNetwork(body *amodel.Network) (*amodel.Network, error) {
	body.UUID = uuid.ShortUUID("net")
	if err := inst.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (inst *DB) UpdateHostNetwork(uuid string, host *amodel.Network) (*amodel.Network, error) {
	m := new(amodel.Network)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, handelNotFound(networkName)
	} else {
		return host, nil
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
