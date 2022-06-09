package controller

import (
	"github.com/NubeIO/rubix-assist/service/edge"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) InstallApp(c *gin.Context) {
	m := &edge.App{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.RunAppInstall(m)
	if err != nil {
		reposeWithCode(404, err, nil, c)
		return
	}
	reposeWithCode(202, data, nil, c)
}

func (inst *Controller) InstallPipeline(c *gin.Context) {
	m := &edge.App{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.PipeRunner(m)
	if err != nil {
		reposeWithCode(404, err, nil, c)
		return
	}
	reposeWithCode(202, data, nil, c)
}
