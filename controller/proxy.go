package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ip"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"
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
	proxyPath := strings.Trim(c.Param("proxyPath"), string(os.PathSeparator))
	proxyPathParts := strings.Split(proxyPath, "/")
	var remote *url.URL = nil
	if len(proxyPathParts) > 0 && proxyPathParts[0] == "eb" {
		proxyPath = path.Join(proxyPathParts[1:]...)
		remote, err = ip.Builder(host.HTTPS, host.IP, host.BiosPort)
	} else if len(proxyPathParts) > 0 && proxyPathParts[0] == "edge" {
		proxyPath = path.Join(proxyPathParts[1:]...)
		remote, err = ip.Builder(host.HTTPS, host.IP, host.Port)
	} else {
		remote, err = ip.Builder(host.HTTPS, host.IP, host.Port)
	}
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	externalToken := host.ExternalToken
	proxyPath = fmt.Sprintf("/%s", proxyPath)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = proxyPath
		authorization := c.GetHeader("jwt-token")
		if authorization != "" {
			req.Header.Set("Authorization", authorization)
		} else if externalToken != "" {
			req.Header.Set("Authorization", composeExternalToken(externalToken))
		}
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
