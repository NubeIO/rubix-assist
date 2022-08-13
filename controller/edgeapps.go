package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

// EdgeListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Controller) EdgeListApps(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListApps(host.UUID, host.Name)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// EdgeListAppsAndService get all the apps by listed in the installation (/data/rubix-service/apps/install) dir and then check the service
func (inst *Controller) EdgeListAppsAndService(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListAppsAndService(host.UUID, host.Name)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// EdgeListNubeServices list all the services by filtering all the service files with name nubeio
func (inst *Controller) EdgeListNubeServices(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListNubeServices(host.UUID, host.Name)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// AddUploadEdgeApp
// upload the build
func (inst *Controller) AddUploadEdgeApp(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *appstore.EdgeApp
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.AddUploadEdgeApp(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) GenerateUploadEdgeService(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *appstore.ServiceFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.GenerateUploadEdgeService(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) InstallEdgeService(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *installer.Install
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.InstallEdgeService(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// EdgeUninstallApp full uninstallation of an app
func (inst *Controller) EdgeUninstallApp(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	deleteApp, _ := strconv.ParseBool(c.Query("delete"))
	data, err := inst.Store.EdgeUnInstallApp(host.UUID, host.Name, c.Query("name"), c.Query("service"), deleteApp)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
