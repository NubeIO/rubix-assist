package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getHostTagsBody(ctx *gin.Context) (dto []*amodel.HostTag, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) UpdateHostTags(c *gin.Context) {
	body, _ := getHostTagsBody(c)
	host, err := inst.DB.UpdateHostTags(c.Params.ByName("host_uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}
