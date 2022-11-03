package edgebioscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"strconv"
)

func (inst *BiosClient) Login(body *user.User) (*model.TokenResponse, error) {
	url := "/api/users/login"
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
	url := "/api/tokens"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]externaltoken.ExternalToken), nil
}

func (inst *BiosClient) BlockToken(uuid string, state bool) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens/%s/block", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&externaltoken.ExternalToken{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*externaltoken.ExternalToken), nil
}

func (inst *BiosClient) RegenerateToken(uuid string) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/api/tokens/%s/regenerate", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&externaltoken.ExternalToken{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*externaltoken.ExternalToken), nil
}

func (inst *BiosClient) DeleteToken(uuid string) (bool, error) {
	url := fmt.Sprintf("/api/tokens/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().Delete(url))
	fmt.Println(resp.String())
	if err != nil {
		return false, err
	}
	out, _ := strconv.ParseBool(resp.String())
	return out, nil
}
