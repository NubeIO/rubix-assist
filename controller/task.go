package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/gin-gonic/gin"
)

func getTaskBody(ctx *gin.Context) (dto *model.Task, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) TasksSchema(ctx *gin.Context) {
}

func (inst *Controller) GetTask(c *gin.Context) {
	team, err := inst.DB.GetTask(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) GetTasks(c *gin.Context) {
	teams, err := inst.DB.GetTasks()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(teams, err, c)
}

func (inst *Controller) CreateTask(c *gin.Context) {
	m := new(model.Task)
	err = c.ShouldBindJSON(&m)
	res, err := inst.DB.CreateTask(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(res, err, c)
}

func (inst *Controller) UpdateTask(c *gin.Context) {
	body, _ := getTaskBody(c)
	team, err := inst.DB.UpdateTask(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) DeleteTask(c *gin.Context) {
	q, err := inst.DB.DeleteTask(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropTasks(c *gin.Context) {
	team, err := inst.DB.DropTasks()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}
