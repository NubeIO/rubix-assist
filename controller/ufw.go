package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/ufw"
	"github.com/gin-gonic/gin"
)

func (base *Controller) UFWInstall(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := base.hostCopy(host)
	u := ufw.UFW{
		Host: h,
	}
	out, err := u.UWFInstall()
	if err != nil {
		reposeHandler(out, err, ctx)
		return
	}
	reposeHandler(out, err, ctx)
	return
}

func (base *Controller) UFWEnable(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := base.hostCopy(host)
	u := ufw.UFW{
		Host: h,
	}
	out, err := u.UWFEnable()
	if err != nil {
		reposeHandler(out, err, ctx)
		return
	}
	reposeHandler(out, err, ctx)
	return
}

func (base *Controller) UFWDisable(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := base.hostCopy(host)
	u := ufw.UFW{
		Host: h,
	}
	out, err := u.UWFDisable()
	if err != nil {
		reposeHandler(out, err, ctx)
		return
	}
	reposeHandler(out, err, ctx)
	return
}

func (base *Controller) UFWAddPort(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := base.hostCopy(host)
	u := ufw.UFW{
		Host: h,
	}
	out, err := u.UWFDefaultPorts()
	if err != nil {
		reposeHandler(out, err, ctx)
		return
	}
	reposeHandler(out, err, ctx)
	return
}
