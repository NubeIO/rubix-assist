package assistcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) EdgeBiosLogin(hostIDName string, externalToken string, body *user.User) (*model.TokenResponse, error) {
	url := "/proxy/eb/api/users/login"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetHeader("Authorization", composeToken(externalToken)).
		SetBody(body).
		SetResult(&model.TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.TokenResponse), nil
}

func (inst *Client) EdgeBiosTokens(hostIDName, externalToken, jwtToken string) (*[]externaltoken.ExternalToken, error) {
	url := "/proxy/eb/api/tokens"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetHeader("Authorization", composeToken(externalToken)).
		SetHeader("jwt_token", jwtToken).
		SetResult(&[]externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosToken(hostIDName, externalToken, jwtToken string, uuid string) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/proxy/eb/api/tokens/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetHeader("Authorization", composeToken(externalToken)).
		SetHeader("jwt_token", jwtToken).
		SetResult(&externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenGenerate(hostIDName, externalToken, jwtToken string, name string) (*externaltoken.ExternalToken, error) {
	url := "/proxy/eb/api/tokens/generate"
	body := externaltoken.ExternalToken{Name: name, Blocked: false}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetHeader("Authorization", composeToken(externalToken)).
		SetHeader("jwt_token", jwtToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenBlock(hostIDName, externalToken, jwtToken string, uuid string, block bool) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/proxy/eb/api/tokens/%s/block", uuid)
	body := map[string]bool{"blocked": block}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetHeader("Authorization", composeToken(externalToken)).
		SetHeader("jwt_token", jwtToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenRegenerate(hostIDName, externalToken, jwtToken string, uuid string) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/proxy/eb/api/tokens/%s/regenerate", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetHeader("Authorization", composeToken(externalToken)).
		SetHeader("jwt_token", jwtToken).
		SetResult(&externaltoken.ExternalToken{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenDelete(hostIDName, externalToken, jwtToken string, uuid string) (bool, error) {
	url := fmt.Sprintf("/proxy/eb/api/tokens/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetHeader("Authorization", composeToken(externalToken)).
		SetHeader("jwt_token", jwtToken).
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}
