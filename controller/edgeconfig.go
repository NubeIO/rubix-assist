package controller

import (
	"errors"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeReadConfig(c *gin.Context) {
	appName := c.Query("app_name")
	configName := c.Query("config_name")
	if appName == "" {
		responseHandler(nil, errors.New("app_name can not be empty"), c)
		return
	}
	if configName == "" {
		configName = "config.yml"
	}
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeReadConfig(host.UUID, host.Name, appName, configName)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeWriteConfig(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *amodel.EdgeConfig
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
	}
	data, err := inst.Store.EdgeWriteConfig(host.UUID, host.Name, m)
	responseHandler(data, err, c)
}
