package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ip"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func composeExternalToken(token string) string {
	return fmt.Sprintf("External %s", token)
}

func (inst *Controller) Proxy(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	remote, err := ip.Builder(host.HTTPS, host.IP, host.Port)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}

	token := host.ExternalToken
	if token == "" {
		responseHandler(nil, errors.New("rubix-edge token is empty"), c)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		req.Header.Set("Authorization", composeExternalToken(token))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
