package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

type AlertStatus struct {
	Status string `json:"status"`
}

func getAlertBody(ctx *gin.Context) (dto *amodel.Alert, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func getAlertStatus(ctx *gin.Context) (status string, err error) {
	statusStruct := AlertStatus{}
	err = ctx.ShouldBindJSON(&statusStruct)
	return statusStruct.Status, err
}

func (inst *Controller) AlertsSchema(ctx *gin.Context) {
}

func (inst *Controller) GetAlert(c *gin.Context) {
	resp, err := inst.DB.GetAlert(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(resp, err, c)
}

func (inst *Controller) GetAlerts(c *gin.Context) {
	resp, err := inst.DB.GetAlerts()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(resp, err, c)
}

func (inst *Controller) GetAlertsByHost(c *gin.Context) {
	resp, err := inst.DB.GetAlertsByHost(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(resp, err, c)
}

func (inst *Controller) CreateAlert(c *gin.Context) {
	m := new(amodel.Alert)
	err := c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	res, err := inst.DB.CreateAlert(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(res, err, c)
}

func (inst *Controller) UpdateAlertStatus(c *gin.Context) {
	status, _ := getAlertStatus(c)
	resp, err := inst.DB.UpdateAlertStatus(c.Params.ByName("uuid"), status)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(resp, err, c)
}

func (inst *Controller) DeleteAlert(c *gin.Context) {
	q, err := inst.DB.DeleteAlert(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
	} else {
		responseHandler(q, err, c)
	}
}

func (inst *Controller) DropAlerts(c *gin.Context) {
	resp, err := inst.DB.DropAlerts()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(resp, err, c)
}
