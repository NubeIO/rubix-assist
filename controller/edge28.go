package controller

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/edge28"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/admin"
	"github.com/NubeIO/rubix-assist/model/schema"
	"github.com/gin-gonic/gin"
)

func getEdgeIPBody(ctx *gin.Context) (dto edge28.EdgeNetworking, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) EdgeIPSchema(ctx *gin.Context) {
	reposeHandler(schema.GetEdge28IPSchema(), nil, ctx)
}

func (base *Controller) EdgeSetIP(ctx *gin.Context) {
	body, err := getEdgeIPBody(ctx)
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
		reposeHandler(nil, errors.New("error on check if host is a edge-28"), ctx)
		return
	}
	if arch.IsBeagleBone {
		ok, _ := edge28.SetIP(body)
		if !ok {
			reposeHandler(nil, errors.New("error on trying to update the networking"), ctx)
			return
		} else {
			reposeHandler("updated networking", nil, ctx)
			return
		}
	} else {
		reposeHandler(nil, errors.New("incorrect host type found"), ctx)
		return
	}

}
