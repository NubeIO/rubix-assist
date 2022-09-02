package controller

import (
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/gin-gonic/gin"
)

// EdgeUploadPlugin upload a plugin to the edge dev
func (inst *Controller) EdgeUploadPlugin(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *appstore.Plugin
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeUploadPlugin(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

// EdgeListPlugins list all the plugins on the edge device
func (inst *Controller) EdgeListPlugins(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListPlugins(host.UUID, host.Name)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

// EdgeDeletePlugin list all the plugins on the edge device
func (inst *Controller) EdgeDeletePlugin(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *appstore.Plugin
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeDeletePlugin(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

// EdgeDeleteAllPlugins list all the plugins on the edge device
func (inst *Controller) EdgeDeleteAllPlugins(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeDeleteAllPlugins(host.UUID, host.Name)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}
