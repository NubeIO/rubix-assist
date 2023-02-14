package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/cligetter"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var createSnapshots = make([]*amodel.CreateSnapshotStatus, 0)
var restoreSnapshots = make([]*amodel.RestoreSnapshotStatus, 0)

type Snapshots struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	CreatedAt string `json:"create_at"`
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
				CreatedAt: file.ModTime().UTC().Format(time.RFC3339),
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
	uuid_ := uuid.ShortUUID()
	createSnapshots = append(createSnapshots, &amodel.CreateSnapshotStatus{UUID: uuid_, HostUUID: host.UUID,
		Status: amodel.Creating})
	go func() {
		cli := cligetter.GetEdgeClient(host)
		snapshot, filename, err := cli.CreateSnapshot()
		if err != nil {
			deleteCreateSnapshots(uuid_)
			log.Errorf("err: %s", err.Error())
			return
		}
		err = os.WriteFile(path.Join(config.Config.GetAbsSnapShotDir(), filename), snapshot, os.FileMode(inst.FileMode))
		if err != nil {
			log.Errorf("err: %s", err.Error())
		}
		deleteCreateSnapshots(uuid_)
	}()
	responseHandler(amodel.Message{Message: "create snapshot process is submitted"}, nil, c)
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
	uuid_ := uuid.ShortUUID()
	restoreSnapshots = append(restoreSnapshots, &amodel.RestoreSnapshotStatus{UUID: uuid_, HostUUID: host.UUID,
		Status: amodel.Restoring})
	go func() {
		cli := cligetter.GetEdgeClient(host)
		reader, err := os.Open(path.Join(config.Config.GetAbsSnapShotDir(), file))
		if err != nil {
			deleteRestoreSnapshots(uuid_)
			log.Errorf("err: %s", err.Error())
			return
		}
		err = cli.RestoreSnapshot(file, reader, useGlobalUUID)
		if err != nil {
			log.Errorf("err: %s", err.Error())
		}
		deleteRestoreSnapshots(uuid_)
	}()
	responseHandler(amodel.Message{Message: "restore snapshot process is submitted"}, nil, c)
}

func (inst *Controller) GetSnapshotsStatus(c *gin.Context) {
	responseHandler(amodel.SnapshotStatus{CreateStatus: createSnapshots, RestoreStatus: restoreSnapshots}, nil, c)
}

func deleteCreateSnapshots(uuid string) {
	for i, createSnapshot := range createSnapshots {
		if createSnapshot.UUID == uuid {
			createSnapshots = append(createSnapshots[:i], createSnapshots[i+1:]...)
		}
	}
}

func deleteRestoreSnapshots(uuid string) {
	for i, restoreSnapshot := range restoreSnapshots {
		if restoreSnapshot.UUID == uuid {
			restoreSnapshots = append(restoreSnapshots[:i], restoreSnapshots[i+1:]...)
		}
	}
}
