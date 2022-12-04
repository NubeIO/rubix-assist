package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (inst *Controller) EdgeListPlugins(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	plugins, connectionErr, requestErr := cli.ListPlugins()
	if connectionErr != nil {
		c.JSON(502, amodel.Message{Message: connectionErr.Error()})
		return
	}
	responseHandler(plugins, requestErr, c)
}

func (inst *Controller) EdgeUploadPlugin(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	var m *amodel.Plugin
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := cli.PluginUpload(m)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeMoveFromDownloadToInstallPlugins(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	resp, err := cli.MovePluginsFromDownloadToInstallDir()
	responseHandler(resp, err, c)
}

func (inst *Controller) EdgeDeletePlugin(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	pluginName := c.Param("plugin_name")
	arch := c.Query("arch")
	cli := cligetter.GetEdgeClient(host)
	installPluginFilePath := global.Installer.GetAppPluginInstallFilePath(pluginName, arch)
	_, connectionErr, requestErr := cli.DeleteFiles(installPluginFilePath)
	if connectionErr != nil {
		log.Errorf(connectionErr.Error())
		c.JSON(502, amodel.Message{Message: connectionErr.Error()})
		return
	}
	if requestErr != nil {
		responseHandler(nil, requestErr, c)
		log.Errorf(requestErr.Error())
		return
	}
	responseHandler(amodel.Message{Message: fmt.Sprintf("successfully deleted plugin %s", pluginName)}, nil, c)
}

func (inst *Controller) EdgeDeleteDownloadPlugins(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	pluginDownloadPath := global.Installer.GetAppPluginDownloadPath()
	msg, connectionErr, requestErr := cli.DeleteFiles(pluginDownloadPath)
	if connectionErr != nil {
		c.JSON(502, amodel.Message{Message: connectionErr.Error()})
		return
	}
	responseHandler(msg, requestErr, c)
}
