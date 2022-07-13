package edgeapi

import (
	"errors"
	"fmt"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
)

func (inst *Manager) getTokens() (token string, tokens []*model.Token, err error) {
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
func (inst *Manager) getHost(body *AppTask) (*model.Host, error, string) {
	var host *model.Host
	var err error
	fmt.Println(body.HostUUID, body.HostName)
	if matchUUID(body.HostUUID) {
		host, err = inst.DB.GetHost(body.HostUUID)
		if err != nil {
			return nil, err, ""
		}
	} else {
		host, err = inst.DB.GetHostByName(body.HostName)
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
			return nil, errors.New(fmt.Sprintf("no valid host was found: host count:%d, host names found:%v uuids:%v", count, hostNames, hostUUIDs)), ""
		}
	}
	token, _, err := inst.getTokens()
	if err != nil {
		//return nil, err, ""
	}
	return host, nil, token
}
