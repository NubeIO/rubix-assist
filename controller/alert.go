package controller

import (
	"github.com/NubeIO/rubix-assist-model/model"
	"github.com/gin-gonic/gin"
)

func getAlertBody(ctx *gin.Context) (dto *model.Alert, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) AlertsSchema(ctx *gin.Context) {
}

func (inst *Controller) GetAlert(c *gin.Context) {
	team, err := inst.DB.GetAlert(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) GetAlerts(c *gin.Context) {
	teams, err := inst.DB.GetAlerts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(teams, err, c)
}

func (inst *Controller) CreateAlert(c *gin.Context) {
	m := new(model.Alert)
	err = c.ShouldBindJSON(&m)
	res, err := inst.DB.CreateAlert(m)
	if err != nil {
		reposeHandler(m, err, c)
		return
	}
	reposeHandler(res, err, c)
}

func (inst *Controller) UpdateAlert(c *gin.Context) {
	body, _ := getAlertBody(c)
	team, err := inst.DB.UpdateAlert(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) DeleteAlert(c *gin.Context) {
	q, err := inst.DB.DeleteAlert(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropAlerts(c *gin.Context) {
	team, err := inst.DB.DropAlerts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}
