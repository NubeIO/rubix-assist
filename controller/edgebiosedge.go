package controller

import (
	"errors"
	"github.com/NubeIO/rubix-assist/helpers"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path"
)

func (inst *Controller) EdgeBiosRubixEdgeUpload(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	var m *assistmodel.FileUpload
	err = c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	err = inst.attachFileOnModel(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	cli := helpers.GetEdgeBiosClient(host)
	data, err := cli.RubixEdgeUpload(m)
	responseHandler(data, err, c)
}

func (inst *Controller) attachFileOnModel(m *assistmodel.FileUpload) error {
	storePath := inst.Store.GetAppsStoreAppWithArchVersionPath("rubix-edge", m.Arch, m.Version)
	files, err := ioutil.ReadDir(storePath)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("rubix-edge store file doesn't exist")
	}
	m.File = path.Join(storePath, files[0].Name())
	return nil
}
