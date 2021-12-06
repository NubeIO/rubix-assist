package controller

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/schema"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func getHostBody(ctx *gin.Context) (dto *model.Host, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) HostsSchema(ctx *gin.Context) {
	reposeHandler(schema.GetHostSchema(), err, ctx)
}

func (base *Controller) GetHost(c *gin.Context) {
	host, err := base.DB.GetHostByName(c.Params.ByName("id"), true)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (base *Controller) GetHosts(c *gin.Context) {
	hosts, err := base.DB.GetHosts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(hosts, err, c)
}

func (base *Controller) CreateHost(c *gin.Context) {
	m := new(model.Host)
	err = c.ShouldBindJSON(&m)
	host, err := base.DB.CreateHost(m)
	if err != nil {
		reposeHandler(m, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (base *Controller) UpdateHost(c *gin.Context) {
	body, _ := getHostBody(c)
	host, err := base.DB.UpdateHost(c.Params.ByName("id"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (base *Controller) DeleteHost(c *gin.Context) {
	q, err := base.DB.DeleteHost(c.Params.ByName("id"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (base *Controller) DropHosts(c *gin.Context) {
	host, err := base.DB.DropHosts()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}
