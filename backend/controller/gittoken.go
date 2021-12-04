package controller

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/logger"
	"github.com/gin-gonic/gin"
)

func (base *Controller) GetTokens(c *gin.Context) {
	var m []model.Token
	if err := base.DB.DB.Find(&m).Error; err != nil {
		logger.Errorf("GetPost error: %v", err)
		c.JSON(200, err)
	} else {
		c.JSON(200, m)
	}
}

func (base *Controller) dbGetToken(id string) (*model.Token, error) {
	m := new(model.Token)
	if err := base.DB.DB.Where("id = ? ", id).First(&m).Error; err != nil {
		logger.Errorf("GetToken error: %v", err)
		return m, err
	}
	return m, err
}

func (base *Controller) GetToken(c *gin.Context) {
	id := c.Params.ByName("id")
	token, err := base.dbGetToken(id)
	if err != nil {
		reposeHandler(token, err, c)
	} else {
		reposeHandler(token, nil, c)
	}
}

func (base *Controller) CreateToken(c *gin.Context) {
	m := new(model.Token)
	err := c.ShouldBindJSON(&m)
	if err := base.DB.DB.Create(&m).Error; err != nil {
		logger.Errorf("CreateToken error: %v", err)
	}
	reposeHandler(m, err, c)
}

func getTokenBody(ctx *gin.Context) (dto *model.Token, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) UpdateToken(c *gin.Context) {
	m := new(model.Token)
	id := c.Params.ByName("id")
	body, _ := getTokenBody(c)
	query := base.DB.DB.Where("id = ?", id).Find(&m)
	query = base.DB.DB.Model(&m).Updates(body)
	if query.Error != nil {
		reposeHandler(m, query.Error, c)
	} else {
		reposeHandler(m, nil, c)
	}
}

func (base *Controller) DeleteToken(c *gin.Context) {
	m := new(model.Token)
	id := c.Params.ByName("id")
	if err := base.DB.DB.Where("id = ? ", id).Delete(&m).Error; err != nil {
		logger.Errorf("DeleteToken error: %v", err)
		reposeHandler(m, err, c)
	} else {
		reposeHandler(m, nil, c)
	}
}
