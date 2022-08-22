package assistapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

func setExternalToken(token string) string {
	return fmt.Sprintf("External %s", token)
}

type Message struct {
	Message string `json:"message"`
}

type Client struct {
	rest        *resty.Client
	URL         string `json:"url"`
	Port        int    `json:"port"`
	HTTPS       bool   `json:"https"`
	AssistToken string `json:"assist_token"`
}

func NewAuth(cli *Client) *Client {
	if cli == nil {
		log.Fatal("rubix-service-rest-client can not be empty")
	}
	var url = cli.URL
	var port = cli.Port
	cli.rest = resty.New()
	if url == "" {
		url = "0.0.0.0"
	}
	if port == 0 {
		port = 1662
	}
	if cli.HTTPS {
		cli.rest.SetBaseURL(fmt.Sprintf("https://%s:%d", url, port))
	} else {
		cli.rest.SetBaseURL(fmt.Sprintf("http://%s:%d", url, port))
	}
	if cli.AssistToken != "" {
		cli.rest.SetHeader("Authorization", setExternalToken(cli.AssistToken))
	}
	return cli
}
