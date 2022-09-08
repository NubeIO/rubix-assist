package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"strconv"
)

func (inst *Client) CreateDir(path string) (*assistmodel.Message, error) {
	url := fmt.Sprintf("/api/dir/create/?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&assistmodel.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*assistmodel.Message), nil
}

// DirExists check if dir exists
func (inst *Client) DirExists(path string) (bool, error) {
	url := fmt.Sprintf("/api/dirs/exists/?path=%s", path)
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
func (inst *Client) DeleteDir(path string, recursively bool) (*assistmodel.Message, error) {
	url := fmt.Sprintf("/api/files/delete/?path=%s&recursively=%s", path, "false")
	if recursively {
		url = fmt.Sprintf("/api/files/delete/?path=%s&recursively=%s", path, "true")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&assistmodel.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*assistmodel.Message), nil
}
