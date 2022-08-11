package edgecli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type path struct {
	path string
}

var paths = struct {
	Apps   path
	Store  path
	System path
}{
	Apps:   path{path: "/api/apps"},
	Store:  path{path: "/api/appstore"},
	System: path{path: "/api/system"},
}

type Client struct {
	Rest *resty.Client
}

// New returns a new instance of the nube common apis
func New(url string, port int) *Client {
	rest := &Client{
		Rest: resty.New(),
	}
	if url == "" {
		url = "0.0.0.0"
	}
	if port == 0 {
		port = 1661
	}
	rest.Rest.SetBaseURL(fmt.Sprintf("http://%s:%d", url, port))
	return rest
}

func (response Response) buildResponse(resp *resty.Response, err error) *Response {
	response.StatusCode = resp.StatusCode()

	if resp.IsError() {
		response.Message = resp.Error()
	}
	if resp.StatusCode() == 0 {
		response.Message = "server is unreachable"
		response.StatusCode = 503
	}
	return &response
}
