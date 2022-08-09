package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) ListStore(c *gin.Context) {
	data, err := inst.Store.ListApps()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

// AddUploadStoreApp *
func (inst *Controller) AddUploadStoreApp(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		Product: c.Query("product"),
		Arch:    c.Query("arch"),
		File:    file,
	}
	data, err := inst.Store.AddUploadStoreApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
