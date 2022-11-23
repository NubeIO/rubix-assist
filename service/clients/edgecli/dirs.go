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

// DirExists check if dir exists
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

// DeleteDir delete a dir - use the full name of file and path
func (inst *Client) DeleteDir(path string, recursively bool) (*amodel.Message, error) {
	url := fmt.Sprintf("/api/files/delete?path=%s&recursively=%s", path, "false")
	if recursively {
		url = fmt.Sprintf("/api/files/delete?path=%s&recursively=%s", path, "true")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&amodel.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.Message), nil
}
