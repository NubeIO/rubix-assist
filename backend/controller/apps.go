package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest"
	"github.com/NubeIO/rubix-updater/service/rubixapi"
	"github.com/NubeIO/rubix-updater/service/rubixmodel"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (base *Controller) restGetToken(id string) (token string, err error) {
	cli := rubixapi.New()
	host, err := base.GetHostDB(id)
	if err != nil {
		return "", err
	}

	h := fmt.Sprintf("http://%s:%d", host.IP, host.RubixPort)
	var rb = rest.RequestBuilder{
		Timeout:        25000 * time.Millisecond,
		BaseURL:        h,
		ContentType:    rest.JSON,
		DisableCache:   false,
		DisableTimeout: false,
	}
	var tb TokenBody
	tb.Username = host.RubixUsername
	tb.Password =  host.RubixPassword
	var r rubixapi.Req
	r.RequestBuilder = &rb
	r.URL = "/api/users/login"
	r.Method = rubixapi.POST
	r.Body = tb
	t, _, err := cli.GetToken(r)
	return  t.AccessToken, err
}

func (base *Controller) restReqBuilder(id string) (r rubixapi.Req, restClient *rubixapi.RestClient, err error) {
	host, err := base.GetHostDB(id)
	ip := fmt.Sprintf("http://%s:%d", host.IP, host.RubixPort)
	token, err := base.restGetToken(id)
	if err != nil {
		fmt.Println("ERROR ON GET RUBIX TOKEN", err)
	}
	fmt.Println("RUBIX IP:", ip)
	fmt.Println("RUBIX TOKEN:", token)
	headers := make(http.Header)
	headers.Add("Authorization", token)
	var rb = rest.RequestBuilder{
		Headers:        headers,
		Timeout:        300000 * time.Millisecond,
		BaseURL:        ip,
		ContentType:    rest.JSON,
		DisableCache:   false,
		DisableTimeout: false,
	}
	var rr rubixapi.Req
	cli := rubixapi.New()
	rr.RequestBuilder = &rb
	return rr, cli, nil
}


func (base *Controller) GetApps(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	r, restCli,  err := base.restReqBuilder(id)
	apps, err := restCli.AppsInstalled(r)
	if err != nil {
		reposeHandler(apps, err, ctx)
	}
	//aa := *apps
	//for i, a := range aa {
	//	fmt.Println(i, a.AppType)
	//}
	reposeHandler(apps, err, ctx)
}

func bodyAppsDownload(ctx *gin.Context) (dto *rubixmodel.AppsDownload, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func bodyInterface(ctx *gin.Context) (dto interface{}, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}


func (base *Controller) DownloadApp(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	r, restCli,  err := base.restReqBuilder(id)
	body, _ := bodyAppsDownload(ctx)
	r.Body = body
	apps, err := restCli.AppsDownload(r)
	if err != nil {
		reposeHandler(apps, err, ctx)
	} else {
		reposeHandler(apps, err, ctx)
	}

}

func (base *Controller) GetDownloadSate(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	r, restCli,  err := base.restReqBuilder(id)
	apps, err := restCli.AppsDownloadState(r)
	if err != nil {
		reposeHandler(apps, err, ctx)
	} else {
		reposeHandler(apps, err, ctx)
	}
}


func (base *Controller) DeleteDownloadSate(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	r, restCli,  err := base.restReqBuilder(id)
	apps, err := restCli.AppsDeleteDownloadState(r)
	if err != nil {
		reposeHandler(apps, err, ctx)
	} else {
		reposeHandler(apps, err, ctx)
	}
}
//FullInstall do a download and install
func (base *Controller) FullInstall(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	r, restCli,  _ := base.restReqBuilder(id)
	body, _ := bodyAppsDownload(ctx)
	r.Body = body
	_, err = restCli.AppsInstall(r)

	var downloadState int
	downloadState = 1

	for ok := true; ok; ok = downloadState != 2 {
		//n, err := downloadState
		if downloadState < 1  {
			fmt.Println("invalid input")
			break
		}
		time.Sleep(2 * time.Second)
		switch downloadState {
		case 1:
			state, _ := restCli.AppsDownloadState(r)

			fmt.Println("state", state)

			if state.State == "DOWNLOADING" {
				//downloadState = 1
			} else if  state.State == "DOWNLOADED" {
				fmt.Println("cause 1")
				downloadState = 5

			}  else if  state.State == "CLEARED" {
				fmt.Println("cause 1")
				//downloadState = 5
			}
		case 5:
			fmt.Println("case 5")
			restCli.AppsInstall(r)
			// Do nothing (we want to exit the loop)
			// In a real program this could be cleanup
		default:
			fmt.Println("not clear")
		}
	}

	//if err != nil {
	//	reposeHandler(apps, err, ctx)
	//} else {
	//	reposeHandler(apps, err, ctx)
	//}
}


func (base *Controller) InstallApp(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	r, restCli,  err := base.restReqBuilder(id)
	body, _ := bodyAppsDownload(ctx)
	r.Body = body
	apps, err := restCli.AppsInstall(r)
	if err != nil {
		reposeHandler(apps, err, ctx)
	} else {
		reposeHandler(apps, err, ctx)
	}
}
