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

// UploadApp
// upload the build
func (inst *Controller) UploadApp(c *gin.Context) {
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
	data, err := inst.Store.UploadApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) AddApp(c *gin.Context) {
	m := &store.App{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.AddApp(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) CheckApp(c *gin.Context) {
	m := &store.App{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.CheckApp(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) CheckApps(c *gin.Context) {
	var m []store.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.CheckApps(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
