package controller

import (
	"github.com/NubeIO/rubix-assist-model/model"
	"github.com/NubeIO/rubix-assist-model/model/schema"

	"github.com/gin-gonic/gin"
)

func getTokenBody(ctx *gin.Context) (dto *model.Token, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) TokenSchema(ctx *gin.Context) {
	reposeHandler(schema.GetTokenSchema(), err, ctx)
}

func (inst *Controller) GetToken(c *gin.Context) {
	Token, err := inst.DB.GetToken(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Token, err, c)
}

func (inst *Controller) GetTokens(c *gin.Context) {
	t, err := inst.DB.GetTokens()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(t, err, c)
}

func (inst *Controller) CreateToken(c *gin.Context) {
	m := new(model.Token)
	err = c.ShouldBindJSON(&m)
	t, err := inst.DB.CreateToken(m)
	if err != nil {
		reposeHandler(m, err, c)
		return
	}
	reposeHandler(t, err, c)
}

func (inst *Controller) UpdateToken(c *gin.Context) {
	body, _ := getTokenBody(c)
	t, err := inst.DB.UpdateToken(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(t, err, c)
}

func (inst *Controller) DeleteToken(c *gin.Context) {
	q, err := inst.DB.DeleteToken(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropTokens(c *gin.Context) {
	t, err := inst.DB.DropTokens()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(t, err, c)
}
