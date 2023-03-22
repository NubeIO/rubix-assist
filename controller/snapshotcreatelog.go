package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getSnapshotCreateLogBody(ctx *gin.Context) (dto *amodel.SnapshotCreateLog, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetSnapshotCreateLogs(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	createLogs, err := inst.DB.GetSnapshotCreateLogs(host.UUID)
	responseHandler(createLogs, err, c)
}

func (inst *Controller) DeleteSnapshotCreateLog(c *gin.Context) {
	q, err := inst.DB.DeleteSnapshotCreateLog(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}

func (inst *Controller) UpdateSnapshotCreateLog(c *gin.Context) {
	body, err := getSnapshotCreateLogBody(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	createLog, err := inst.DB.UpdateSnapshotCreateLog(c.Params.ByName("uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(createLog, err, c)
}
