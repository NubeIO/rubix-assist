package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
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

func (inst *Controller) EdgeFileExists(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeFileExists(host.UUID, host.Name, path)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeReadFile(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeReadFile(host.UUID, host.Name, path)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	c.Data(http.StatusOK, ContentTypeText, data)
}

func (inst *Controller) EdgeCreateFile(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *assistmodel.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeCreateFile(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

func (inst *Controller) EdgeWriteString(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *assistmodel.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeWriteString(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

func (inst *Controller) EdgeWriteFileJson(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *assistmodel.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeWriteFileJson(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

func (inst *Controller) EdgeWriteFileYml(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *assistmodel.WriteFile
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeWriteFileYml(host.UUID, host.Name, m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}
