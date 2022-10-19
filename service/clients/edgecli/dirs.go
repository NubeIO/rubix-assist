package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/assistcli/nresty"
	"strconv"
)

func (inst *Client) CreateDir(path string) (*model.Message, error) {
	url := fmt.Sprintf("/api/dirs/create?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

// DirExists check if dir exists
func (inst *Client) DirExists(path string) (bool, error) {
	url := fmt.Sprintf("/api/dirs/exists?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Get(url))
	if err != nil {
		return false, err
	}
	found, err := strconv.ParseBool(resp.String())
	if err != nil {
		return false, err
	}
	return found, nil
}

// DeleteDir delete a dir - use the full name of file and path
func (inst *Client) DeleteDir(path string, recursively bool) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/delete?path=%s&recursively=%s", path, "false")
	if recursively {
		url = fmt.Sprintf("/api/files/delete?path=%s&recursively=%s", path, "true")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}
