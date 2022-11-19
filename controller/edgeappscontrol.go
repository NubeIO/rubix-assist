package controller

import (
	"github.com/NubeIO/rubix-assist/model"
	"github.com/gin-gonic/gin"
)

// EdgeSystemCtlAction start, stop, enable, disable a service
func (inst *Controller) EdgeSystemCtlAction(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *model.SystemCtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeSystemCtlAction(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}

// EdgeSystemCtlStatus check isRunning, isInstalled, isEnabled, isActive, isFailed for a service
func (inst *Controller) EdgeSystemCtlStatus(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *model.SystemCtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeSystemCtlStatus(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}

// EdgeServiceMassAction start, stop, enable, disable a service
func (inst *Controller) EdgeServiceMassAction(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *model.SystemCtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeServiceMassAction(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}

// EdgeServiceMassStatus on mass check isRunning, isInstalled, isEnabled, isActive, isFailed for a service
func (inst *Controller) EdgeServiceMassStatus(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *model.SystemCtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeServiceMassStatus(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}
