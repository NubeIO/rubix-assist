package controller

import (
	"encoding/json"
	"errors"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/wirescli"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) wiresToken(c *gin.Context) (*wirescli.Token, *model.Host, error) {
	body := &wirescli.WiresTokenBody{}
	err = c.ShouldBindJSON(body)
	host, err := inst.resolveHost(c)
	if err != nil {
		return nil, nil, err
	}
	body.Username = host.Username
	body.Password = host.Password
	data, _ := wirescli.New(host.IP, host.WiresPort).GetToken(body)
	return data, host, nil
}

func (inst *Controller) WiresUpload(c *gin.Context) {
	body := &wirescli.NodesBody{}
	var nodes interface{}
	data, _ := c.GetRawData()
	if err := json.Unmarshal(data, &nodes); err != nil {
		responseHandler(nil, err, c)
	}
	body.Nodes = nodes
	body.Pos = []float64{-1250, -1600}
	token, host, err := inst.wiresToken(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	if token.Token == "" {
		responseHandler(nil, errors.New("failed to get wires token"), c)
		return
	}
	ok, res := wirescli.New(host.IP, host.WiresPort).Upload(body, token.Token)
	r := map[string]interface{}{
		"imported": ok,
		"code":     res.StatusCode,
	}
	responseHandler(r, err, c)
}

func (inst *Controller) WiresBackup(c *gin.Context) {
	token, host, err := inst.wiresToken(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	if token.Token == "" {
		responseHandler(nil, errors.New("failed to get wires token"), c)
		return
	}
	if host.WiresPort == 0 {
		host.WiresPort = 1313
	}
	data, err := wirescli.New(host.IP, host.WiresPort).Backup(token.Token)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, err, c)
}
