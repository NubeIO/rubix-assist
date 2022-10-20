package edgebioscli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	mutex   = &sync.RWMutex{}
	clients = map[string]*BiosClient{}
)

type BiosClient struct {
	Rest          *resty.Client
	Ip            string `json:"ip"`
	Port          int    `json:"port"`
	HTTPS         *bool  `json:"https"`
	ExternalToken string `json:"external_token"`
}

func New(cli *BiosClient) *BiosClient {
	mutex.Lock()
	defer mutex.Unlock()

	if cli == nil {
		log.Fatal("edge bios client cli can not be empty")
		return nil
	}
	cli.Rest = resty.New()
	if cli.Ip == "" {
		cli.Ip = "0.0.0.0"
	}
	if cli.Port == 0 {
		cli.Port = 1659
	}
	var baseURL string
	if cli.HTTPS != nil && *cli.HTTPS {
		baseURL = fmt.Sprintf("https://%s:%d", cli.Ip, cli.Port)
	} else {
		baseURL = fmt.Sprintf("http://%s:%d", cli.Ip, cli.Port)
	}
	if client, found := clients[baseURL]; found {
		return client
	}
	cli.Rest.SetBaseURL(baseURL)
	cli.SetTokenHeader(cli.ExternalToken)
	clients[baseURL] = cli
	return cli
}

func (inst *BiosClient) SetTokenHeader(token string) *BiosClient {
	inst.Rest.Header.Set("Authorization", composeToken(token))
	return inst
}

func composeToken(token string) string {
	return fmt.Sprintf("External %s", token)
}
