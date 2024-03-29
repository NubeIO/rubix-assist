package controller

import (
	"github.com/NubeIO/lib-networking/networking"
	"github.com/gin-gonic/gin"
)

var nets = networking.New()

func (inst *Controller) Networking(c *gin.Context) {
	data, err := nets.GetNetworks()
	if err != nil {
		responseHandler(data, err, c)
		return
	}
	responseHandler(data, err, c)
}

func (inst *Controller) GetInterfacesNames(c *gin.Context) {
	data, err := nets.GetInterfacesNames()
	if err != nil {
		responseHandler(data, err, c)
		return
	}
	responseHandler(data, err, c)
}

func (inst *Controller) InternetIP(c *gin.Context) {
	data, err := nets.GetInternetIP()
	if err != nil {
		responseHandler(data, err, c)
		return
	}
	responseHandler(data, err, c)
}
