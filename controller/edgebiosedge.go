package controller

import (
	"github.com/NubeIO/rubix-assist/helpers"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeBiosEdgeUpload(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *assistmodel.PingBody
	err = c.ShouldBindJSON(&m)
	cli := helpers.GetEdgeBiosClient(host)
	data, err := cli.Upload("", "", nil) // TODO
	responseHandler(data, err, c)
}
