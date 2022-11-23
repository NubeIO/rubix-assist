package controller

import (
	"errors"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"github.com/NubeIO/rubix-assist/service/systemctl"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeAppUpload(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	var m *amodel.AppUpload
	err = c.ShouldBindJSON(&m)
	data, err := cli.AppUpload(m)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeAppInstall(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	var m *systemctl.ServiceFile
	err = c.ShouldBindJSON(&m)
	data, err := cli.AppInstall(m)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeAppUninstall(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	appName := c.Query("app_name")
	if appName == "" {
		responseHandler(nil, errors.New("app_name can't be empty"), c)
		return
	}
	data, err := cli.AppUninstall(appName)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeListAppsStatus(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	data, err := cli.AppsStatus()
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeGetInstallVersion(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	appStatus, err := cli.GetAppStatus(c.Query("app_name"))
	responseHandler(appStatus, err, c)
}
