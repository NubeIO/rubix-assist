package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/store"
	"github.com/gin-gonic/gin"
)

// AddUploadEdgeApp
// upload the build
func (inst *Controller) AddUploadEdgeApp(c *gin.Context) {
	var m *store.EdgeApp
	err = c.ShouldBindJSON(&m)
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.AddUploadEdgeApp(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) GenerateUploadEdgeService(c *gin.Context) {
	var m *store.ServiceFile
	err = c.ShouldBindJSON(&m)
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.GenerateUploadEdgeService(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) InstallEdgeService(c *gin.Context) {
	var m *installer.Install
	err = c.ShouldBindJSON(&m)
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.InstallEdgeService(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
