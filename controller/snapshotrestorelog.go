package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) GetSnapshotRestoreLogs(c *gin.Context) {
	restoreLogs, err := inst.DB.GetSnapshotRestoreLogs()
	responseHandler(restoreLogs, err, c)
}

func (inst *Controller) DeleteSnapshotRestoreLog(c *gin.Context) {
	q, err := inst.DB.DeleteSnapshotRestoreLog(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}
