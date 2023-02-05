package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getHostBody(ctx *gin.Context) (dto *amodel.Host, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetHostSchema(ctx *gin.Context) {
	mod := amodel.GetHostSchema()
	responseHandler(mod, nil, ctx)
}

func (inst *Controller) GetHost(c *gin.Context) {
	host, err := inst.DB.GetHost(c.Params.ByName("uuid"))
	responseHandler(host, err, c)
}

func (inst *Controller) GetHosts(c *gin.Context) {
	hosts, err := inst.DB.GetHosts(true)
	responseHandler(hosts, err, c)
}

func (inst *Controller) CreateHost(c *gin.Context) {
	m := new(amodel.Host)
	err := c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateHost(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) UpdateHost(c *gin.Context) {
	body, _ := getHostBody(c)
	host, err := inst.DB.UpdateHost(c.Params.ByName("uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) DeleteHost(c *gin.Context) {
	q, err := inst.DB.DeleteHost(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}

func (inst *Controller) DropHosts(c *gin.Context) {
	host, err := inst.DB.DropHosts()
	responseHandler(host, err, c)
}

func (inst *Controller) UpdateStatus(c *gin.Context) {
	hosts, err := inst.DB.UpdateStatus()
	responseHandler(hosts, err, c)
}

func (inst *Controller) ConfigureOpenVPN(c *gin.Context) {
	hosts, err := inst.DB.ConfigureOpenVPN(c.Params.ByName("uuid"))
	responseHandler(hosts, err, c)
}
