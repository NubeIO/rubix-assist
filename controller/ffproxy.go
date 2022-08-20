package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/helpers/ip"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (inst *Controller) FFProxy(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	remote, err := ip.Builder(host.IP, 1660)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
