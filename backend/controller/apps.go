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

type TokenBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type TokenResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	Message     *string `json:"message,omitempty"`
}

func bodyAppsDownload(ctx *gin.Context) (dto appsDownload, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

var AppsUrls = struct {
	Install  string
	Download string
	State    string
}{
	Install:  "/api/app/install",
	Download: "/api/app/download",
	State:    "/api/app/download_state",
}

type appsDownload []apps

type apps struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

func appsInstall(proxyReq *nrest.Service, opt *nrest.ReqOpt) (response interface{}, err error) {
	fmt.Println(opt.Json)

	getState := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
	if getState.StatusCode != 200 {
		return "", errors.New("error on get state")
	}
	log.Info("try and delete app state")
	deleteState := proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
	if deleteState.StatusCode != 200 {
		log.Error("try and delete app state failed")
		return "", errors.New("error on delete state")

	}
	log.Info("try and download app")
	appDownload := proxyReq.Do(nrest.POST, AppsUrls.Download, opt)
	if appDownload.StatusCode != 200 {
		log.Error("try and download app failed")
		return "", errors.New("error on download")
	}
	retryCount := 0
	for {
		retryCount++
		req := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
		state := new(rubix.AppsDownloadState)
		req.ToInterface(&state)
		time.Sleep(4 * time.Second)
		log.Println("DOWNLOADING APP count", retryCount, "STATE", state.State, "CODE", req.StatusCode)
		if state.State == "DOWNLOADED" {
			break
		}
		if retryCount > 30 {
			return "", errors.New("error on download: tried to download 30 times and failed")
		}
	}
	log.Info("try and install app")
	appInstall := proxyReq.Do(nrest.POST, AppsUrls.Install, opt)
	if appInstall.StatusCode != 200 {
		log.Error("try and install app failed")
		return "", errors.New("error install failed")
	}
	deleteState = proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
	if deleteState.StatusCode != 200 {
		log.Error("try delete  app state failed")
		return "", errors.New("error on delete state")
	} else {
		res, err := appInstall.AsJson()
		if err != nil {
			return res, nil

		}
		return res, nil
	}
}

func (base *Controller) AppsFullInstall(ctx *gin.Context) {
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

	install, err := appsInstall(proxyReq, opt)
	if err != nil {
		reposeHandler(nil, err, ctx)
	} else {
		reposeHandler(install, nil, ctx)
	}

	//downloadCount := 0
	////get state
	//getState := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
	//fmt.Println(getState.StatusCode)
	//fmt.Println(getState.AsString())
	//
	//var msg TMsg
	//msg.Topic = "apps.install.state"
	//msg.Message = getState.AsString()
	//base.publishMSG(msg)
	////delete state
	//deleteState := proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
	//fmt.Println(deleteState.StatusCode)
	//fmt.Println(deleteState.AsString())
	//
	//appDownload := proxyReq.Do(nrest.POST, AppsUrls.Download, opt)
	//fmt.Println(appDownload.Err)
	//fmt.Println(appDownload.StatusCode)
	//fmt.Println(appDownload.AsString())
	//
	//msg.Topic = "apps.install.download"
	//msg.Message = "download completed"
	//base.publishMSG(msg)
	//
	////
	//for {
	//	req := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
	//	state := new(rubix.AppsDownloadState)
	//	req.ToInterface(&state)
	//	fmt.Println(req.Err)
	//	fmt.Println(req.StatusCode)
	//	fmt.Println(4444, state.State)
	//	time.Sleep(4 * time.Second)
	//	downloadCount++
	//	fmt.Println("downloaded")
	//	msg.Topic = "apps.install.downloading"
	//	msg.Message = state.State
	//	base.publishMSG(msg)
	//	if state.State == "DOWNLOADED" {
	//		break
	//	}
	//	msg.Topic = "apps.install.downloaded"
	//	msg.Message = "DOWNLOADED"
	//	base.publishMSG(msg)
	//}
	//appInstall := proxyReq.Do(nrest.POST, AppsUrls.Install, opt)
	//fmt.Println(appInstall.Err)
	//fmt.Println(appInstall.StatusCode)
	//fmt.Println(appInstall.AsString())
	//msg.Topic = "apps.install.install"
	//msg.Message = appInstall.AsString()
	//base.publishMSG(msg)
	//
	//deleteState = proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
	//fmt.Println(deleteState.StatusCode)
	//fmt.Println(deleteState.AsString())

}
