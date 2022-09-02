package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/helpers/ip"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func setExternalToken(token string) string {
	return fmt.Sprintf("External %s", token)
}

// first login to assist and get an JWT token to generate the assist-token, store this token is the UI DB ("/api/users/login")
// first login to edge and get an JWT token to generate the edge-token ("/api/users/login")
// store that token in the host as EdgeToken

func (inst *Controller) Proxy(c *gin.Context) {
	host, err := inst.resolveHost(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	remote, err := ip.Builder(host.IP, host.Port)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}

	token := host.RubixToken // rubix-ui must first get and store the token (by using the uname/pass)
	if token == "" {
		// responseHandler(nil, errors.New("rubix-edge token is empty"), c)
		// return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		req.Header.Set("Authorization", setExternalToken(token))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
