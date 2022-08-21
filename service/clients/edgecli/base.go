package edgecli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Client struct {
	Rest  *resty.Client
	URL   string `json:"url"`
	Port  int    `json:"port"`
	HTTPS bool   `json:"https"`
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
		port = 1661
	}
	if cli.HTTPS {
		cli.Rest.SetBaseURL(fmt.Sprintf("https://%s:%d", url, port))
	} else {
		cli.Rest.SetBaseURL(fmt.Sprintf("http://%s:%d", url, port))
	}
	return cli
}

func (inst *Client) SetTokenHeader(token string) *Client {
	inst.Rest.Header.Set("Authorization", setToken(token))
	return inst
}

func setToken(token string) string {
	return fmt.Sprintf("External %s", token)
}
