package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TUpdatePlugins struct {
	RemoveOldPlugins string `json:"remove_old_plugins_pass"`
	MakeUploadDir    string `json:"make_upload_dir_pass"`
	UpLoadPlugins    string `json:"upload_plugins_pass"`
	CleanUp          string `json:"clean_up_pass"`
	RestartFF        string `json:"restart_flow_framework"`
}

type TMsg struct {
	Topic   string
	Message string
	IsError bool
}



//UpdatePlugins full install of the plugins as in upload, unzip and restart flow framework
func (base *Controller) UpdatePlugins(ctx *gin.Context) {
	var msg TMsg
	msg.Topic = "plugins.update"
	msg.Message = "start update of plugins"
	base.publishMSG(msg)
	body := uploadBody(ctx)
	id := ctx.Params.ByName("id")
	result := new(TUpdatePlugins)
	result.RemoveOldPlugins = "PASS: removed old plugins"
	result.MakeUploadDir = "PASS: made tmp upload dir"
	result.UpLoadPlugins = "PASS: upload and unzip plugins"
	result.CleanUp = "PASS: deleted tmp upload dir"
	result.RestartFF = "PASS: deleted tmp upload dir"
	_, err := base.clearDir(id, "/data/flow-framework/data/plugins")
	if err != nil {
		result.RemoveOldPlugins = "PASS or FAIL: failed to remove existing OR there where no existing plugins installed"
		msg.Message = result.RemoveOldPlugins
		base.publishMSG(msg)
	} else {
		msg.Message = result.RemoveOldPlugins
		base.publishMSG(msg)
	}
	//_, err = base.mkDir(id, body.ToPath, false)
	if err != nil {
		result.MakeUploadDir = "FAIL: failed to make OR new dir or was exiting"
		msg.Message = result.MakeUploadDir
		base.publishMSG(msg)
	} else {
		msg.Message = result.MakeUploadDir
		base.publishMSG(msg)
	}
	err = base.uploadZip(id, body)
	if err != nil {
		result.UpLoadPlugins = fmt.Sprint(err)
		msg.Message = result.UpLoadPlugins
		base.publishMSG(msg)
	} else {
		msg.Message = result.UpLoadPlugins
		base.publishMSG(msg)
	}
	_, err = base.rmDir(id, body.ToPath, false)
	if err != nil {
		result.CleanUp = fmt.Sprint(err)
		msg.Message = result.CleanUp
		base.publishMSG(msg)
	} else {
		msg.Message = result.CleanUp
		base.publishMSG(msg)
	}
	//_, err = base.runCommand(id, "sudo systemctl restart nubeio-flow-framework.service", true)
	//if err != nil {
	//	msg.Message = result.RestartFF
	//	base.publishMSG(msg)
	//} else {
	//	msg.Message = result.RestartFF
	//	base.publishMSG(msg)
	//}
	reposeHandler(result, err, ctx)
}

func (base *Controller) UploadPlugins(ctx *gin.Context) {
	body := uploadBody(ctx)
	id := ctx.Params.ByName("id")
	err := base.uploadZip(id, body)
	if err != nil {
		//return
	}
	reposeHandler("string(out)", err, ctx)
}

func (base *Controller) DeleteAllPlugins(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	dir, err := base.clearDir(id, "/data/flow-framework/data/plugins")
	if err != nil {
		reposeHandler(dir, err, ctx)
	} else {
		reposeHandler(dir, err, ctx)
	}
}
