package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/ufw"
	"github.com/gin-gonic/gin"
)

func (base *Controller) InstallUFW(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := hostCopy(host)
	u := ufw.UFW{
		Host: h,
	}
	out, err := u.UWFInstall()
	if err != nil {
		reposeHandler(out, err, ctx)
		return
	}
	reposeHandler(out, err, ctx)
}

func (base *Controller) InstallUFW2(ctx *gin.Context) {
	host, useID, err := base.resolveHost(ctx)
	debug := false
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	hostName := host.Name
	if useID {
		hostName = host.ID
	}
	opts := commandOpts{
		id:    hostName,
		cmd:   "sudo ufw ",
		debug: debug,
		host:  *host,
	}
	_, install, err := base.runCommand(opts, host.IsLocalhost)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}

	fmt.Println("install", install)
	reposeHandler("install completed", err, ctx)
}
