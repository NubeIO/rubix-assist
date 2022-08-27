package assitcli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Rest        *resty.Client
	URL         string `json:"url"`
	Port        int    `json:"port"`
	HTTPS       bool   `json:"https"`
	AssistToken string `json:"assist_token"`
}

type ResponseBody struct {
	Response ResponseCommon `json:"response"`
	Status   string         `json:"status"`
	Count    string         `json:"count"`
}

type ResponseCommon struct {
	UUID string `json:"uuid"`
	// Name        string `json:"name"`
}

func buildUrl(overrideUrl ...string) string {
	if len(overrideUrl) > 0 {
		if overrideUrl[0] != "" {
			return overrideUrl[0]
		}
	}
	return ""
}

func setExternalToken(token string) string {
	return fmt.Sprintf("External %s", token)
}

// New returns a new instance of the nube common apis
func New(cli *Client) *Client {
	if cli == nil {
		log.Fatal("rubix-service-rest-client can not be empty")
	}
	var url = cli.URL
	var port = cli.Port
	cli.Rest = resty.New()
	if url == "" {
		url = "0.0.0.0"
	}
	if port == 0 {
		port = 1662
	}
	if cli.HTTPS {
		cli.Rest.SetBaseURL(fmt.Sprintf("https://%s:%d", url, port))
	} else {
		cli.Rest.SetBaseURL(fmt.Sprintf("http://%s:%d", url, port))
	}
	if cli.AssistToken != "" {
		cli.Rest.SetHeader("Authorization", setExternalToken(cli.AssistToken))
	}
	return cli
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
	Wires        Path
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
	Wires:        Path{Path: "/api/wires"},
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
