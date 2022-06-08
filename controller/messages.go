package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/gin-gonic/gin"
)

func getMessageBody(ctx *gin.Context) (dto *model.Message, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) MessagesSchema(ctx *gin.Context) {
	//reposeHandler(schema.GetMessageSchema(), nil, ctx)
}

func (inst *Controller) GetMessage(c *gin.Context) {
	team, err := inst.DB.GetMessage(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) GetMessages(c *gin.Context) {
	teams, err := inst.DB.GetMessages()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(teams, err, c)
}

func (inst *Controller) CreateMessage(c *gin.Context) {
	m := new(model.Message)
	err = c.ShouldBindJSON(&m)
	team, err := inst.DB.CreateMessage(m)
	if err != nil {
		reposeHandler(m, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) UpdateMessage(c *gin.Context) {
	body, _ := getMessageBody(c)
	team, err := inst.DB.UpdateMessage(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) DeleteMessage(c *gin.Context) {
	q, err := inst.DB.DeleteMessage(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropMessages(c *gin.Context) {
	team, err := inst.DB.DropMessages()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}
