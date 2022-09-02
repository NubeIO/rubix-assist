package controller

import (
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) AssistPing(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data := inst.Store.AssistPing(host.UUID, host.Name)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

func (inst *Controller) EdgePing(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *edgecli.PingBody
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgePing(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}
