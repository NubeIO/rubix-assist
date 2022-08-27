package edgecli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

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

func (inst *Client) GetUser() (*user.User, error) {
	url := fmt.Sprintf("/api/users")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&user.User{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*user.User), nil
}

func (inst *Client) Login(body *user.User) (*TokenResponse, error) {
	url := fmt.Sprintf("/api/users/login")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*TokenResponse), nil
}

func (inst *Client) GenerateToken(jtwToken string, body *TokenCreate) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens/generate")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("Authorization", jtwToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*externaltoken.ExternalToken), nil
}

func (inst *Client) GetTokens() ([]externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]externaltoken.ExternalToken)
	return *data, nil
}
