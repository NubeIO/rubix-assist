package base

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"sync"
)

const hostName = "host"

func (inst *DB) GetHost(uuid string) (*amodel.Host, error) {
	host := amodel.Host{}
	if err := inst.DB.Where("uuid = ? ", uuid).First(&host).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no host was found with uuid: %s", uuid))
	}
	return &host, nil
}

func (inst *DB) GetHostByName(name string) (*amodel.Host, error) {
	host := amodel.Host{}
	if err := inst.DB.Where("name = ? ", name).First(&host).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no host was found with name: %s", name))
	}
	return &host, nil
}

func (inst *DB) GetHosts(withOpenVPN bool) ([]*amodel.Host, error) {
	resetHostClient := func(host *amodel.Host) {
		host.VirtualIP = ""
		host.ReceivedBytes = 0
		host.SentBytes = 0
		host.ConnectedSince = ""
	}

	resetHostsClient := func(hosts []*amodel.Host) {
		for _, host := range hosts {
			resetHostClient(host)
		}
	}

	var hosts []*amodel.Host
	if err := inst.DB.Find(&hosts).Error; err != nil {
		return nil, err
	}
	if withOpenVPN {
		oCli, _ := cligetter.GetOpenVPNClient()
		if oCli != nil {
			clients, _ := oCli.GetClients()
			if clients != nil {
				for _, host := range hosts {
					if client, found := (*clients)[host.GlobalUUID]; found {
						host.VirtualIP = client.VirtualIP
						host.ReceivedBytes = client.ReceivedBytes
						host.SentBytes = client.SentBytes
						host.ConnectedSince = client.ConnectedSince
					} else {
						resetHostClient(host)
					}
				}
			} else {
				resetHostsClient(hosts)
			}
		} else {
			resetHostsClient(hosts)
		}
	}
	return hosts, nil
}

func (inst *DB) CreateHost(host *amodel.Host) (*amodel.Host, error) {
	host.UUID = uuid.ShortUUID("hos")
	if host.HTTPS == nil {
		host.HTTPS = nils.NewFalse()
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

func (inst *DB) UpdateStatus() ([]*amodel.Host, error) {
	var hosts []*amodel.Host
	if err := inst.DB.Find(&hosts).Error; err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	for _, host := range hosts {
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
	for _, host := range hosts {
		if err := tx.Where("uuid = ?", host.UUID).Updates(&host).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return hosts, nil
}

func (inst *DB) ConfigureOpenVPN(uuid string) (*amodel.Message, error) {
	host := amodel.Host{}
	if err := inst.DB.Where("uuid = ? ", uuid).First(&host).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no host was found with uuid: %s", uuid))
	}
	cli := cligetter.GetEdgeClient(&host)
	globalUUID, pingable, isValidToken := cli.Ping()
	if pingable == false {
		return nil, errors.New("make it accessible at first")
	}
	if isValidToken == false || globalUUID == nil {
		return nil, errors.New("configure valid token at first")
	}
	host.GlobalUUID = *globalUUID
	oCli, err := cligetter.GetOpenVPNClient()
	if err != nil {
		return nil, err
	}
	openVPNConfig, err := oCli.GetOpenVPNConfig(host.GlobalUUID)
	if err != nil {
		return nil, err
	}
	_, err = cli.ConfigureOpenVPN(openVPNConfig)
	if err != nil {
		return nil, err
	}
	host.IsOnline = &pingable
	host.IsValidToken = &isValidToken
	if err := inst.DB.Where("uuid = ?", host.UUID).Updates(&host).Error; err != nil {
		return nil, err
	}
	return &amodel.Message{Message: "OpenVPN is configured!"}, nil
}
