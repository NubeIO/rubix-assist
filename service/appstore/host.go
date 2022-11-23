package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
)

func (inst *Store) getClient(hostUUID, hostName string) (*edgecli.Client, error) {
	host, err := inst.getHost(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return inst.newClient(host)
}

func (inst *Store) newClient(host *amodel.Host) (*edgecli.Client, error) {
	cli := edgecli.New(&edgecli.Client{
		Rest:          nil,
		Ip:            host.IP,
		Port:          host.Port,
		HTTPS:         host.HTTPS,
		ExternalToken: host.ExternalToken,
	})
	return cli, nil
}

func matchUUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "hos_" {
			return true
		}
	}
	return false
}

func (inst *Store) getHost(hostUUID, hostName string) (*amodel.Host, error) {
	var host *amodel.Host
	var err error
	if matchUUID(hostUUID) {
		host, err = inst.DB.GetHost(hostUUID)
		if err != nil {
			return nil, err
		}
	} else {
		host, err = inst.DB.GetHostByName(hostName)
		if err != nil {
			var hostNames []string
			var hostUUIDs []string
			var count int
			hosts, _ := inst.DB.GetHosts()
			for _, h := range hosts {
				hostNames = append(hostNames, h.Name)
				hostUUIDs = append(hostUUIDs, h.UUID)
				count++
			}
			return nil, errors.New(fmt.Sprintf("no valid host was found: host count: %d, host names found: %v uuids: %v", count, hostNames, hostUUIDs))
		}
	}
	return host, nil
}
