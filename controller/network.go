package controller

import (
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/gin-gonic/gin"
)

func getNetworkBody(ctx *gin.Context) (dto *model.Network, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetNetworkSchema(ctx *gin.Context) {
	mod := model.GetNetworkSchema()
	reposeHandler(mod, nil, ctx)
}

func (inst *Controller) GetHostNetwork(c *gin.Context) {
	host, err := inst.DB.GetHostNetworkByName(c.Params.ByName("uuid"), true)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) GetHostNetworks(c *gin.Context) {
	hosts, err := inst.DB.GetHostNetworks()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(hosts, err, c)
}

func (inst *Controller) CreateHostNetwork(c *gin.Context) {
	m := new(model.Network)
	err = c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateHostNetwork(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) UpdateHostNetwork(c *gin.Context) {
	body, _ := getNetworkBody(c)
	host, err := inst.DB.UpdateHostNetwork(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) DeleteHostNetwork(c *gin.Context) {
	q, err := inst.DB.DeleteHostNetwork(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropHostNetworks(c *gin.Context) {
	host, err := inst.DB.DropHostNetworks()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}
