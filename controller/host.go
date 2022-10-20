package controller

import (
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/gin-gonic/gin"
)

func getHostBody(ctx *gin.Context) (dto *model.Host, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetHostSchema(ctx *gin.Context) {
	mod := model.GetHostSchema()
	responseHandler(mod, nil, ctx)
}

func (inst *Controller) GetHost(c *gin.Context) {
	host, err := inst.DB.GetHost(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) GetHosts(c *gin.Context) {
	hosts, err := inst.DB.GetHosts()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(hosts, err, c)
}

func (inst *Controller) CreateHost(c *gin.Context) {
	m := new(model.Host)
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
	if err != nil {
		responseHandler(nil, err, c)
	} else {
		responseHandler(q, err, c)
	}
}

func (inst *Controller) DropHosts(c *gin.Context) {
	host, err := inst.DB.DropHosts()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}
