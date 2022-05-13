package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func proxyPath(u string) (ok bool, path, proxyPath string) {
	p, err := url.Parse(fmt.Sprintf("http://%s", u))
	if err != nil {
		return false, path, proxyPath
	}
	parts := strings.SplitAfter(p.String(), "proxy")
	if len(parts) >= 1 {
		s := strings.Split(parts[1], "/")
		if len(s) > 1 {
			proxyPath = s[1]
			path = strings.ReplaceAll(parts[1], fmt.Sprintf("/%s", proxyPath), "")
			ok = true
		}
	}
	return
}

func (base *Controller) RubixProxyRequest(ctx *gin.Context) {

	method := getMethod(ctx.Request.Method)
	ok, path, rubixProxy := proxyPath(ctx.Request.URL.String())
	if !ok {
		fmt.Println("failed to set proxy path", err)
	}

	body, err := bodyAsJSON(ctx)
	if err != nil {
		fmt.Println("proxy err", err)
	}

	host, b, err := base.resolveHost(ctx)
	if err != nil {
		return
	}

	fmt.Println(host, b, err)
	restService := &rest.Service{}
	restService.Url = "192.168.15.191"
	restService.Port = 1616
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = true
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	nubeProxy.RubixProxyPath = rubixProxy
	restService.NubeProxy = nubeProxy

	base.Rest = restService
	req := restService.
		SetMethod(method).
		SetPath(path).
		SetBody(body).
		DoRequest()
	response := base.Rest.RestResponse(req, nil)
	if response.GetError() != nil {
		reposeHandler(nil, response.GetError(), ctx)
	} else {
		reposeHandler(response.AsJsonNoErr(), nil, ctx)
	}

	return
}
