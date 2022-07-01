package edgeapi

import (
	"errors"
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

// getHost returns the host and a GitHub token
func (inst *Manager) getHost(body *AppTask) (*model.Host, error, string) {
	var host *model.Host
	var err error
	if body.HostUUID != "" {
		host, err = inst.DB.GetHost(body.HostUUID)
		if err != nil {
			return nil, err, ""
		}
	} else {
		host, err = inst.DB.GetHostByLocationName(body.HostName, body.NetworkName, body.LocationName)
		if err != nil {
			return nil, err, ""
		}
	}
	token, _, err := inst.getTokens()
	if err != nil {
		//return nil, err, ""
	}
	return host, err, token
}
