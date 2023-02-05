package base

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"sync"
)

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
	err := inst.DB.Where("uuid = ? ", uuid).Preload("Hosts").First(&m).Error
	if err != nil {
		return nil, err
	}
	inst.attachOpenVPN(m.Hosts)
	return &m, nil
}

func (inst *DB) CreateHostNetwork(body *amodel.Network) (*amodel.Network, error) {
	body.UUID = uuid.ShortUUID("net")
	err := inst.DB.Create(&body).Error
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (inst *DB) UpdateHostNetwork(uuid string, host *amodel.Network) (*amodel.Network, error) {
	m := amodel.Network{}
	err := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(host).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (inst *DB) UpdateHostsStatus(uuid string) (*amodel.Network, error) {
	m := amodel.Network{}
	err := inst.DB.Where("uuid = ?", uuid).Preload("Hosts").Find(&m).Error
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	for _, host := range m.Hosts {
		wg.Add(1)
		cli := cligetter.GetEdgeClientFastTimeout(host)
		go func(h *amodel.Host) {
			defer wg.Done()
			globalUUID, pingable, isValidToken := cli.Ping()
			if globalUUID != nil {
				h.GlobalUUID = *globalUUID
			}
			h.IsOnline = &pingable
			h.IsValidToken = &isValidToken
		}(host)
	}
	wg.Wait()
	tx := inst.DB.Begin()
	for _, host := range m.Hosts {
		if err := tx.Where("uuid = ?", host.UUID).Updates(&host).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return &m, nil
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
