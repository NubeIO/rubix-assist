package controller

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"net/http"

	dbase "github.com/NubeIO/rubix-assist/database"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/remote"
	"github.com/NubeIO/rubix-assist/service/remote/ssh"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/melbahja/goph"
)

const nonRoot = 0700
const root = 0777

var fileUtils = fileutils.New()
var filePerm = root
var err error

type Controller struct {
	SSH   *goph.Client
	DB    *dbase.DB
	Rest  *rest.Service
	Store *appstore.Store
}

func (inst *Controller) resolveHost(c *gin.Context) (*model.Host, error) {
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
	return nil, errors.New(fmt.Sprintf("no valid host was found: host count:%d, host names found:%v uuids:%v", count, hostNames, hostUUIDs))
}

func (inst *Controller) getHost(c *gin.Context) (host *model.Host, session *remote.Admin, err error) {
	host, err = inst.resolveHost(c)
	rs := &remote.Admin{
		SSH: &ssh.Host{
			Host: &model.Host{
				IP:       host.IP,
				Port:     host.Port,
				Username: host.SSHUsername,
				Password: host.SSHPassword,
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
