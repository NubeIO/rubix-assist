package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/edge28/edgeip"
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
	//body, err := getEdgeIPBody(ctx)
	_, session, err := inst.getHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}

	session.EdgeSetIP(nil)

}
