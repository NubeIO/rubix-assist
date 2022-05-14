package controller

import (
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/model/schema"
	"github.com/gin-gonic/gin"
)

func getTeamBody(ctx *gin.Context) (dto *model.Team, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) TeamsSchema(ctx *gin.Context) {
	reposeHandler(schema.GetTeamSchema(), nil, ctx)
}

func (base *Controller) GetTeam(c *gin.Context) {
	team, err := base.DB.GetTeam(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) GetTeams(c *gin.Context) {
	teams, err := base.DB.GetTeams()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(teams, err, c)
}

func (base *Controller) CreateTeam(c *gin.Context) {
	m := new(model.Team)
	err = c.ShouldBindJSON(&m)
	team, err := base.DB.CreateTeam(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) UpdateTeam(c *gin.Context) {
	body, _ := getTeamBody(c)
	team, err := base.DB.UpdateTeam(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) DeleteTeam(c *gin.Context) {
	q, err := base.DB.DeleteTeam(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (base *Controller) DropTeams(c *gin.Context) {
	team, err := base.DB.DropTeams()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}
