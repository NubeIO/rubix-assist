package edgecli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"os"
	"strconv"
)

// FileExists check if file exists
func (inst *Client) FileExists(path string) (bool, error) {
	url := fmt.Sprintf("/api/files/exists?path=%s", path)
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

// ListFiles list all files/dirs in a dir
func (inst *Client) ListFiles(path string) ([]string, error) {
	url := fmt.Sprintf("/api/files/list?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]string{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]string), nil
}

// Walk list all files/dirs in a dir
func (inst *Client) Walk(path string) ([]string, error) {
	url := fmt.Sprintf("/api/files/walk?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]string{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]string), nil
}

// ReadFile read a files content
func (inst *Client) ReadFile(path string) ([]byte, error) {
	url := fmt.Sprintf("/api/files/read?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (inst *Client) CreateFile(body *assistmodel.WriteFile) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/create")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

func (inst *Client) WriteString(body *assistmodel.WriteFile) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/write/string")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

func (inst *Client) WriteFileJson(body *assistmodel.WriteFile) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/write/json")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

func (inst *Client) WriteFileYml(body *assistmodel.WriteFile) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/write/yml")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

// RenameFile rename a file - use the full name of file and path
func (inst *Client) RenameFile(oldNameAndPath, newNameAndPath string) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/rename?old=%s&new=%s", oldNameAndPath, newNameAndPath)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

// CopyFile copy a file - use the full name of file and path
func (inst *Client) CopyFile(from, to string) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/copy?from=%s&to=%s", from, to)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

// MoveFile move a file - use the full name of file and path
func (inst *Client) MoveFile(from, to string) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/move?from=%s&to=%s", from, to)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

// DeleteFile delete a file - use the full name of file and path
func (inst *Client) DeleteFile(path string) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/delete?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

// DeleteAllFiles delete all file's in a dir - use the full name of file and path
func (inst *Client) DeleteAllFiles(path string) (*model.Message, error) {
	url := fmt.Sprintf("/api/files/delete/all?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}

func (inst *Client) UploadLocalFile(file, destination string) (*assistmodel.EdgeUploadResponse, error) {
	reader, err := os.Open(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open file: %s err: %s", file, err.Error()))
	}
	resp, err := inst.Rest.R().
		SetResult(&assistmodel.EdgeUploadResponse{}).
		SetFileReader("file", file, reader).
		Post(fmt.Sprintf("/api/files/upload?destination=%s", destination))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() > 299 {
		return nil, errors.New(resp.String())
	}
	return resp.Result().(*assistmodel.EdgeUploadResponse), nil
}

// DownloadFile download a file
func (inst *Client) DownloadFile(path, file, destination string) (*assistmodel.EdgeDownloadResponse, error) {
	url := fmt.Sprintf("/api/files/download?path=%s&file=%s", path, file)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetOutput(destination).
		Post(url))
	if err != nil {
		return nil, err
	}
	return &assistmodel.EdgeDownloadResponse{FileName: file, Path: path, Destination: destination}, nil
}
