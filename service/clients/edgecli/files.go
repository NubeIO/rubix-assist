package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) ListFiles(path string) ([]fileutils.FileDetails, error) {
	url := fmt.Sprintf("/api/files/list?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]fileutils.FileDetails{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]fileutils.FileDetails), nil
}

func (inst *Client) ListFilesV2(path string) ([]fileutils.FileDetails, error, error) {
	url := fmt.Sprintf("/api/files/list?path=%s", path)
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetResult(&[]fileutils.FileDetails{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	return *resp.Result().(*[]fileutils.FileDetails), nil, nil
}

func (inst *Client) DeleteFiles(path string) (*amodel.Message, error, error) {
	url := fmt.Sprintf("/api/files/delete-all?path=%s", path)
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetResult(&amodel.Message{}).
		Delete(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	return resp.Result().(*amodel.Message), nil, nil
}

func (inst *Client) ReadFile(path string) ([]byte, error) {
	url := fmt.Sprintf("/api/files/read?file=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (inst *Client) WriteString(body *amodel.WriteFile) (*amodel.Message, error) {
	url := fmt.Sprintf("/api/files/write/string")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&amodel.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.Message), nil
}

func (inst *Client) WriteFileJson(body *amodel.WriteFile) (*amodel.Message, error) {
	url := fmt.Sprintf("/api/files/write/json")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&amodel.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.Message), nil
}

func (inst *Client) WriteFileYml(body *amodel.WriteFile) (*amodel.Message, error) {
	url := fmt.Sprintf("/api/files/write/yml")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&amodel.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.Message), nil
}
