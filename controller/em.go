package controller

import (
	"github.com/NubeIO/rubix-assist/service/em"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) InstallApp(c *gin.Context) {

	m := &em.App{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.InstallApp(m)
	if err != nil {
		reposeWithCode(404, err, nil, c)
		return
	}
	reposeWithCode(202, data, nil, c)
}
