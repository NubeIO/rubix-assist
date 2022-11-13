package edgecli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	mutex   = &sync.RWMutex{}
	clients = map[string]*Client{}
)

type Client struct {
	Rest          *resty.Client
	Ip            string `json:"ip"`
	Port          int    `json:"port"`
	HTTPS         *bool  `json:"https"`
	ExternalToken string `json:"external_token"`
}

func New(cli *Client) *Client {
	mutex.Lock()
	defer mutex.Unlock()

	if cli == nil {
		log.Fatal("edge client cli can not be empty")
		return nil
	}
	baseURL := getBaseUrl(cli)
	if client, found := clients[baseURL]; found {
		return client
	}
	rest := resty.New()
	rest.SetBaseURL(baseURL)
	rest.Header.Set("Authorization", composeToken(cli.ExternalToken))
	cli.Rest = rest
	clients[baseURL] = cli
	return cli
}

func NewForce(cli *Client) *Client {
	mutex.Lock()
	defer mutex.Unlock()

	if cli == nil {
		log.Fatal("edge client cli can not be empty")
		return nil
	}
	baseURL := getBaseUrl(cli)
	rest := resty.New()
	rest.SetBaseURL(baseURL)
	rest.Header.Set("Authorization", composeToken(cli.ExternalToken))
	cli.Rest = rest
	clients[baseURL] = cli
	return cli
}

func getBaseUrl(cli *Client) string {
	cli.Rest = resty.New()
	if cli.Ip == "" {
		cli.Ip = "0.0.0.0"
	}
	if cli.Port == 0 {
		cli.Port = 1661
	}
	var baseURL string
	if cli.HTTPS != nil && *cli.HTTPS {
		baseURL = fmt.Sprintf("https://%s:%d", cli.Ip, cli.Port)
	} else {
		baseURL = fmt.Sprintf("http://%s:%d", cli.Ip, cli.Port)
	}
	return baseURL
}

func composeToken(token string) string {
	return fmt.Sprintf("External %s", token)
}
