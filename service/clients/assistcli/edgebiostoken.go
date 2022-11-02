package assistcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) EdgeBiosLogin(hostIDName string, body *user.User) (*model.TokenResponse, error) {
	url := fmt.Sprintf("/api/edge-bios/users/login")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetBody(body).
		SetResult(&model.TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.TokenResponse), nil
}

func (inst *Client) EdgeBiosTokens(hostIDName, jwtToken string) (*[]externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/edge-bios/tokens")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetQueryParam("jwt_token", jwtToken).
		SetResult(&[]externaltoken.ExternalToken{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]externaltoken.ExternalToken)
	return data, nil
}
