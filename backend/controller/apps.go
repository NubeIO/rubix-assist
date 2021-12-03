package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/rubix-updater/service/rubixmodel"

	"github.com/gin-gonic/gin"
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

func (base *Controller) AppsFullInstall(ctx *gin.Context) {

	body, err := bodyAppsDownload(ctx)
	po := proxyOptions{
		ctx:          ctx,
		refreshToken: true,
		NonProxyReq:  true,
	}
	proxyReq, opt, rtn, err := base.buildProxyReq(po)
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
		state := new(rubixmodel.AppsDownloadState)
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
