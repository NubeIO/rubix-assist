package controller

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/edge28/edgeip"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/remote/v1/remote"
	"github.com/NubeIO/rubix-assist/model/schema"
	"github.com/gin-gonic/gin"
)

func getEdgeIPBody(ctx *gin.Context) (dto edgeip.EdgeNetworking, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) EdgeIPSchema(ctx *gin.Context) {
	reposeHandler(schema.GetEdge28IPSchema(), nil, ctx)
}

func (inst *Controller) EdgeSetIP(ctx *gin.Context) {
	body, err := getEdgeIPBody(ctx)
	host, _, err := inst.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	h, err := inst.hostCopy(host)
	_host := remote.Admin{
		Host: h,
	}
	arch, _, err := _host.DetectArch()
	if err != nil {
		reposeHandler(nil, errors.New("error on check if host is a edge-28"), ctx)
		return
	}
	if arch.IsBeagleBone {
		ok, _ := edgeip.SetIP(body)
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