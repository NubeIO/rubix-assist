package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/schema"
	"github.com/gin-gonic/gin"
	sh "github.com/helloyi/go-sshclient"
	"github.com/melbahja/goph"
)

type Message struct {
	Message string `json:"message"`
}

func getHostBody(ctx *gin.Context) (dto *model.Host, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) HostsSchema(ctx *gin.Context) {
	reposeHandler(schema.GetHostSchema(), nil, ctx)
}

func (base *Controller) GetHost(c *gin.Context) {
	host, err := base.DB.GetHostByName(c.Params.ByName("uuid"), true)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (base *Controller) GetHosts(c *gin.Context) {
	hosts, err := base.DB.GetHosts()

	fmt.Println("222222")
	client1, err := sh.DialWithPasswd("120.151.62.75:2221", "debian", "N00BConnect")
	if err != nil {
		fmt.Println(err)
	}
	defer client1.Close()

	ccc, _ := client1.Cmd("pwd").Output()
	fmt.Println(string(ccc))

	fmt.Println("222222")
	fmt.Println("222222")

	client2, err := goph.NewConn(&goph.Config{
		User:     "debian",
		Addr:     "120.151.62.75",
		Port:     2221,
		Auth:     goph.Password("N00BConnect"),
		Callback: VerifyHost,
	})

	//client2, err := goph.New("debian", "120.151.62.75:2221", goph.Password("N00BConnect"))
	if err != nil {
		reposeHandler(nil, err, c)
	}
	cmd, err := client2.Command("pwd")
	fmt.Println(cmd.String())

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
	host, err := base.DB.UpdateHost(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}

func (base *Controller) DeleteHost(c *gin.Context) {
	q, err := base.DB.DeleteHost(c.Params.ByName("uuid"))
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
