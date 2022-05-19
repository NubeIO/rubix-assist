package controller

import (
	"github.com/NubeIO/rubix-assist/model/schema"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) DeviceInfoSchema(ctx *gin.Context) {
	reposeHandler(schema.GetHostSchema(), nil, ctx)
}
