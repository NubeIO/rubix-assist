package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) ListAppsWithVersions(c *gin.Context) {
	data, err := inst.Store.ListAppsWithVersions()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) ListAppsBuildDetails(c *gin.Context) {
	data, err := inst.Store.ListAppsBuildDetails()
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

// UploadStorePlugin upload a plugin
func (inst *Controller) UploadStorePlugin(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		File: file,
	}
	data, err := inst.Store.UploadStorePlugin(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
