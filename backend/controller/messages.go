package controller

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/schema"
	"github.com/gin-gonic/gin"
)

func getMessageBody(ctx *gin.Context) (dto *model.Message, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) MessagesSchema(ctx *gin.Context) {
	reposeHandler(schema.GetMessageSchema(), nil, ctx)
}

func (base *Controller) GetMessage(c *gin.Context) {
	team, err := base.DB.GetMessage(c.Params.ByName("id"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) GetMessages(c *gin.Context) {
	teams, err := base.DB.GetMessages()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(teams, err, c)
}

func (base *Controller) CreateMessage(c *gin.Context) {
	m := new(model.Message)
	err = c.ShouldBindJSON(&m)
	team, err := base.DB.CreateMessage(m)
	if err != nil {
		reposeHandler(m, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) UpdateMessage(c *gin.Context) {
	body, _ := getMessageBody(c)
	team, err := base.DB.UpdateMessage(c.Params.ByName("id"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (base *Controller) DeleteMessage(c *gin.Context) {
	q, err := base.DB.DeleteMessage(c.Params.ByName("id"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (base *Controller) DropMessages(c *gin.Context) {
	team, err := base.DB.DropMessages()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}
