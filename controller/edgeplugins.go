package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeUploadPlugin(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	var m *amodel.Plugin
	err = c.ShouldBindJSON(&m)
	data, err := cli.PluginUpload(m)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeListPlugins(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := cligetter.GetEdgeClient(host)
	plugins, err := cli.ListPlugins()
	responseHandler(plugins, err, c)
}
