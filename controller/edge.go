package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-assist/service/edgeapi"
	"github.com/NubeIO/rubix-assist/service/store"
	"github.com/gin-gonic/gin"
)

// UploadEdgeApp
// upload the build
func (inst *Controller) UploadEdgeApp(c *gin.Context) {
	var m *store.EdgeApp
	err = c.ShouldBindJSON(&m)
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.UploadEdgeApp(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// InstallEdgeApp
// make all the dirs and install the uploaded build
func (inst *Controller) InstallEdgeApp(c *gin.Context) {
	var m *installer.Install
	err = c.ShouldBindJSON(&m)
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Store.InstallEdgeApp(host.UUID, host.Name, m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// UploadEdgeService
// upload the service file
func (inst *Controller) UploadEdgeService(c *gin.Context) {
	//file, err := c.FormFile("file")
	//if err != nil {
	//	reposeHandler(nil, err, c)
	//	return
	//}
	//m := &installer.Upload{
	//	Name:      c.Query("name"),
	//	BuildName: c.Query("buildName"),
	//	Version:   c.Query("version"),
	//	File:      file,
	//}
	//data, err := inst.Rubix.UploadServiceFile(m)
	//if err != nil {
	//	reposeHandler(nil, err, c)
	//	return
	//}
	//reposeHandler(data, nil, c)
}

func (inst *Controller) InstallEdgeService(c *gin.Context) {
	//var m *installer.Install
	//err = c.ShouldBindJSON(&m)
	//data, err := inst.Rubix.InstallService(m)
	//if err != nil {
	//	reposeHandler(nil, err, c)
	//	return
	//}
	//reposeHandler(data, nil, c)
}

func (inst *Controller) TaskBuilder(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, resp := inst.Edge.TaskBuilder(m)
	if data == nil {
		reposeHandler(resp.Message, nil, c)
	} else {
		reposeHandler(data, nil, c)
	}

}

func (inst *Controller) TaskRunner(c *gin.Context) {
	m := &edgeapi.AppTask{}
	err = c.ShouldBindJSON(&m)
	data, err := inst.Edge.TaskRunner(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
