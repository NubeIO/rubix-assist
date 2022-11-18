package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/ssh"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	model2 "github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"net/http"

	dbase "github.com/NubeIO/rubix-assist/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/melbahja/goph"
)

type Controller struct {
	SSH      *goph.Client
	DB       *dbase.DB
	Rest     *rest.Service
	Store    *appstore.Store
	FileMode int
}

func (inst *Controller) resolveHost(c *gin.Context) (*model2.Host, error) {
	uuid := matchHostUUID(c)
	name := matchHostName(c)
	if uuid == "" && name == "" {
		return nil, errors.New("host-uuid, and host-name can bot not be empty")
	}
	if uuid != "" {
		host, _ := inst.DB.GetHost(uuid)
		if host != nil {
			return host, nil
		}
	}
	if name != "" {
		host, _ := inst.DB.GetHostByName(name)
		if host != nil {
			return host, nil
		}
	}
	var hostNames []string
	var hostUUIDs []string
	var count int
	hosts, err := inst.DB.GetHosts()
	if err != nil {
		return nil, err
	}
	for _, h := range hosts {
		hostNames = append(hostNames, h.Name)
		hostUUIDs = append(hostUUIDs, h.UUID)
		count++
	}
	return nil, errors.New(fmt.Sprintf("no valid host was found: host count: %d, host names found: %v uuids: %v", count, hostNames, hostUUIDs))
}

func matchUUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "hos_" {
			return true
		}
	}
	return false
}

func matchHostUUID(ctx *gin.Context) string {
	hostID := resolveHeaderHostID(ctx)
	if len(hostID) == 16 {
		if matchUUID(hostID) {
			return hostID
		}
	}
	return ""
}

func matchHostName(ctx *gin.Context) string {
	name := resolveHeaderHostName(ctx)
	return name
}

func resolveHeaderHostID(ctx *gin.Context) string {
	return ctx.GetHeader("host_uuid")
}

func resolveHeaderHostName(ctx *gin.Context) string {
	return ctx.GetHeader("host_name")
}

func responseHandler(body interface{}, err error, c *gin.Context, statusCode ...int) {
	var code int
	if err != nil {
		if len(statusCode) > 0 {
			code = statusCode[0]
		} else {
			code = http.StatusNotFound
		}
		msg := model2.Message{
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

// hostCopy copy same types from this host to the host needed for ssh.Host
func (inst *Controller) hostCopy(host *model2.Host) (ssh.Host, error) {
	h := new(ssh.Host)
	err := copier.Copy(&h, &host)
	if err != nil {
		fmt.Println(err)
		return *h, err
	} else {
		return *h, err
	}
}
