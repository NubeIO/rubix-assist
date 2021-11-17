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

func (base *Controller)  restGetToken() (token string, err error) {
	cli := rubixapi.New()
	h := fmt.Sprintf("http://%s:%s", "123.209.84.230", "1616")
	var rb = rest.RequestBuilder{
		Timeout:        5000 * time.Millisecond,
		BaseURL:        h,
		ContentType:    rest.JSON,
		DisableCache:   false,
		DisableTimeout: false,
	}
	var tb TokenBody
	tb.Username = "admin"
	tb.Password = "N00BWires"
	var r rubixapi.Req
	r.RequestBuilder = &rb
	r.URL = "/api/users/login"
	r.Method = rubixapi.POST
	r.Body = tb
	t, _, err := cli.GetToken(r)
	return  t.AccessToken, err
}

func (base *Controller) restReqBuilder() (r rubixapi.Req, restClient *rubixapi.RestClient, err error) {
	h := fmt.Sprintf("http://%s:%s", "123.209.84.230", "1616")
	token, err := base.restGetToken()
	if err != nil {
		//return
	}
	headers := make(http.Header)
	headers.Add("Authorization", token)

	var rb = rest.RequestBuilder{
		Headers:        headers,
		Timeout:        5000 * time.Millisecond,
		BaseURL:        h,
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
	r, restCli,  err := base.restReqBuilder()
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

func bodyAppsDownload(ctx *gin.Context) (dto rubixmodel.AppsDownload, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) DownloadApp(ctx *gin.Context) {
	r, restCli,  err := base.restReqBuilder()
	body, _ := bodyAppsDownload(ctx)
	r.Body = body
	apps, err := restCli.AppsInstall(r)
	if err != nil {
		reposeHandler(apps, err, ctx)
	}
	//aa := *apps
	//for i, a := range aa {
	//	fmt.Println(i, a.AppType)
	//}
	reposeHandler(apps, err, ctx)
}
