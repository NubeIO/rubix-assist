package controller

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/schema"
	"github.com/gin-gonic/gin"
)

func getAlertBody(ctx *gin.Context) (dto *model.Alert, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) AlertsSchema(ctx *gin.Context) {
	reposeHandler(schema.GetAlertSchema(), nil, ctx)
}

func (base *Controller) GetAlert(c *gin.Context) {
	team, err := base.DB.GetAlert(c.Params.ByName("id"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) GetAlerts(c *gin.Context) {
	teams, err := base.DB.GetAlerts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(teams, err, c)
}

func (base *Controller) CreateAlert(c *gin.Context) {
	m := new(model.Alert)
	err = c.ShouldBindJSON(&m)
	team, err := base.DB.CreateAlert(m)
	if err != nil {
		reposeHandler(m, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) UpdateAlert(c *gin.Context) {
	body, _ := getAlertBody(c)
	team, err := base.DB.UpdateAlert(c.Params.ByName("id"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) DeleteAlert(c *gin.Context) {
	q, err := base.DB.DeleteAlert(c.Params.ByName("id"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (base *Controller) DropAlerts(c *gin.Context) {
	team, err := base.DB.DropAlerts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}
