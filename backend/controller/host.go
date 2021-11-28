package controller

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func (base *Controller) GetHosts(c *gin.Context) {
	var m []model.Host
	if err := base.DB.Find(&m).Error; err != nil {
		logger.Errorf("GetPost error: %v", err)
		c.JSON(200, err)
	} else {
		c.JSON(200, m)
	}
	var msg TMsg
	msg.Topic = "plugins.update"
	msg.Message = "start update of plugins"
	base.publishMSG(msg)
}

func (base *Controller) GetHostDB(id string) (*model.Host, error) {
	m := new(model.Host)
	if err := base.DB.Where("id = ? ", id).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (base *Controller) GetHostByName(name string) (*model.Host, error) {
	m := new(model.Host)
	if err := base.DB.Where("name = ? ", name).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (base *Controller) GetHost(c *gin.Context) {
	id := c.Params.ByName("id")
	d, err := base.GetHostDB(id)
	if err != nil {
		reposeHandler(d, err, c)
	} else {
		reposeHandler(d, err, c)
	}
}

func (base *Controller) CreateHost(c *gin.Context) {
	m := new(model.Host)
	err = c.ShouldBindJSON(&m)
	if err := base.DB.Create(&m).Error; err != nil {
		reposeHandler(m, err, c)
	} else {
		reposeHandler(m, err, c)
	}

}

func getHostBody(ctx *gin.Context) (dto *model.Host, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) DBUpdateHost(id uint, host *model.Host) (*model.Host, error)  {
	m := new(model.Host)
	query := base.DB.Where("id = ?", id).Find(&m).Updates(host)
	if query.Error != nil {
		return nil, query.Error
	}else {
		return host,  query.Error
	}
}

func (base *Controller) UpdateHost(c *gin.Context) {
	m := new(model.Host)
	id := c.Params.ByName("id")
	body, _ := getHostBody(c)
	query := base.DB.Where("id = ?", id).Find(&m)
	query = base.DB.Model(&m).Updates(body)
	if query.Error != nil {
		reposeHandler(m, query.Error, c)
	} else {
		reposeHandler(m, nil, c)
	}
}

func (base *Controller) DeleteHost(c *gin.Context) {
	m := new(model.Host)
	id := c.Params.ByName("id")
	if err := base.DB.Where("id = ? ", id).Delete(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		reposeHandler(m, err, c)
	} else {
		reposeHandler(m, nil, c)
	}
}
