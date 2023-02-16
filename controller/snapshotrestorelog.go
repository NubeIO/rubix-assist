package controller

import "github.com/gin-gonic/gin"

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
