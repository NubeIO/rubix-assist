package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
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
		c.JSON(http.StatusNotFound, amodel.Message{Message: "not found"})
		return
	}
	proxyPath = fmt.Sprintf("/%s", proxyPath)
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
		req.URL.Path = proxyPath
		authorization := c.GetHeader("jwt_token")
		if authorization == "" {
			req.Header.Set("Authorization", composeExternalToken(host.ExternalToken))
		} else {
			req.Header.Set("Authorization", authorization)
		}
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
