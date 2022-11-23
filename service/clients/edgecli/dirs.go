package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) CreateDir(path string) (*amodel.Message, error) {
	url := fmt.Sprintf("/api/dirs/create?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&amodel.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.Message), nil
}

func (inst *Client) DirExists(path string) (*amodel.DirExistence, error) {
	url := fmt.Sprintf("/api/dirs/exists?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&amodel.DirExistence{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.DirExistence), nil
}
