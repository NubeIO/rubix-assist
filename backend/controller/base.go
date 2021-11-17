package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/melbahja/goph"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
	SSH *goph.Client
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if body == nil {
			ctx.JSON(404, Message{Message: "unknown error"})
		} else {
			ctx.JSON(404, Message{Message: err.Error()})
		}
	} else {
		ctx.JSON(200, body)
	}
}
