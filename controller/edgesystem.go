package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgePing(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *assistmodel.PingBody
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgePing(host.UUID, host.Name)
	responseHandler(data, err, c)
}

// EdgeGetDeviceInfo get edge details
func (inst *Controller) EdgeGetDeviceInfo(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeGetDeviceInfo(host.UUID, host.Name)
	responseHandler(data, err, c)
}

// EdgeProductInfo get edge details
func (inst *Controller) EdgeProductInfo(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeProductInfo(host.UUID, host.Name)
	responseHandler(data, err, c)
}
