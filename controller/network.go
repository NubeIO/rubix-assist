package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getNetworkBody(ctx *gin.Context) (dto *amodel.Network, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetNetworkSchema(ctx *gin.Context) {
	mod := amodel.GetNetworkSchema()
	responseHandler(mod, nil, ctx)
}

func (inst *Controller) GetHostNetworks(c *gin.Context) {
	hosts, err := inst.DB.GetHostNetworks()
	responseHandler(hosts, err, c)
}

func (inst *Controller) GetHostNetwork(c *gin.Context) {
	host, err := inst.DB.GetHostNetwork(c.Params.ByName("uuid"))
	responseHandler(host, err, c)
}

func (inst *Controller) CreateHostNetwork(c *gin.Context) {
	m := new(amodel.Network)
	err := c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateHostNetwork(m)
	responseHandler(host, err, c)
}

func (inst *Controller) UpdateHostNetwork(c *gin.Context) {
	body, _ := getNetworkBody(c)
	host, err := inst.DB.UpdateHostNetwork(c.Params.ByName("uuid"), body)
	responseHandler(host, err, c)
}

func (inst *Controller) DeleteHostNetwork(c *gin.Context) {
	q, err := inst.DB.DeleteHostNetwork(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}

func (inst *Controller) DropHostNetworks(c *gin.Context) {
	host, err := inst.DB.DropHostNetworks()
	responseHandler(host, err, c)
}
