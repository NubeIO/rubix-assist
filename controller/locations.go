package controller

import (
	"github.com/NubeIO/rubix-assist/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetLocationSchema(ctx *gin.Context) {
	mod := model.GetLocationSchema()
	responseHandler(mod, nil, ctx)
}

func getLocationBody(ctx *gin.Context) (dto *model.Location, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetLocation(c *gin.Context) {
	host, err := inst.DB.GetLocationsByName(c.Params.ByName("uuid"), true)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) GetLocations(c *gin.Context) {
	hosts, err := inst.DB.GetLocations()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(hosts, err, c)
}

func (inst *Controller) CreateLocationWizard(c *gin.Context) {
	m := new(model.Location)
	err := c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateLocationWizard(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) CreateLocation(c *gin.Context) {
	m := new(model.Location)
	err := c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateLocation(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) UpdateLocation(c *gin.Context) {
	body, _ := getLocationBody(c)
	host, err := inst.DB.UpdateLocation(c.Params.ByName("uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}

func (inst *Controller) DeleteLocation(c *gin.Context) {
	q, err := inst.DB.DeleteLocation(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
	} else {
		responseHandler(q, err, c)
	}
}

func (inst *Controller) DropLocations(c *gin.Context) {
	host, err := inst.DB.DropLocations()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(host, err, c)
}
