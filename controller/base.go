package controller

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/rubix-assist/service/edge"

	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"

	"github.com/NubeIO/rubix-assist/service/remote"
	"github.com/NubeIO/rubix-assist/service/remote/ssh"
	log "github.com/sirupsen/logrus"

	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/melbahja/goph"
	"gopkg.in/olahol/melody.v1"
)

type Controller struct {
	//DB  *gorm.DB
	SSH  *goph.Client
	WS   *melody.Melody //web socket
	DB   *dbase.DB
	Rest *rest.Service
	Edge *edge.Manager
}

type WsMsg struct {
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
	IsError bool        `json:"is_error"`
}

var err error

////publishMSG send websocket message
func (inst *Controller) publishMSG(in *WsMsg) ([]byte, error) {
	msg := map[string]interface{}{
		"topic":    in.Topic,
		"msg":      in.Message,
		"is_error": in.IsError,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		//panic(err)
	}
	if in.IsError {
		log.Errorf("ERROR: publish websocket topic: %v\n", in.Topic)
	} else {
		log.Infof("INFO: publish websocket topic: %v\n", in.Topic)
	}
	err = inst.WS.Broadcast(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

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

func resolveHeaderGitToken(ctx *gin.Context) string {
	return ctx.GetHeader("git_token")
}

func reposeWithCode(code int, body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(code, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(code, body)
			} else {
				ctx.JSON(code, Message{Message: err.Error()})
			}

		}
	} else {
		ctx.JSON(code, body)
	}
}

type Response struct {
	StatusCode   int         `json:"status_code"`
	ErrorMessage string      `json:"error_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(404, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(404, body)
			} else {
				ctx.JSON(404, Message{Message: err.Error()})
			}
		}
	} else {
		ctx.JSON(200, body)
	}
}

func reposeMessage(code int, body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(code, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(code, body)
			} else {
				ctx.JSON(code, Message{Message: err.Error()})
			}

		}
	} else {
		ctx.JSON(code, body)
	}
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
