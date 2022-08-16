package controller

import (
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeReplaceConfig(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var m *appstore.EdgeReplaceConfig
	err = c.ShouldBindJSON(&m)
	data, err := inst.Store.EdgeReplaceConfig(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
