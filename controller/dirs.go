package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
	"os"
)

func (inst *Controller) DirExists(c *gin.Context) {
	path := c.Query("path")
	exists := fileutils.DirExists(path)
	dirExistence := amodel.DirExistence{Path: path, Exists: exists}
	responseHandler(dirExistence, nil, c)
}

func (inst *Controller) CreateDir(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		responseHandler(nil, errors.New("path can not be empty"), c)
		return
	}
	err := os.MkdirAll(path, os.FileMode(inst.FileMode))
	responseHandler(amodel.Message{Message: fmt.Sprintf("created directory: %s", path)}, err, c)
}
