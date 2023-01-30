package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetLocationSchema(ctx *gin.Context) {
	mod := amodel.GetLocationSchema()
	responseHandler(mod, nil, ctx)
}

func getLocationBody(ctx *gin.Context) (dto *amodel.Location, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetLocations(c *gin.Context) {
	locations, err := inst.DB.GetLocations()
	responseHandler(locations, err, c)
}

func (inst *Controller) GetLocation(c *gin.Context) {
	location, err := inst.DB.GetLocation(c.Params.ByName("uuid"))
	responseHandler(location, err, c)
}

func (inst *Controller) CreateLocation(c *gin.Context) {
	m := new(amodel.Location)
	err := c.ShouldBindJSON(&m)
	location, err := inst.DB.CreateLocation(m)
	responseHandler(location, err, c)
}

func (inst *Controller) UpdateLocation(c *gin.Context) {
	body, err := getLocationBody(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	location, err := inst.DB.UpdateLocation(c.Params.ByName("uuid"), body)
	responseHandler(location, err, c)
}

func (inst *Controller) DeleteLocation(c *gin.Context) {
	q, err := inst.DB.DeleteLocation(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}

func (inst *Controller) DropLocations(c *gin.Context) {
	location, err := inst.DB.DropLocations()
	responseHandler(location, err, c)
}
