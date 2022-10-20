package assistcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) GetUsers() (data []model.User, response *Response) {
	path := fmt.Sprintf(Paths.Users.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.User{}).
		Get(path)
	return *resp.Result().(*[]model.User), response.buildResponse(resp, err)
}

func (inst *Client) AddUser(body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf(Paths.Users.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Post(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}

func (inst *Client) UpdateUser(uuid string, body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Users.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Patch(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}

func (inst *Client) DeleteUser(uuid string, body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Users.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Patch(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
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

func (inst *Client) GetUser(jtwToken string) (*user.User, error) {
	url := fmt.Sprintf("/api/users")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("Authorization", jtwToken).
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

func (inst *Client) DeleteToken(jtwToken, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/tokens/%s", uuid)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("Authorization", jtwToken).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return &Message{Message: "deleted ok"}, nil
}

func (inst *Client) GetTokens(jtwToken string) ([]externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("Authorization", jtwToken).
		SetResult(&[]externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]externaltoken.ExternalToken)
	return *data, nil
}
