package controller

import (
	"github.com/NubeIO/edge/service/client"

	"github.com/gin-gonic/gin"
)

func (inst *Controller) UploadFile(c *gin.Context) {
	destination := c.Query("destination")
	file, err := c.FormFile("file")
	host, _, err := inst.resolveHost(c)
	if err != nil {
		reposeWithCode(404, nil, err, c)
		return
	}
	data, err := client.New(host.IP, host.WiresPort).UploadFile(file.Filename, destination)
	if err != nil {
		reposeWithCode(404, data, nil, c)
		return
	}
	reposeWithCode(404, data, nil, c)

}
