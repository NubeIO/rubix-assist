package controller

import (
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func (inst *Controller) EdgeReadFile(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeReadFile(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	c.Data(http.StatusOK, ContentTypeText, data)
}

func (inst *Controller) EdgeWriteFile(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *edgecli.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeWriteFile(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeWriteFileJson(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *edgecli.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeWriteFileJson(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeFileExists(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeFileExists(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeDirExists(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeDirExists(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
func (inst *Controller) EdgeWriteFileYml(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *edgecli.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeWriteFileYml(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeCreateFile(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *edgecli.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeCreateFile(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeCreateDir(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeCreateDir(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
