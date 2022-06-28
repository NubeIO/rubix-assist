package ffclient

import (
	"context"
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	mutex       = &sync.RWMutex{}
	flowClients = map[string]*FlowClient{}
)

type FlowClient struct {
	client *resty.Client
}

// The dialTimeout normally catches: when the server is unreachable and returns i/o timeout within 2 seconds.
// Otherwise, the i/o timeout takes 1.3 minutes on default; which is a very long time for waiting.
// It uses the DialTimeout function of the net package which connects to a server address on a named network before
// a specified timeout.
func dialTimeout(_ context.Context, network, addr string) (net.Conn, error) {
	timeout := 2 * time.Second
	return net.DialTimeout(network, addr, timeout)
}

var transport = http.Transport{
	DialContext: dialTimeout,
}

type Connection struct {
	Ip   string
	Port int
}

func NewLocalClient(conn *Connection) *FlowClient {
	mutex.RLock()
	defer mutex.RUnlock()
	ip := conn.Ip
	port := conn.Port
	if ip == "" {
		ip = "0.0.0.0"
	}
	if port == 0 {
		port = 1660
	}

	url := fmt.Sprintf("%s://%s:%d", getSchema(port), ip, port)
	//if flowClient, found := flowClients[url]; found {
	//	flowClient.client.SetHeader("Authorization", auth.GetInternalToken(true))
	//	return flowClient
	//}
	client := resty.New()
	client.SetDebug(false)
	client.SetBaseURL(url)
	client.SetError(&nresty.Error{})
	client.SetTransport(&transport)
	//client.SetHeader("Authorization", auth.GetInternalToken(true))
	flowClient := &FlowClient{client: client}
	flowClients[url] = flowClient
	return flowClient
}

func newSessionWithToken(ip string, port int, token string, isTokenAuth bool) *FlowClient {
	mutex.RLock()
	defer mutex.RUnlock()
	url := fmt.Sprintf("%s://%s:%d/ff", getSchema(port), ip, port)
	if isTokenAuth {
		url = fmt.Sprintf("%s://%s:%d", getSchema(port), ip, port)
		token = fmt.Sprintf("External %s", token)
	}
	if flowClient, found := flowClients[url]; found {
		flowClient.client.SetHeader("Authorization", token)
		return flowClient
	}
	client := resty.New()
	client.SetDebug(false)
	client.SetBaseURL(url)
	client.SetError(&nresty.Error{})
	client.SetHeader("Authorization", token)
	client.SetTransport(&transport)
	flowClient := &FlowClient{client: client}
	flowClients[url] = flowClient
	return flowClient
}

func getSchema(port int) string {
	if port == 443 {
		return "https"
	}
	return "http"
}
