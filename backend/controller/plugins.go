package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/rubix-updater/model/rubix"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

var PluginsUrls = struct {
	Install    string
	Installed  string
	GetPlugins string
	UnInstall  string
	Download   string
	State      string
	Apps       string
}{
	Install:    "/api/app/plugins/FLOW_FRAMEWORK/install",
	Installed:  "/api/app/plugins/FLOW_FRAMEWORK",
	GetPlugins: "/ff/api/plugins",
	UnInstall:  "/api/app/plugins/FLOW_FRAMEWORK/uninstall",
	Download:   "/api/app/plugins/FLOW_FRAMEWORK/download",
	State:      "/api/app/plugins/FLOW_FRAMEWORK/download_state",
	Apps:       "/api/app/control",
}

type blankSlice map[string]interface{}

type pluginsList []Plugins

type Plugins struct {
	Author       string      `json:"author"`
	Capabilities []string    `json:"capabilities"`
	Enabled      bool        `json:"enabled"`
	HasNetwork   bool        `json:"has_network"`
	Integration  interface{} `json:"integration"`
	Jobs         interface{} `json:"jobs"`
	ModulePath   string      `json:"module_path"`
	Name         string      `json:"name"`
	ProtocolType string      `json:"protocol_type"`
	Uuid         string      `json:"uuid"`
	Website      string      `json:"website"`
}

func pluginsInstall(proxyReq *nrest.Service, opt *nrest.ReqOpt, ctx *gin.Context) {

	getState := proxyReq.Do(nrest.GET, PluginsUrls.State, opt)
	if getState.StatusCode != 200 {
		reposeHandler("", errors.New("error on plugin get state"), ctx)
		return
	}
	deleteState := proxyReq.Do(nrest.DELETE, PluginsUrls.State, opt)
	if deleteState.StatusCode != 200 {
		reposeHandler("", errors.New("error on plugin delete state"), ctx)
		return
	}
	appDownload := proxyReq.Do(nrest.POST, PluginsUrls.Download, opt)
	if appDownload.StatusCode != 200 {
		reposeHandler("", errors.New("error on plugin download"), ctx)
		return
	}
	retryCount := 0
	for {
		retryCount++
		req := proxyReq.Do(nrest.GET, PluginsUrls.State, opt)
		state := new(rubix.AppsDownloadState)
		req.ToInterface(&state)
		time.Sleep(4 * time.Second)
		log.Println("DOWNLOADING count", retryCount, "STATE", state.State, "CODE", req.StatusCode)
		if state.State == "DOWNLOADED" {
			break
		}
		if retryCount > 30 {
			reposeHandler("", errors.New("error on plugin download: tried to download the plugin 30 times and failed"), ctx)
			return
		}
	}
	appInstall := proxyReq.Do(nrest.POST, PluginsUrls.Install, opt)
	if appInstall.StatusCode != 200 {
		reposeHandler("", errors.New("error on plugin install plugin failed"), ctx)
		return
	}
	deleteState = proxyReq.Do(nrest.DELETE, PluginsUrls.State, opt)
	if deleteState.StatusCode != 200 {
		reposeHandler("", errors.New("error on plugin delete state"), ctx)
		return
	} else {
		res, err := appInstall.AsJson()
		if err != nil {
			reposeHandler(res, err, ctx)
			return
		}
		reposeHandler(res, nil, ctx)
		return
	}

}

func pluginsInstalled(proxyReq *nrest.Service, opt *nrest.ReqOpt, ctx *gin.Context) (installedPlugins []blankSlice) {
	appInstall := proxyReq.Do(nrest.GET, PluginsUrls.GetPlugins, opt)
	if appInstall.StatusCode != 200 {
		reposeHandler("", errors.New("error on get installed plugins"), ctx)
		return
	} else {
		res, err := appInstall.AsJson()
		if err != nil {
			reposeHandler(res, err, ctx)
			return
		}

		var data pluginsList
		err = appInstall.ToInterface(&data)
		if err != nil {
			reposeHandler("", errors.New("error on get installed plugins"), ctx)
			return
		}
		var plugins []blankSlice
		for _, a := range data {
			plugins = append(plugins, blankSlice{"plugin": a.Name})
		}
		return plugins
	}
}

func pluginsUninstall(proxyReq *nrest.Service, opt *nrest.ReqOpt, ctx *gin.Context, pluginArr []blankSlice) {
	opt.Json = pluginArr
	appInstall := proxyReq.Do(nrest.POST, PluginsUrls.UnInstall, opt)
	fmt.Println(appInstall.StatusCode)
	fmt.Println(appInstall.AsString())
	if appInstall.StatusCode != 200 {
		reposeHandler("", errors.New("error on get installed plugins"), ctx)
		return
	} else {
		res, err := appInstall.AsJson()
		if err != nil {
			reposeHandler(res, err, ctx)
			return
		} else {
			reposeHandler(res, err, ctx)
			return
		}
	}
}

