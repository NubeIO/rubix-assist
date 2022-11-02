package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/rubix-assist/helpers"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) EdgeBiosLogin(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := helpers.GetEdgeBiosClient(host)
	var m *user.User
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	data, err := cli.Login(m)
	responseHandler(data, err, c)
}

func (inst *Controller) EdgeTokens(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := helpers.GetEdgeBiosClient(host)
	cli.SetJwtTokenHeader(c.Param("jwt_token"))
	data, err := cli.Tokens()
	responseHandler(data, err, c)
}
