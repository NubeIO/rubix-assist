package controller

import (
	"github.com/NubeIO/rubix-assist/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) SystemPing(c *gin.Context) {
	responseHandler(model.Message{Message: "pong"}, nil, c)
}