func restartApp(proxyReq *nrest.Service, opt *nrest.ReqOpt, ctx *gin.Context) {

	var apps []blankSlice
	apps = append(apps, blankSlice{"service": "FLOW_FRAMEWORK", "action": "restart"})
	opt.Json = apps
	appInstall := proxyReq.Do(nrest.POST, PluginsUrls.Apps, opt)
	if appInstall.StatusCode != 200 {
		reposeHandler("", errors.New("error on get installed plugins"), ctx)
		return
	} else {
		res, err := appInstall.AsJson()
		if err != nil {
			reposeHandler(res, err, ctx)
			return
		} else {
			reposeHandler(res, err, ctx)
			return
		}
	}
}

func (base *Controller) PluginFullInstall(ctx *gin.Context) {
	body, err := bodyAsJSON(ctx)
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
	installed := pluginsInstalled(proxyReq, opt, ctx)
	fmt.Println(1111)
	fmt.Println(installed)
	fmt.Println(1111)
	pluginsUninstall(proxyReq, opt, ctx, installed)
	pluginsInstall(proxyReq, opt, ctx)
	restartApp(proxyReq, opt, ctx)
	fmt.Println(22222)
	//
}

//FlowFrameworkUpgrade UPDATE FF
//check what plugins where installed and enabled
//delete old plugins
//install new version of FF
//install existing plugins and enable as required
func (base *Controller) FlowFrameworkUpgrade(ctx *gin.Context) {
	body, err := bodyAsJSON(ctx)
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
	//get installed plugins
	pluginsInstalled(proxyReq, opt, ctx)

	//pluginsInstall(proxyReq, opt, ctx)
}

//UpdatePlugins full install of the plugins as in upload, unzip and restart flow framework
//func (base *Controller) UpdatePlugins(ctx *gin.Context) {
//	var msg TMsg
//	msg.Topic = "plugins.update"
//	msg.Message = "start update of plugins"
//	base.publishMSG(msg)
//	body := uploadBody(ctx)
//	uuid := ctx.Params.ByName("uuid")
//	result := new(TUpdatePlugins)
//	result.RemoveOldPlugins = "PASS: removed old plugins"
//	result.MakeUploadDir = "PASS: made tmp upload dir"
//	result.UpLoadPlugins = "PASS: upload and unzip plugins"
//	result.CleanUp = "PASS: deleted tmp upload dir"
//	result.RestartFF = "PASS: deleted tmp upload dir"
//	_, err := base.clearDir(uuid, "/data/flow-framework/data/plugins")
//	if err != nil {
//		result.RemoveOldPlugins = "PASS or FAIL: failed to remove existing OR there where no existing plugins installed"
//		msg.Message = result.RemoveOldPlugins
//		base.publishMSG(msg)
//	} else {
//		msg.Message = result.RemoveOldPlugins
//		base.publishMSG(msg)
//	}
//	//_, err = base.mkDir(uuid, body.ToPath, false)
//	if err != nil {
//		result.MakeUploadDir = "FAIL: failed to make OR new dir or was exiting"
//		msg.Message = result.MakeUploadDir
//		base.publishMSG(msg)
//	} else {
//		msg.Message = result.MakeUploadDir
//		base.publishMSG(msg)
//	}
//	err = base.uploadZip(uuid, body)
//	if err != nil {
//		result.UpLoadPlugins = fmt.Sprint(err)
//		msg.Message = result.UpLoadPlugins
//		base.publishMSG(msg)
//	} else {
//		msg.Message = result.UpLoadPlugins
//		base.publishMSG(msg)
//	}
//	_, err = base.rmDir(uuid, body.ToPath, false)
//	if err != nil {
//		result.CleanUp = fmt.Sprint(err)
//		msg.Message = result.CleanUp
//		base.publishMSG(msg)
//	} else {
//		msg.Message = result.CleanUp
//		base.publishMSG(msg)
//	}
//	//_, err = base.runCommand(uuid, "sudo systemctl restart nubeio-flow-framework.service", true)
//	//if err != nil {
//	//	msg.Message = result.RestartFF
//	//	base.publishMSG(msg)
//	//} else {
//	//	msg.Message = result.RestartFF
//	//	base.publishMSG(msg)
//	//}
//	reposeHandler(result, err, ctx)
//}

//func (base *Controller) UploadPlugins(ctx *gin.Context) {
//	body := uploadBody(ctx)
//	uuid := ctx.Params.ByName("uuid")
//	err := base.uploadZip(uuid, body)
//	if err != nil {
//		//return
//	}
//	reposeHandler("string(out)", err, ctx)
//}
//
//func (base *Controller) DeleteAllPlugins(ctx *gin.Context) {
//	uuid := ctx.Params.ByName("uuid")
//	dir, err := base.clearDir(uuid, "/data/flow-framework/data/plugins")
//	if err != nil {
//		reposeHandler(dir, err, ctx)
//	} else {
//		reposeHandler(dir, err, ctx)
//	}
//}
