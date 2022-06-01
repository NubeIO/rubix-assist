package controller

import (
	"github.com/NubeIO/rubix-assist-model/model"
	"github.com/gin-gonic/gin"
)

func getLocationBody(ctx *gin.Context) (dto *model.Location, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetLocation(c *gin.Context) {
	host, err := inst.DB.GetLocationsByName(c.Params.ByName("uuid"), true)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) GetLocations(c *gin.Context) {
	hosts, err := inst.DB.GetLocations()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(hosts, err, c)
}

func (inst *Controller) CreateLocation(c *gin.Context) {
	m := new(model.Location)
	err = c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateLocation(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) UpdateLocation(c *gin.Context) {
	body, _ := getLocationBody(c)
	host, err := inst.DB.UpdateLocation(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (inst *Controller) DeleteLocation(c *gin.Context) {
	q, err := inst.DB.DeleteLocation(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropLocations(c *gin.Context) {
	host, err := inst.DB.DropLocations()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}
