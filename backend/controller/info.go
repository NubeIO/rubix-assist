package controller

import (
	"github.com/NubeIO/rubix-updater/model/schema"
	"github.com/gin-gonic/gin"
)

func (base *Controller) DeviceInfoSchema(ctx *gin.Context) {
	reposeHandler(schema.GetHostSchema(), nil, ctx)
}
