package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) GetSnapshotCreateLogs(c *gin.Context) {
	createLogs, err := inst.DB.GetSnapshotCreateLogs()
	responseHandler(createLogs, err, c)
}

func (inst *Controller) DeleteSnapshotCreateLog(c *gin.Context) {
	q, err := inst.DB.DeleteSnapshotCreateLog(c.Params.ByName("uuid"))
	responseHandler(q, err, c)
}
