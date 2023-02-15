package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ttime"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

func (inst *Controller) GetSnapshots(c *gin.Context) {
	snapshots, err := inst.listFiles(config.Config.GetAbsSnapShotDir())
	responseHandler(snapshots, err, c)
}

func (inst *Controller) DeleteSnapshot(c *gin.Context) {
	file := c.Query("file")
	if file == "" {
		responseHandler(nil, errors.New("file can not be empty"), c)
		return
	}
	err := os.Remove(path.Join(config.Config.GetAbsSnapShotDir(), file))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(amodel.Message{Message: fmt.Sprintf("deleted file: %s", file)}, err, c)
}

func (inst *Controller) CreateSnapshot(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	createLog, err := inst.DB.CreateSnapshotCreateLog(&amodel.SnapshotCreateLog{UUID: "", HostUUID: host.UUID, Msg: "",
		Status: amodel.Creating, CreatedAt: ttime.Now()})
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	go func() {
		cli := cligetter.GetEdgeClient(host)
		snapshot, filename, err := cli.CreateSnapshot()
		if err == nil {
			err = os.WriteFile(path.Join(config.Config.GetAbsSnapShotDir(), filename), snapshot,
				os.FileMode(inst.FileMode))
		}
		createLog.Status = amodel.Created
		createLog.Msg = filename
		if err != nil {
			createLog.Status = amodel.CreateFailed
			createLog.Msg = err.Error()
		}
		_, _ = inst.DB.UpdateSnapshotCreateLog(createLog.UUID, createLog)
	}()
	responseHandler(amodel.Message{Message: "create snapshot started"}, nil, c)
}

func (inst *Controller) RestoreSnapshot(c *gin.Context) {
	file := c.Query("file")
	if file == "" {
		responseHandler(nil, errors.New("file can not be empty"), c)
		return
	}
	useGlobalUUID := c.Query("use_global_uuid")
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	restoreLog, err := inst.DB.CreateSnapshotRestoreLog(&amodel.SnapshotRestoreLog{UUID: "", HostUUID: host.UUID,
		Msg: "", Status: amodel.Restoring, CreatedAt: ttime.Now()})
	go func() {
		cli := cligetter.GetEdgeClient(host)
		reader, err := os.Open(path.Join(config.Config.GetAbsSnapShotDir(), file))
		if err == nil {
			err = cli.RestoreSnapshot(file, reader, useGlobalUUID)
		}
		restoreLog.Status = amodel.Restored
		restoreLog.Msg = file
		if err != nil {
			restoreLog.Status = amodel.RestoreFailed
			restoreLog.Msg = err.Error()
		}
		_, _ = inst.DB.UpdateSnapshotRestoreLog(restoreLog.UUID, restoreLog)
	}()
	responseHandler(amodel.Message{Message: "restore snapshot started"}, nil, c)
}
