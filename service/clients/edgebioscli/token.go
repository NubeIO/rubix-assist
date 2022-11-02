package edgebioscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *BiosClient) Login(body *user.User) (*model.TokenResponse, error) {
	url := fmt.Sprintf("/api/users/login")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.TokenResponse), nil
}

func (inst *BiosClient) Tokens() (*[]externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]externaltoken.ExternalToken), nil
}
