package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func getHostBody(ctx *gin.Context) (dto *model.Host, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetHostSchema(ctx *gin.Context) {
	mod := model.GetHostSchema()
	res := map[string]interface{}{
		"properties": mod,
	}
	reposeHandler(res, nil, ctx)
}

func (inst *Controller) GetHost(c *gin.Context) {
	host, err := inst.DB.GetHostByName(c.Params.ByName("uuid"), true)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	host.Password = "***"
	host.RubixPassword = "***"
	reposeHandler(host, err, c)
}

func (inst *Controller) GetHosts(c *gin.Context) {
	hosts, err := inst.DB.GetHosts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	for _, host := range hosts {
		host.Password = "***"
		host.RubixPassword = "***"
	}
	reposeHandler(hosts, err, c)
}

func (inst *Controller) CreateHost(c *gin.Context) {
	m := new(model.Host)
	err = c.ShouldBindJSON(&m)
	host, err := inst.DB.CreateHost(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	host.Password = "***"
	host.RubixPassword = "***"
	reposeHandler(host, err, c)
}

func (inst *Controller) UpdateHost(c *gin.Context) {
	body, _ := getHostBody(c)
	host, _, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}

	host, err = inst.DB.UpdateHost(host.UUID, body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	host.Password = "***"
	host.RubixPassword = "***"
	reposeHandler(host, err, c)
}

func (inst *Controller) DeleteHost(c *gin.Context) {
	q, err := inst.DB.DeleteHost(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropHosts(c *gin.Context) {
	host, err := inst.DB.DropHosts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}
