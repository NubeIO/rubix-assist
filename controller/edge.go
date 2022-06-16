package controller

import (
	"github.com/NubeIO/rubix-assist/service/edgeapi"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) InstallApp(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, _ := inst.Edge.RunInstall(m)
	reposeHandler(data, err, c)
}

func (inst *Controller) TaskBuilder(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	data, _ := inst.Edge.TaskBuilder(m)
	reposeHandler(data, err, c)
}

func (inst *Controller) TaskRunner(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.TaskRunner(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
