package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/rubix-updater/model/rubix"
	"github.com/gin-gonic/gin"
	"time"
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

func (base *Controller) PluginFullInstall(ctx *gin.Context) {
	body, err := bodyAppsDownload(ctx)
	po := proxyOptions{
		ctx:          ctx,
		refreshToken: true,
		NonProxyReq:  true,
	}
	proxyReq, opt, rtn, err := base.buildReq(po)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	opt = &nrest.ReqOpt{
		Timeout:          500 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    0 * time.Second,
		RetryMaxWaitTime: 0,
		Headers:          map[string]interface{}{"Authorization": rtn.Token},
		Json:             body,
	}

	downloadCount := 0
	//get state
	getState := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
	fmt.Println(getState.StatusCode)
	fmt.Println(getState.AsString())
	//delete state
	deleteState := proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
	fmt.Println(deleteState.StatusCode)
	fmt.Println(deleteState.AsString())

	appDownload := proxyReq.Do(nrest.POST, AppsUrls.Download, opt)
	fmt.Println(appDownload.Err)
	fmt.Println(appDownload.StatusCode)
	fmt.Println(appDownload.AsString())

	//
	for {
		req := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
		state := new(rubix.AppsDownloadState)
		req.ToInterface(&state)
		fmt.Println(req.Err)
		fmt.Println(req.StatusCode)
		fmt.Println(4444, state.State)
		time.Sleep(4 * time.Second)
		downloadCount++
		fmt.Println("downloaded")
		if state.State == "DOWNLOADED" {
			break
		}
	}
	appInstall := proxyReq.Do(nrest.POST, AppsUrls.Install, opt)
	fmt.Println(appInstall.Err)
	fmt.Println(appInstall.StatusCode)
	fmt.Println(appInstall.AsString())

	deleteState = proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
	fmt.Println(deleteState.StatusCode)
	fmt.Println(deleteState.AsString())

}

//UpdatePlugins full install of the plugins as in upload, unzip and restart flow framework
func (base *Controller) UpdatePlugins(ctx *gin.Context) {
	var msg TMsg
	msg.Topic = "plugins.update"
	msg.Message = "start update of plugins"
	base.publishMSG(msg)
	body := uploadBody(ctx)
	uuid := ctx.Params.ByName("uuid")
	result := new(TUpdatePlugins)
	result.RemoveOldPlugins = "PASS: removed old plugins"
	result.MakeUploadDir = "PASS: made tmp upload dir"
	result.UpLoadPlugins = "PASS: upload and unzip plugins"
	result.CleanUp = "PASS: deleted tmp upload dir"
	result.RestartFF = "PASS: deleted tmp upload dir"
	_, err := base.clearDir(uuid, "/data/flow-framework/data/plugins")
	if err != nil {
		result.RemoveOldPlugins = "PASS or FAIL: failed to remove existing OR there where no existing plugins installed"
		msg.Message = result.RemoveOldPlugins
		base.publishMSG(msg)
	} else {
		msg.Message = result.RemoveOldPlugins
		base.publishMSG(msg)
	}
	//_, err = base.mkDir(uuid, body.ToPath, false)
	if err != nil {
		result.MakeUploadDir = "FAIL: failed to make OR new dir or was exiting"
		msg.Message = result.MakeUploadDir
		base.publishMSG(msg)
	} else {
		msg.Message = result.MakeUploadDir
		base.publishMSG(msg)
	}
	err = base.uploadZip(uuid, body)
	if err != nil {
		result.UpLoadPlugins = fmt.Sprint(err)
		msg.Message = result.UpLoadPlugins
		base.publishMSG(msg)
	} else {
		msg.Message = result.UpLoadPlugins
		base.publishMSG(msg)
	}
	_, err = base.rmDir(uuid, body.ToPath, false)
	if err != nil {
		result.CleanUp = fmt.Sprint(err)
		msg.Message = result.CleanUp
		base.publishMSG(msg)
	} else {
		msg.Message = result.CleanUp
		base.publishMSG(msg)
	}
	//_, err = base.runCommand(uuid, "sudo systemctl restart nubeio-flow-framework.service", true)
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
	uuid := ctx.Params.ByName("uuid")
	err := base.uploadZip(uuid, body)
	if err != nil {
		//return
	}
	reposeHandler("string(out)", err, ctx)
}

func (base *Controller) DeleteAllPlugins(ctx *gin.Context) {
	uuid := ctx.Params.ByName("uuid")
	dir, err := base.clearDir(uuid, "/data/flow-framework/data/plugins")
	if err != nil {
		reposeHandler(dir, err, ctx)
	} else {
		reposeHandler(dir, err, ctx)
	}
}
