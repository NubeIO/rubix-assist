package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getSnapshotRestoreLogBody(ctx *gin.Context) (dto *amodel.SnapshotRestoreLog, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetSnapshotRestoreLogs(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	restoreLogs, err := inst.DB.GetSnapshotRestoreLogs(host.UUID)
	responseHandler(restoreLogs, err, c)
}

func (inst *Controller) DeleteSnapshotRestoreLog(c *gin.Context) {
	q, err := inst.DB.DeleteSnapshotRestoreLog(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}

func (inst *Controller) UpdateSnapshotRestoreLog(c *gin.Context) {
	body, err := getSnapshotRestoreLogBody(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	restoreLog, err := inst.DB.UpdateSnapshotRestoreLog(c.Params.ByName("uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(restoreLog, err, c)
}
