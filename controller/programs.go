package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/admin"
	"github.com/gin-gonic/gin"
)

func (base *Controller) NodeJsVersion(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := base.hostCopy(host)
	u := admin.Admin{
		Host: h,
	}
	out, _, err := u.NodeGetVersion()
	if err != nil {
		reposeHandler(out, err, ctx)
		return
	}
	reposeHandler(out, err, ctx)
}

func (base *Controller) NodeJsInstall(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := base.hostCopy(host)
	u := admin.Admin{
		Host: h,
	}
	out, err := u.InstallNode14()
	if err != nil {
		reposeHandler(out, err, ctx)
		return
	}
	reposeHandler(out, err, ctx)
}
