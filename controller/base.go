package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/edgeapi"
	"net/http"

	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"

	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/NubeIO/rubix-assist/service/remote"
	"github.com/NubeIO/rubix-assist/service/remote/ssh"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/melbahja/goph"
)

type Controller struct {
	SSH  *goph.Client
	DB   *dbase.DB
	Rest *rest.Service
	Edge *edgeapi.Manager
}

var err error

func (inst *Controller) resolveHost(ctx *gin.Context) (host *model.Host, useID bool, err error) {
	idName, useID := useHostNameOrID(ctx)
	host, err = inst.DB.GetHostByName(idName, useID)
	return host, useID, err
}

func (inst *Controller) getHost(ctx *gin.Context) (host *model.Host, session *remote.Admin, err error) {
	idName, useID := useHostNameOrID(ctx)
	host, err = inst.DB.GetHostByName(idName, useID)

	rs := &remote.Admin{
		SSH: &ssh.Host{
			Host: &model.Host{
				IP:       host.IP,
				Port:     host.Port,
				Username: host.Username,
				Password: host.Password,
			},
		},
	}
	session = remote.New(rs)
	return host, session, err
}

func bodyAsJSON(ctx *gin.Context) (interface{}, error) {
	var body interface{} //get the body and put it into an interface
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func useHostNameOrID(ctx *gin.Context) (idName string, useID bool) {
	hostID := resolveHeaderHostID(ctx)
	hostName := resolveHeaderHostName(ctx)
	if hostID != "" {
		return hostID, true
	} else if hostName != "" {
		return hostName, false
	} else {
		return "", false
	}
}

func resolveHeaderHostID(ctx *gin.Context) string {
	return ctx.GetHeader("host_uuid")
}

func resolveHeaderHostName(ctx *gin.Context) string {
	return ctx.GetHeader("host_name")
}

func reposeHandler(body interface{}, err error, c *gin.Context, statusCode ...int) {
	var code int
	if err != nil {
		if len(statusCode) > 0 {
			code = statusCode[0]
		} else {
			code = http.StatusNotFound
		}
		msg := Message{
			Message: err.Error(),
		}
		c.JSON(code, msg)
	} else {
		if len(statusCode) > 0 {
			code = statusCode[0]
		} else {
			code = http.StatusOK
		}
		c.JSON(code, body)

	}
}

type Message struct {
	Message interface{} `json:"message"`
}

//hostCopy copy same types from this host to the host needed for ssh.Host
func (inst *Controller) hostCopy(host *model.Host) (ssh.Host, error) {
	h := new(ssh.Host)
	err = copier.Copy(&h, &host)
	if err != nil {
		fmt.Println(err)
		return *h, err
	} else {
		return *h, err
	}
}
