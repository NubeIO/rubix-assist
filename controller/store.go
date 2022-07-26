package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/store"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) ListStore(c *gin.Context) {
	data, err := inst.Store.ListStore()
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
		File:    file,
	}
	data, err := inst.Store.AddUploadStoreApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) CheckStoreApp(c *gin.Context) {
	m := &store.App{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.CheckApp(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) CheckStoreApps(c *gin.Context) {
	var m []store.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.CheckApps(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
