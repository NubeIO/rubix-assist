package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/service/rubixmodel"
	"github.com/NubeIO/rubix-updater/utils/rest"
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

func bodyAppsDownload(ctx *gin.Context) (dto *rubixmodel.AppsDownload, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func bodyInterface(ctx *gin.Context) (dto interface{}, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

var AppsUrls = struct {
	Install  string
	Download string
	State string
}{
	Install:  "/install",
	Download: "/download",
	State: "/state",

}

func (base *Controller) AppsRequest(ctx *gin.Context) {
	po := proxyOptions{
		ctx:          ctx,
		refreshToken: true,
		NonProxyReq: true,
	}
	proxyReq, opt, rtn, err := base.buildProxyReq(po)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	opt = &rest.ReqOpt{
		Timeout:          2 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    0 * time.Second,
		RetryMaxWaitTime: 0,
		Headers:          map[string]interface{}{"Authorization": rtn.Token},
		Json:             rtn.Body,
	}

	apps := "/apps"

	//downloaded := false
	//downloadCount := 0
	//
	//installation := false
	//installCount := 0

	switch rtn.RequestURL {
	case apps+AppsUrls.Install:
		req := proxyReq.Do(rtn.Method, rtn.RequestURL, opt)
		json, err := req.AsJson()
		fmt.Println(req.Err)
		fmt.Println(req.StatusCode)
		if err != nil {
			reposeHandler(nil, err, ctx)
		} else {
			reposeHandler(json, err, ctx)
		}
	case apps+AppsUrls.Download:
		req := proxyReq.Do(rtn.Method, rtn.RequestURL, opt)
		json, err := req.AsJson()
		fmt.Println(req.Err)
		fmt.Println(req.StatusCode)
		if err != nil {
			reposeHandler(nil, err, ctx)
		} else {
			reposeHandler(json, err, ctx)
		}
	}



	req := proxyReq.Do(rtn.Method, rtn.RequestURL, opt)
	json, err := req.AsJson()
	fmt.Println(req.Err)
	fmt.Println(req.StatusCode)
	if err != nil {
		reposeHandler(nil, err, ctx)
	} else {
		reposeHandler(json, err, ctx)
	}

}

func download(c int) bool {
	fmt.Println("downloading")
	if c == 3 {
		return true
	} else {
		return false
	}
}

func install(c int) bool {
	fmt.Println("install")
	if c == 3 {
		return true
	} else {
		return false
	}
}



//FullInstall do a download and install
func (base *Controller) fullInstall(ctx *gin.Context) {
	//checkState
	//start download and check state every 4 sec
	//once state is downloaded then start install
	//once download is completed the start install

	downloaded := false
	downloadCount := 0

	installation := false
	installCount := 0

	for {
		downloaded = download(downloadCount)
		time.Sleep(4 * time.Second)
		downloadCount++
		fmt.Println("downloaded")
		if downloaded {
			break
		}
	}
	for {
		installation = install(installCount)
		time.Sleep(4 * time.Second)
		fmt.Println("installation")
		installCount++
		if installation {
			break
		}
	}
	reposeHandler(222, err, ctx)

}

//for ok := true; ok; ok = downloadState != 2 {
//	//n, err := downloadState
//	if downloadState < 1  {
//		fmt.Println("invalid input")
//		break
//	}
//	time.Sleep(2 * time.Second)
//	switch downloadState {
//	case 1:
//		state, _ := restCli.AppsDownloadState(r)
//
//		fmt.Println("state", state)
//
//		if state.State == "DOWNLOADING" {
//			//downloadState = 1
//		} else if  state.State == "DOWNLOADED" {
//			fmt.Println("cause 1")
//			downloadState = 5
//
//		}  else if  state.State == "CLEARED" {
//			fmt.Println("cause 1")
//			//downloadState = 5
//		}
//	case 5:
//		fmt.Println("case 5")
//		restCli.AppsInstall(r)
//		// Do nothing (we want to exit the loop)
//		// In a real program this could be cleanup
//	default:
//		fmt.Println("not clear")
//	}
//}

//if err != nil {
//	reposeHandler(apps, err, ctx)
//} else {
//	reposeHandler(apps, err, ctx)
//}
//}
