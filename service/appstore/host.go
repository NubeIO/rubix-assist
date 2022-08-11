package appstore

import (
	"errors"
	"fmt"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
)

func (inst *Store) getClient(hostUUID, hostName string) (*edgecli.Client, error) {

	host, err := inst.getHost(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	ip := host.IP
	port := host.Port
	return inst.newClient(ip, host.Username, host.Password, port)
}

func (inst *Store) newClient(ip, user, password string, port int) (*edgecli.Client, error) {
	cli := edgecli.New(ip, port)
	return cli, nil
}

func (inst *Store) getTokens() (token string, tokens []*model.Token, err error) {
	tokens = []*model.Token{}
	tokens, err = inst.DB.GetTokens()
	if err != nil {
		return "", nil, errors.New("no token provided")
	}
	if len(tokens) == 0 {
		return "", nil, errors.New("no token provided")
	}
	return tokens[0].Token, tokens, nil
}

func matchUUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "hos_" {
			return true
		}
	}
	return false
}

// getHost returns the host and a GitHub token
func (inst *Store) getHost(hostUUID, hostName string) (*model.Host, error) {
	var host *model.Host
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
			return nil, errors.New(fmt.Sprintf("no valid host was found: host count:%d, host names found:%v uuids:%v", count, hostNames, hostUUIDs))
		}
	}
	return host, nil
}
