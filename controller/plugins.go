package controller

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/rubix-assist/model/rubix"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"reflect"
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

func pluginsInstall(proxyReq *nrest.Service, opt *nrest.ReqOpt) (response interface{}, err error) {

	if reflect.ValueOf(opt.Json).IsValid() {
		if reflect.ValueOf(opt.Json).Len() < 1 {
			return "", errors.New("plugins body was empty")
		}
	} else {
		return "", errors.New("invalid plugins body")
	}

	getState := proxyReq.Do(nrest.GET, PluginsUrls.State, opt)
	if getState.StatusCode != 200 {
		return "", errors.New("error on get state")
	}
	deleteState := proxyReq.Do(nrest.DELETE, PluginsUrls.State, opt)
	if deleteState.StatusCode != 200 {
		return "", errors.New("error on delete state")

	}
	log.Info("try download plugins")

	appDownload := proxyReq.Do(nrest.POST, PluginsUrls.Download, opt)
	if appDownload.StatusCode != 200 {
		log.Error("try download plugins failed")
		return "", errors.New("error on download")
	}
	retryCount := 0
	for {
		retryCount++
		req := proxyReq.Do(nrest.GET, PluginsUrls.State, opt)
		state := new(rubix.AppsDownloadState)
		req.ToInterface(&state)
		time.Sleep(4 * time.Second)
		log.Println("DOWNLOADING PLUGINS count", retryCount, "STATE", state.State, "CODE", req.StatusCode)
		if state.State == "DOWNLOADED" {
			log.Info("DOWNLOADED PLUGINS")
			break
		}
		if retryCount > 30 {
			return "", errors.New("error on download: tried to download 30 times and failed")
		}
	}
	log.Info("try install plugins")
	appInstall := proxyReq.Do(nrest.POST, PluginsUrls.Install, opt)
	if appInstall.StatusCode != 200 {
		log.Error("try install plugins failed")
		return "", errors.New("error install failed")
	}
	log.Info("try delete state plugins")
	deleteState = proxyReq.Do(nrest.DELETE, PluginsUrls.State, opt)
	if deleteState.StatusCode != 200 {
		log.Error("error on delete state")
		return "", errors.New("error on delete state")
	} else {
		res, err := appInstall.AsJson()
		if err != nil {
			return res, nil

		}
		return res, nil
	}
}

func pluginsInstalled(proxyReq *nrest.Service, opt *nrest.ReqOpt) (installedPlugins []blankSlice, nilPlugins bool, err error) {
	req := proxyReq.Do(nrest.GET, PluginsUrls.GetPlugins, opt)
	if req.StatusCode != 200 {
		return nil, true, errors.New("error on http req to get installed plugins")
	} else {
		var data pluginsList
		err = req.ToInterface(&data)
		if err != nil {
			return nil, true, errors.New("error on body return from get plugins")
		}
		var plugins []blankSlice
		for _, a := range data {
			plugins = append(plugins, blankSlice{"plugin": a.Name})
		}
		if len(plugins) == 0 {
			nilPlugins = true
		}
		return plugins, nilPlugins, nil
	}
}

func pluginsUninstall(proxyReq *nrest.Service, opt *nrest.ReqOpt, pluginArr []blankSlice) (response interface{}, err error) {
	opt.Json = pluginArr
	appInstall := proxyReq.Do(nrest.POST, PluginsUrls.UnInstall, opt)
	if appInstall.StatusCode != 200 {
		return nil, errors.New("error on http req to get uninstall plugins")
	} else {
		res, err := appInstall.AsJson()
		if err != nil {
			return nil, errors.New("error on http repose body uninstall plugins")
		} else {
			return res, nil
		}
	}
}

func restartApp(proxyReq *nrest.Service, opt *nrest.ReqOpt, apps []blankSlice) (ok bool, err error) {
	opt.Json = apps
	r := proxyReq.Do(nrest.POST, PluginsUrls.Apps, opt)
	if r.StatusCode != 200 {
		return false, errors.New("failed to restart flow-framework")
	} else {
		return true, nil
	}
}

