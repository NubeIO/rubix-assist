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
		remote, err = ip.Builder(host.HTTPS, host.IP, host.BiosPort)
		proxyPath = path.Join(proxyPathParts[1:]...)
	} else {
		remote, err = ip.Builder(host.HTTPS, host.IP, host.Port)
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
		jwtToken := c.Query("jwt_token")
		fmt.Println("jwtToken", jwtToken)
		if jwtToken == "" {
			req.Header.Set("Authorization", composeExternalToken(host.ExternalToken))
		} else {
			req.Header.Set("Authorization", jwtToken)
		}
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
