package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type Snapshots struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}

func (inst *Controller) GetSnapshots(c *gin.Context) {
	snapshots, err := inst.getSnapshots()
	responseHandler(snapshots, err, c)
}

func (inst *Controller) getSnapshots() ([]Snapshots, error) {
	_path := config.Config.GetAbsSnapShotDir()
	fileInfo, err := os.Stat(_path)
	dirContent := make([]Snapshots, 0)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(_path)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			dirContent = append(dirContent, Snapshots{
				Name:      file.Name(),
				Size:      file.Size(),
				CreatedAt: file.ModTime(),
			})
		}
	} else {
		return nil, errors.New("it needs to be a directory, found a file")
	}
	return dirContent, nil
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
		Status: amodel.Creating, CreatedAt: time.Now()})
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
	responseHandler(amodel.Message{Message: "create snapshot process has submitted"}, nil, c)
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
		Msg: "", Status: amodel.Restoring, CreatedAt: time.Now()})
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
	responseHandler(amodel.Message{Message: "restore snapshot process has submitted"}, nil, c)
}
