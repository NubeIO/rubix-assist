package controller

import (
	"errors"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeWriteConfig(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *appstore.EdgeConfig
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeWriteConfig(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}
func (inst *Controller) EdgeReadConfig(c *gin.Context) {
	appName := c.Query("name")
	configName := c.Query("config")
	if appName == "" {
		responseHandler(nil, errors.New("file path can not be empty"), c)
		return
	}
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeReadConfig(host.UUID, host.Name, appName, configName)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}
