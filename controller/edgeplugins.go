package controller

import (
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/gin-gonic/gin"
)

// EdgeUploadPlugin upload a plugin to the edge dev
func (inst *Controller) EdgeUploadPlugin(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *appstore.Plugin
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeUploadPlugin(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// EdgeListPlugins list all the plugins on the edge device
func (inst *Controller) EdgeListPlugins(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListPlugins(host.UUID, host.Name)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
