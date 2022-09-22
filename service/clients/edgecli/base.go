package edgecli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Rest  *resty.Client
	Ip    string `json:"ip"`
	Port  int    `json:"port"`
	HTTPS bool   `json:"https"`
}

// New returns a new instance of the nube common apis
func New(cli *Client) *Client {
	if cli == nil {
		log.Fatal("rubix-service-rest-client can not be empty")
		return nil
	}
	var ip = cli.Ip
	var port = cli.Port
	cli.Rest = resty.New()
	if ip == "" {
		ip = "0.0.0.0"
	}
	if port == 0 {
		port = 1661
	}
	if cli.HTTPS {
		cli.Rest.SetBaseURL(fmt.Sprintf("https://%s:%d", ip, port))
	} else {
		cli.Rest.SetBaseURL(fmt.Sprintf("http://%s:%d", ip, port))
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
