package controller

import (
	"github.com/gin-gonic/gin"
)

func (inst *Controller) AssistPing(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data := inst.Store.AssistPing(host.UUID, host.Name)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}
