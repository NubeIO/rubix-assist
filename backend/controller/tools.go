package controller

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/admin"
	"github.com/NubeIO/rubix-updater/model/schema"
	"github.com/gin-gonic/gin"
)

func (base *Controller) ToolsEndPoints(ctx *gin.Context) {
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := base.hostCopy(host)
	_host := admin.Admin{
		Host: h,
	}
	arch, _, err := _host.DetectArch()
	if err != nil {
		reposeHandler(nil, errors.New("error on check host"), ctx)
		return
	}
	if arch.IsBeagleBone {
		reposeHandler(schema.GetToolsEndPointsSchema(), nil, ctx)
	} else {
		reposeHandler(schema.GetToolsEndPointsSchema(), nil, ctx)
	}

}
