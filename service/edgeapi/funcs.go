package edgeapi

import (
	"errors"
	"github.com/NubeIO/rubix-assist/pkg/model"
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
func (inst *Manager) getHost(body *App) (*model.Host, error, string) {
	host, err := inst.DB.GetHostByLocationName(body.HostName, body.NetworkName, body.LocationName)
	if err != nil {
		return nil, err, ""
	}
	token, _, err := inst.getTokens()
	if err != nil {
		return nil, err, ""
	}
	return host, err, token
}
