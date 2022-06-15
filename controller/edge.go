package controller

import (
	"github.com/NubeIO/rubix-assist/service/edgeapi"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) InstallApp(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.RunInstall(m)
	if err != nil {
		reposeWithCode(404, data, nil, c)
		return
	}
	reposeWithCode(202, data, nil, c)
}

func (inst *Controller) TaskBuilder(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.TaskBuilder(m)
	if err != nil {
		reposeWithCode(404, err, nil, c)
		return
	}
	reposeWithCode(202, data, nil, c)
}

func (inst *Controller) TaskRunner(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.TaskRunner(m)
	if err != nil {
		reposeWithCode(404, data, nil, c)
		return
	}
	reposeWithCode(202, data, nil, c)
}
