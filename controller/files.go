package controller

import (
	"github.com/NubeIO/rubix-assist/service/edgecli"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (inst *Controller) UploadFile(c *gin.Context) {
	destination := c.Query("destination")
	file, err := c.FormFile("file")
	if err != nil {
		reposeWithCode(http.StatusBadRequest, nil, err, c)
		return
	}
	f, err := file.Open()
	if err != nil {
		reposeWithCode(http.StatusBadRequest, nil, err, c)
		return
	}
	host, _, err := inst.resolveHost(c)
	if err != nil {
		reposeWithCode(http.StatusBadRequest, nil, err, c)
		return
	}
	data, err := edgecli.New(host.IP, host.RubixPort).UploadFile(file.Filename, destination, f)
	if err != nil {
		reposeWithCode(http.StatusBadRequest, data, nil, c)
		return
	}
	reposeWithCode(http.StatusOK, data, nil, c)

}
