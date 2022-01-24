package controller

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/schema"

	"github.com/gin-gonic/gin"
)

func getTokenBody(ctx *gin.Context) (dto *model.Token, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) TokenSchema(ctx *gin.Context) {
	reposeHandler(schema.GetTokenSchema(), err, ctx)
}

func (base *Controller) GetToken(c *gin.Context) {
	Token, err := base.DB.GetToken(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Token, err, c)
}

func (base *Controller) GetTokens(c *gin.Context) {
	t, err := base.DB.GetTokens()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(t, err, c)
}

func (base *Controller) CreateToken(c *gin.Context) {
	m := new(model.Token)
	err = c.ShouldBindJSON(&m)
	t, err := base.DB.CreateToken(m)
	if err != nil {
		reposeHandler(m, err, c)
		return
	}
	reposeHandler(t, err, c)
}

func (base *Controller) UpdateToken(c *gin.Context) {
	body, _ := getTokenBody(c)
	t, err := base.DB.UpdateToken(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(t, err, c)
}

func (base *Controller) DeleteToken(c *gin.Context) {
	q, err := base.DB.DeleteToken(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (base *Controller) DropTokens(c *gin.Context) {
	t, err := base.DB.DropTokens()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(t, err, c)
}
