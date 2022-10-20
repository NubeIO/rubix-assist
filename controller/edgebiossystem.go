package controller

import (
	"github.com/NubeIO/rubix-assist/helpers"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeBiosPing(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := helpers.GetEdgeBiosClient(host)
	data, err := cli.Ping()
	responseHandler(data, err, c)
}
