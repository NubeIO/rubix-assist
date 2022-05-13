package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest/v1/rest"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
	"time"
)

func getMethod(method string) (out string) {
	out = rest.GET
	switch method {
	case "GET":
		out = rest.GET
	case "PATCH":
		out = rest.PATCH
	case "DELETE":
		out = rest.DELETE
	case "POST":
		out = rest.POST
	case "PUT":
		out = rest.PUT
	}
	return out
}

func urlProxyPath(u string, nonProxyReq bool) (clean string) {
	p, err := url.Parse(fmt.Sprintf("http://%s", u))
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
	//reqOpt       rest.ReqOpt
	NonProxyReq bool
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

//func (base *Controller) buildReq(proxyOptions proxyOptions) (s *nrest.Service, options *nrest.ReqOpt, rtn proxyReturn, err error) {
//	ctx := proxyOptions.ctx
//	host, _, err := base.resolveHost(ctx)
//	if err != nil {
//		return nil, nil, rtn, err
//	}
//	method := getMethod(ctx.Request.Method)
//	rtn.Method = method
//	ru := urlProxyPath(ctx.Request.URL.String(), proxyOptions.NonProxyReq)
//	rtn.RequestURL = ru
//	body, err := bodyAsJSON(ctx)
//	rtn.Body = body
//
//	http := "http"
//	if bools.IsTrue(host.HTTPS) {
//		http = "https"
//	}
//
//	ip := fmt.Sprintf("%s://%s:%d", http, host.IP, host.RubixPort)
//	if host.RubixPort == 0 {
//		ip = fmt.Sprintf("%s://%s", http, host.IP)
//	}
//	s = &nrest.Service{
//		BaseUri: ip,
//	}
//	token := host.RubixToken
//	if token == "" || proxyOptions.refreshToken || tokenTimeDiffMin(host.RubixTokenLastUpdate, 15) {
//		options = &nrest.ReqOpt{
//			Timeout:          2 * time.Second,
//			RetryCount:       2,
//			RetryWaitTime:    2 * time.Second,
//			RetryMaxWaitTime: 0,
//			Json:             map[string]interface{}{"username": host.RubixUsername, "password": host.RubixPassword},
//		}
//		req := s.Do(nrest.POST, "/api/users/login", options)
//		res := new(rubix.TokenResponse)
//		err = req.ToInterface(&res)
//		if err != nil {
//			return nil, nil, rtn, err
//		}
//		if res.AccessToken == "" {
//			return nil, nil, rtn, errors.New("ERROR: failed to get token")
//		}
//		token = res.AccessToken
//		rtn.Token = token
//		var h model.Host
//		h.RubixToken = token
//		h.RubixTokenLastUpdate = time.Now()
//		_, err := base.DB.UpdateHost(host.UUID, &h)
//		if err != nil {
//			log.Println("ERROR: failed to update host token in db", err)
//			return nil, nil, rtn, errors.New("ERROR: failed to update host token in db")
//		}
//	}
//	return s, options, rtn, nil
//}

//func (base *Controller) RubixProxyRequest(ctx *gin.Context) {
//
//	method := getMethod(ctx.Request.Method)
//	rtn.Method = method
//	ru := urlProxyPath(ctx.Request.URL.String(), proxyOptions.NonProxyReq)
//	rtn.RequestURL = ru
//	body, err := bodyAsJSON(ctx)
//
//	restService := &rest.Service{}
//	restService.Url = "192.168.15.191"
//	restService.Port = 1616
//	restOptions := &rest.Options{}
//	restService.Options = restOptions
//	restService = rest.New(restService)
//
//	//po := proxyOptions{
//	//	ctx:          ctx,
//	//	refreshToken: true,
//	//}
//	//proxyReq, opt, rtn, err := base.buildReq(po)
//	//if err != nil {
//	//	reposeHandler(nil, err, ctx)
//	//} else {
//	//	opt = &nrest.ReqOpt{
//	//		Timeout:          500 * time.Second,
//	//		RetryCount:       0,
//	//		RetryWaitTime:    0 * time.Second,
//	//		RetryMaxWaitTime: 0,
//	//		Headers:          map[string]interface{}{"Authorization": rtn.Token},
//	//		Json:             rtn.Body,
//	//	}
//	//
//	//	_url := rtn.RequestURL
//	//	//get query parameters eg: ?thing=true
//	//	parts := strings.SplitAfter(rtn.RequestURL, "?")
//	//	if len(parts) >= 2 {
//	//		opt.SetQueryString = parts[1]
//	//		_url = strings.TrimRight(parts[0], "?")
//	//	}
//	//	req := proxyReq.Do(rtn.Method, _url, opt)
//	//	fmt.Println(rtn.Method, _url, opt.Data)
//	//	json, _ := req.AsJson()
//	//	log.Println(rtn.RequestURL)
//	//	log.Println(req.Err)
//	//	log.Println(req.AsString())
//	//	log.Println(req.StatusCode)
//	//	if err != nil {
//	//		reposeHandler(nil, req.Err, ctx)
//	//	} else {
//	//		reposeHandler(json, nil, ctx)
//	//	}
//	//}
//}
