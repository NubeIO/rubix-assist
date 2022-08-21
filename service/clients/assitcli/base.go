package assitcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Client struct {
	Rest *resty.Client
}

type FlowClient struct {
	client *resty.Client
}

// New returns a new instance of the nube common apis
func New(url string, port int) *Client {
	rest := &Client{
		Rest: resty.New(),
	}
	rest.Rest.SetBaseURL(fmt.Sprintf("http://%s:%d", url, port))
	return rest
}

type AssistClient struct {
	rest        *resty.Client
	URL         string `json:"url"`
	Port        int    `json:"port"`
	HTTPS       bool   `json:"https"`
	AssistToken string `json:"assist_token"`
}

func setExternalToken(token string) string {
	return fmt.Sprintf("External %s", token)
}

func NewAuth(cli *AssistClient) *AssistClient {
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
		port = 1661
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

func (inst *AssistClient) ProxyGET(hostIDName, path string) (*resty.Response, error) {
	path = fmt.Sprintf("proxy/%s", path)
	path = strings.Replace(path, "//", "/", -1) // makes this /api//host/123 => /api/host/123
	resp, err := inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		Get(path)
	if err != nil {
		return nil, err
	}
	return resp, err
}

type TokenCreate struct {
	Name    string `json:"name" binding:"required"`
	Blocked *bool  `json:"blocked" binding:"required"`
}

type TokenBlock struct {
	Blocked *bool `json:"blocked" binding:"required"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (inst *AssistClient) GetUser(jtwToken string) (*user.User, error) {
	url := fmt.Sprintf("/api/users")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("Authorization", jtwToken).
		SetResult(&user.User{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*user.User), nil
}

func (inst *AssistClient) Login(body *user.User) (*TokenResponse, error) {
	url := fmt.Sprintf("/api/users/login")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetBody(body).
		SetResult(&TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*TokenResponse), nil
}

func (inst *AssistClient) GenerateToken(jtwToken string, body *TokenCreate) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens/generate")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("Authorization", jtwToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*externaltoken.ExternalToken), nil
}

func (inst *AssistClient) DeleteToken(jtwToken, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/tokens/%s", uuid)
	_, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("Authorization", jtwToken).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return &Message{Message: "deleted ok"}, nil
}

func (inst *AssistClient) GetTokens(jtwToken string) ([]externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens")
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("Authorization", jtwToken).
		SetResult(&[]externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]externaltoken.ExternalToken)
	return *data, nil
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
