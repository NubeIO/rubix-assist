package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) EdgeDirExists(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeDirExists(host.UUID, host.Name, path)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}

func (inst *Controller) EdgeCreateDir(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	path := c.Query("path")
	data, err := inst.Store.EdgeCreateDir(host.UUID, host.Name, path)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(data, nil, c)
}
