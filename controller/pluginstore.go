package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

// GetPluginsStorePlugins list all the plugins
func (inst *Controller) GetPluginsStorePlugins(c *gin.Context) {
	data, err := inst.Store.GetPluginsStorePlugins()
	responseHandler(data, err, c)
}

// UploadPluginStorePlugin upload a plugin
func (inst *Controller) UploadPluginStorePlugin(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	m := &amodel.Upload{
		File: file,
	}
	data, err := inst.Store.UploadPluginStorePlugin(m)
	responseHandler(data, err, c)
}
