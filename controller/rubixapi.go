package controller

//import (
//	"fmt"
//	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
//	"github.com/NubeIO/rubix-updater/model/rubix"
//	"github.com/NubeIO/rubix-updater/model/schema"
//	"github.com/gin-gonic/gin"
//	"time"
//)
//
//func (base *Controller) RubixHost(ctx *gin.Context) {
//	body, err := bodyAppsDownload(ctx)
//	po := proxyOptions{
//		ctx:          ctx,
//		refreshToken: true,
//		NonProxyReq:  true,
//	}
//	proxyReq, opt, rtn, err := base.buildReq(po)
//	if err != nil {
//		reposeHandler(nil, err, ctx)
//		return
//	}
//	opt = &nrest.ReqOpt{
//		Timeout:          500 * time.Second,
//		RetryCount:       0,
//		RetryWaitTime:    0 * time.Second,
//		RetryMaxWaitTime: 0,
//		Headers:          map[string]interface{}{"Authorization": rtn.Token},
//		Json:             body,
//	}
//
//	downloadCount := 0
//	//get state
//	getState := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
//	fmt.Println(getState.StatusCode)
//	fmt.Println(getState.AsString())
//	//delete state
//	deleteState := proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
//	fmt.Println(deleteState.StatusCode)
//	fmt.Println(deleteState.AsString())
//
//	appDownload := proxyReq.Do(nrest.POST, AppsUrls.Download, opt)
//	fmt.Println(appDownload.Err)
//	fmt.Println(appDownload.StatusCode)
//	fmt.Println(appDownload.AsString())
//
//	//
//	for {
//		req := proxyReq.Do(nrest.GET, AppsUrls.State, opt)
//		state := new(rubix.AppsDownloadState)
//		req.ToInterface(&state)
//		fmt.Println(req.Err)
//		fmt.Println(req.StatusCode)
//		fmt.Println(4444, state.State)
//		time.Sleep(4 * time.Second)
//		downloadCount++
//		fmt.Println("downloaded")
//		if state.State == "DOWNLOADED" {
//			break
//		}
//	}
//	appInstall := proxyReq.Do(nrest.POST, AppsUrls.Install, opt)
//	fmt.Println(appInstall.Err)
//	fmt.Println(appInstall.StatusCode)
//	fmt.Println(appInstall.AsString())
//
//	deleteState = proxyReq.Do(nrest.DELETE, AppsUrls.State, opt)
//	fmt.Println(deleteState.StatusCode)
//	fmt.Println(deleteState.AsString())
//
//}
//
//func (base *Controller) RubixPlatSchema(ctx *gin.Context) {
//	reposeHandler(schema.GetRubixPlatSchema(), nil, ctx)
//}
//
//func (base *Controller) RubixDiscoverSchema(ctx *gin.Context) {
//	reposeHandler(schema.GetRubixDiscover(), nil, ctx)
//}
//
//func (base *Controller) RubixMasterSchema(ctx *gin.Context) {
//	reposeHandler(schema.GetRubixSlaves(), nil, ctx)
//}
