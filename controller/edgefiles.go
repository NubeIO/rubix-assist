package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeWalkFile(c *gin.Context) {
	path := c.Query("path")
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeWalkFiles(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeListFiles(c *gin.Context) {
	path := c.Query("path")
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeListFiles(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeRenameFile(c *gin.Context) {
	oldName := c.Query("old")
	newName := c.Query("new")
	if oldName == "" || newName == "" {
		reposeHandler(nil, errors.New("old and new, from and to files name can not be empty"), c)
		return
	}
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeRenameFile(host.UUID, host.Name, oldName, newName)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeCopyFile(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if from == "" || to == "" {
		reposeHandler(nil, errors.New("from and to, from and to files name can not be empty"), c)
		return
	}
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeRenameFile(host.UUID, host.Name, from, to)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeMoveFile(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if from == "" || to == "" {
		reposeHandler(nil, errors.New("from and to, from and to files name can not be empty"), c)
		return
	}
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeMoveFile(host.UUID, host.Name, from, to)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeDeleteFile(c *gin.Context) {
	path := c.Query("path")
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeDeleteFile(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) EdgeDeleteAllFiles(c *gin.Context) {
	path := c.Query("path")
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.EdgeDeleteAllFiles(host.UUID, host.Name, path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
