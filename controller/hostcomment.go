package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getHostCommentBody(ctx *gin.Context) (dto *amodel.HostComment, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) CreateHostComment(c *gin.Context) {
	m := new(amodel.HostComment)
	err := c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateHostComment(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) UpdateHostComment(c *gin.Context) {
	body, _ := getHostCommentBody(c)
	host, err := inst.DB.UpdateHostComment(c.Params.ByName("uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) DeleteHostComment(c *gin.Context) {
	q, err := inst.DB.DeleteHostComment(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}
