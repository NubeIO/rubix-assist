package controller

import (
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) UploadFile(c *gin.Context) {
	destination := c.Query("destination")
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	f, err := file.Open()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := edgecli.New(host.IP, host.RubixPort).UploadFile(file.Filename, destination, f)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)

}
