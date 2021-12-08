package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/rubixmodel"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
	"time"
)

func getMethod(method string) (out string) {
	out = nrest.GET
	switch method {
	case "GET":
		out = nrest.GET
	case "PATCH":
		out = nrest.PATCH
	case "DELETE":
		out = nrest.DELETE
	case "POST":
		out = nrest.POST
	}
	return out
}

func (base *Controller) resolveHost(ctx *gin.Context) (host *model.Host, useID bool, err error) {
	idName, useID := useHostNameOrID(ctx)
	host, err = base.DB.GetHostByName(idName, useID)
	return host, useID, err
}

func urlProxyPath(u string, nonProxyReq bool) (clean string) {
	_url := fmt.Sprintf("http://%s", u)
	p, err := url.Parse(_url)
	if err != nil {
		return ""
	}
	var parts []string
	if nonProxyReq {
		parts = strings.SplitAfter(p.String(), "api")
	} else {
		parts = strings.SplitAfter(p.String(), "proxy")
	}

	if len(parts) >= 1 {
		return parts[1]
	} else {
		return ""
	}
}

type proxyOptions struct {
	ctx          *gin.Context
	refreshToken bool
	reqOpt       nrest.ReqOpt
	NonProxyReq  bool
}

type proxyReturn struct {
	Token, Method, RequestURL string
	Body                      interface{}
}

func tokenTimeDiffMin(t time.Time, timeDiff float64) bool {
	t1 := time.Now()
	if t1.Sub(t).Minutes() > timeDiff {
		return true
	} else {
		return false
	}
}

func (base *Controller) buildProxyReq(proxyOptions proxyOptions) (s *nrest.Service, options *nrest.ReqOpt, rtn proxyReturn, err error) {
	ctx := proxyOptions.ctx
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		return nil, nil, rtn, err
	}
	method := getMethod(ctx.Request.Method)
	rtn.Method = method
	ru := urlProxyPath(ctx.Request.URL.String(), proxyOptions.NonProxyReq)
	rtn.RequestURL = ru
	body, err := bodyAsJSON(ctx)
	rtn.Body = body
	ip := fmt.Sprintf("http://%s:%d", host.IP, host.RubixPort)
	s = &nrest.Service{
		BaseUri: ip,
	}
	token := host.RubixToken
	if token == "" || proxyOptions.refreshToken || tokenTimeDiffMin(host.RubixTokenLastUpdate, 15) {
		options = &nrest.ReqOpt{
			Timeout:          2 * time.Second,
			RetryCount:       2,
			RetryWaitTime:    2 * time.Second,
			RetryMaxWaitTime: 0,
			Json:             map[string]interface{}{"username": host.RubixUsername, "password": host.RubixPassword},
		}
		req := s.Do(nrest.POST, "/api/users/login", options)
		fmt.Println("REQ GET TOKEN:", req.AsString())
		res := new(rubixmodel.TokenResponse)
		err = req.ToInterface(&res)
		if err != nil {
			return nil, nil, rtn, err
		}
		if res.AccessToken == "" {
			return nil, nil, rtn, errors.New("ERROR: failed to get token")
		}
		token = res.AccessToken
		rtn.Token = token
		var h model.Host
		h.RubixToken = token
		h.RubixTokenLastUpdate = time.Now()
		_, err := base.DB.UpdateHost(host.ID, &h)
		if err != nil {
			fmt.Println("ERROR: failed to update host token in db", err)
			return nil, nil, rtn, errors.New("ERROR: failed to update host token in db")
		}
	}
	return s, options, rtn, nil
}

func (base *Controller) RubixProxyRequest(ctx *gin.Context) {
	po := proxyOptions{
		ctx:          ctx,
		refreshToken: true,
	}
	proxyReq, opt, rtn, err := base.buildProxyReq(po)
	if err != nil {
		reposeHandler(nil, err, ctx)
	} else {
		opt = &nrest.ReqOpt{
			Timeout:          500 * time.Second,
			RetryCount:       0,
			RetryWaitTime:    0 * time.Second,
			RetryMaxWaitTime: 0,
			Headers:          map[string]interface{}{"Authorization": rtn.Token},
			Json:             rtn.Body,
		}

		_url := rtn.RequestURL
		//get query parameters eg: ?thing=true
		parts := strings.SplitAfter(rtn.RequestURL, "?")
		if len(parts) >= 2 {
			opt.SetQueryString = parts[1]
			_url = strings.TrimRight(parts[0], "?")
		}
		req := proxyReq.Do(rtn.Method, _url, opt)
		json, _ := req.AsJson()
		fmt.Println(rtn.RequestURL)
		fmt.Println(req.Err)
		fmt.Println(req.AsString())
		fmt.Println(req.StatusCode)
		if err != nil {
			reposeHandler(nil, req.Err, ctx)
		} else {
			reposeHandler(json, nil, ctx)
		}
	}
}
