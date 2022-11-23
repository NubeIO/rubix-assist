package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getTeamBody(ctx *gin.Context) (dto *amodel.Team, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) TeamsSchema(ctx *gin.Context) {
}

func (inst *Controller) GetTeam(c *gin.Context) {
	team, err := inst.DB.GetTeam(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}

func (inst *Controller) GetTeams(c *gin.Context) {
	teams, err := inst.DB.GetTeams()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(teams, err, c)
}

func (inst *Controller) CreateTeam(c *gin.Context) {
	m := new(amodel.Team)
	err := c.ShouldBindJSON(&m)
	team, err := inst.DB.CreateTeam(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}

func (inst *Controller) UpdateTeam(c *gin.Context) {
	body, _ := getTeamBody(c)
	team, err := inst.DB.UpdateTeam(c.Params.ByName("uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}

func (inst *Controller) DeleteTeam(c *gin.Context) {
	q, err := inst.DB.DeleteTeam(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
	} else {
		responseHandler(q, err, c)
	}
}

func (inst *Controller) DropTeams(c *gin.Context) {
	team, err := inst.DB.DropTeams()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}
