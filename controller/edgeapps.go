package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/systemctl"
	"github.com/gin-gonic/gin"
	"strconv"
)

// EdgeListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Controller) EdgeListApps(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListApps(host.UUID, host.Name)
	responseHandler(data, err, c)
}

// EdgeListAppsStatus get all the apps status
func (inst *Controller) EdgeListAppsStatus(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListAppsStatus(host.UUID, host.Name)
	responseHandler(data, err, c)
}

// EdgeUploadApp uploads the build
func (inst *Controller) EdgeUploadApp(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *installer.Upload
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeUploadApp(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}

func (inst *Controller) GenerateServiceFileAndEdgeUpload(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *systemctl.ServiceFile
	err = c.ShouldBindJSON(&m)

	data, err := inst.Store.GenerateServiceFileAndEdgeUpload(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeInstallService(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *installer.Install
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeInstallService(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}

// EdgeUninstallApp full uninstallation of an app
func (inst *Controller) EdgeUninstallApp(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	deleteApp, _ := strconv.ParseBool(c.Query("delete"))
	data, err := inst.Store.EdgeUninstallApp(host.UUID, host.Name, c.Query("name"), deleteApp)
	responseHandler(data, err, c)
}
