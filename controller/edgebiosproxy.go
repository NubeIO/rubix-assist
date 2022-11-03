package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/helpers/ip"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (inst *Controller) EdgeBiosProxy(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	remote, err := ip.Builder(host.HTTPS, host.IP, host.BiosPort)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		jwtToken := c.Param("jwt_token")
		if jwtToken == "" {
			req.Header.Set("Authorization", composeExternalToken(host.ExternalToken))
		} else {
			req.Header.Set("Authorization", jwtToken)
		}

	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
