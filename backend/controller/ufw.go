package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (base *Controller) InstallUFW(ctx *gin.Context) {
	_uwf, _ := base.UWF.UFWLoadProfile(true)

	fmt.Println(_uwf.PortsCurrentState)

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
		cmd:   "sudo apt-get install ufw",
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
