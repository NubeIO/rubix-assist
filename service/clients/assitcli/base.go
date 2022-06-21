package assitcli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Rest *resty.Client
}

// New returns a new instance of the nube common apis
func New(url string, port int) *Client {
	rest := &Client{
		Rest: resty.New(),
	}
	rest.Rest.SetBaseURL(fmt.Sprintf("http://%s:%d", url, port))
	return rest
}

type Path struct {
	Path string
}

var Paths = struct {
	Hosts        Path
	Ping         Path
	HostNetwork  Path
	Location     Path
	Users        Path
	Edge         Path
	Apps         Path
	Tasks        Path
	Transactions Path
	System       Path
	Networking   Path
}{
	Hosts:        Path{Path: "/api/hosts"},
	Ping:         Path{Path: "/api/system/ping"},
	HostNetwork:  Path{Path: "/api/networks"},
	Location:     Path{Path: "/api/locations"},
	Users:        Path{Path: "/api/locations"},
	Edge:         Path{Path: "/api/edgeapi"},
	Apps:         Path{Path: "/api/edgeapi/apps"},
	Tasks:        Path{Path: "/api/Tasks"},
	Transactions: Path{Path: "/api/transactions"},
	System:       Path{Path: "/api/system"},
	Networking:   Path{Path: "/api/networking"},
}

type Response struct {
	StatusCode int         `json:"code"`
	Message    interface{} `json:"message"`
	resty      *resty.Response
}

func (response Response) buildResponse(resp *resty.Response, err error) *Response {
	response.StatusCode = resp.StatusCode()
	response.resty = resp
	if resp.IsError() {
		response.Message = resp.Error()
	}
	if resp.StatusCode() == 0 {
		response.Message = "server is unreachable"
		response.StatusCode = 503
	}
	return &response
}
