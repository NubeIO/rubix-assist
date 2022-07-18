package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) ListStore(c *gin.Context) {
	data, err := inst.Store.ListApps()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