func (base *Controller) PluginFullUninstall(ctx *gin.Context) {
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
	//uninstall
	r := proxyReq.Do(nrest.POST, PluginsUrls.UnInstall, opt)
	if r.StatusCode != 200 {
		reposeHandler(nil, errors.New("error on http req to get uninstall plugins"), ctx)
	}
	res, err := r.AsJson()
	//restart FF to load the plugins
	var apps []blankSlice
	apps = append(apps, blankSlice{"service": "FLOW_FRAMEWORK", "action": "restart"})
	log.Info("try and re-start flow-framework")
	_, err = restartApp(proxyReq, opt, apps)
	if err != nil {
		log.Errorln("try and re-start flow-framework failed")
		reposeHandler(nil, err, ctx)
	}
	reposeHandler(res, err, ctx)
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
	install, err := pluginsInstall(proxyReq, opt)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	//restart FF to load the plugins
	var apps []blankSlice
	apps = append(apps, blankSlice{"service": "FLOW_FRAMEWORK", "action": "restart"})
	log.Info("try and re-start flow-framework")
	_, err = restartApp(proxyReq, opt, apps)
	if err != nil {
		log.Errorln("try and re-start flow-framework failed")
		reposeHandler(nil, err, ctx)
	}
	reposeHandler(install, err, ctx)

}

type pluginsResponse struct {
	GetInstalledPlugins           string `json:"get_installed_plugins"`
	DeletePlugins                 string `json:"delete_plugins"`
	ReInstallFlowFramework        string `json:"reinstall_flow_framework"`
	ReInstallFlowFrameworkPlugins string `json:"reinstall_flow_framework_plugins"`
	ReStartFlowFramework          string `json:"restart_flow_framework"`
	Error                         string `json:"error"`
}

//FlowFrameworkUpgrade UPDATE FF
//check what plugins where installed and enabled
//delete old plugins
//install new version of FF
//install existing plugins and enable as required
func (base *Controller) FlowFrameworkUpgrade(ctx *gin.Context) {
	httpRes := new(pluginsResponse)
	body, err := bodyAppsDownload(ctx)
	po := proxyOptions{
		ctx:          ctx,
		refreshToken: true,
		NonProxyReq:  true,
	}
	proxyReq, opt, rtn, err := base.buildReq(po)
	if err != nil {
		httpRes.Error = err.Error()
		reposeHandler(nil, errors.New(httpRes.Error), ctx)
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
	var uninstall interface{}
	log.Info("try get installed plugins")
	installed, nilPlugins, err := pluginsInstalled(proxyReq, opt)
	httpRes.GetInstalledPlugins = "pass"
	if !nilPlugins {
		log.Info("try and uninstalled plugins")
		uninstall, err = pluginsUninstall(proxyReq, opt, installed)
		httpRes.DeletePlugins = "pass"
		if err != nil {
			httpRes.DeletePlugins = "error on uninstall plugins"
			httpRes.GetInstalledPlugins = "error on uninstall plugins"
		}
	} else {
		httpRes.GetInstalledPlugins = "no plugins where installed"
		httpRes.DeletePlugins = "no plugins where installed"
	}
	//updateFF
	log.Info("try and re-install flow-framework")
	opt.Json = body
	httpRes.ReInstallFlowFramework = "pass"
	_, err = appsInstall(proxyReq, opt)
	if err != nil {
		httpRes.ReInstallFlowFramework = err.Error()

	}
	//reinstall plugins
	log.Info("try and re-install flow-framework PLUGINS", uninstall)
	opt.Json = installed
	_, err = pluginsInstall(proxyReq, opt)
	httpRes.ReInstallFlowFrameworkPlugins = "pass"
	if err != nil {
		httpRes.ReInstallFlowFrameworkPlugins = err.Error()
	}
	//restart FF to load the plugins
	var apps []blankSlice
	apps = append(apps, blankSlice{"service": "FLOW_FRAMEWORK", "action": "restart"})
	log.Info("try and re-start flow-framework")
	_, err = restartApp(proxyReq, opt, apps)
	httpRes.ReStartFlowFramework = "pass"
	if err != nil {
		log.Errorln("try and re-start flow-framework failed")
		httpRes.ReStartFlowFramework = err.Error()
	}
	reposeHandler(httpRes, err, ctx)
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
