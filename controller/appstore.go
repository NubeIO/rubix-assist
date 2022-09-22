package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) ListAppsWithVersions(c *gin.Context) {
	data, err := inst.Store.ListAppsWithVersions()
	responseHandler(data, err, c)
}

func (inst *Controller) UploadAddOnAppStore(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		Arch:    c.Query("arch"),
		File:    file,
	}
	data, err := inst.Store.UploadAddOnAppStore(m)
	responseHandler(data, err, c)
}
