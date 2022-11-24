package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) SystemPing(c *gin.Context) {
	responseHandler(amodel.Message{Message: "pong"}, nil, c)
}
