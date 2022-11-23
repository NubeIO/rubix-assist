package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/helpers"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func (inst *Controller) ListAppsWithVersions(c *gin.Context) {
	data, err := inst.Store.ListAppsWithVersions()
	responseHandler(data, err, c)
}

func (inst *Controller) UploadAddOnAppStore(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	m := &amodel.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		Arch:    c.Query("arch"),
		File:    file,
	}
	data, err := inst.Store.UploadAddOnAppStore(m)
	responseHandler(data, err, c)
}

func (inst *Controller) CheckAppExistence(c *gin.Context) {
	name := c.Query("name")
	arch := c.Query("arch")
	version := c.Query("version")
	if err := inst.checkAppExistence(name, arch, version); err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(amodel.FoundMessage{Found: true}, nil, c)
}

func (inst *Controller) checkAppExistence(name, arch, version string) error {
	if name == "" {
		return errors.New("name can not be empty")
	}
	if err := helpers.CheckVersion(version); err != nil {
		return err
	}
	if arch == "" {
		return errors.New("arch can not be empty")
	}
	p := inst.Store.GetAppsStoreAppWithArchVersionPath(name, arch, version)
	found := fileutils.DirExists(p)
	if !found {
		return errors.New(fmt.Sprintf("failed to find app: %s with arch: %s & version: %s with  in app store", name, arch, version))
	}
	files, _ := ioutil.ReadDir(p)
	if len(files) == 0 {
		return errors.New(fmt.Sprintf("failed to find app: %s with arch: %s & version: %s with  in app store", name, arch, version))
	}
	return nil
}
