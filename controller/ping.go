package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) SystemPing(c *gin.Context) {
	responseHandler(Message{Message: "boo-ya"}, err, c)
}
