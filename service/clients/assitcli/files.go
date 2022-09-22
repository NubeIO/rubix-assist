package assitcli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"io"
	"os"
	"path/filepath"
)

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

// RenameFile rename a file - use the full name of file and path
func (inst *Client) RenameFile(oldNameAndPath, newNameAndPath string) (*Message, error) {
	url := fmt.Sprintf("/api/files/rename?old=%s&new=%s", oldNameAndPath, newNameAndPath)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

// CopyFile copy a file - use the full name of file and path
func (inst *Client) CopyFile(from, to string) (*Message, error) {
	url := fmt.Sprintf("/api/files/copy?from=%s&to=%s", from, to)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

// MoveFile move a file - use the full name of file and path
func (inst *Client) MoveFile(from, to string) (*Message, error) {
	url := fmt.Sprintf("/api/files/move?from=%s&to=%s", from, to)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

// DeleteFile delete a file - use the full name of file and path
func (inst *Client) DeleteFile(path string) (*Message, error) {
	url := fmt.Sprintf("/api/files/delete?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

// DeleteAllFiles delete all file's in a dir - use the full name of file and path
func (inst *Client) DeleteAllFiles(path string) (*Message, error) {
	url := fmt.Sprintf("/api/files/delete/all?path=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

// DeleteDir delete a dir - use the full name of file and path
func (inst *Client) DeleteDir(path string, recursively bool) (*Message, error) {
	url := fmt.Sprintf("/api/files/delete?path=%s&recursively=%s", path, "false")
	if recursively {
		url = fmt.Sprintf("/api/files/delete?path=%s&recursively=%s", path, "true")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

// UploadFile upload file
func (inst *Client) UploadFile(destination, fileName string, file io.Reader) (*EdgeUploadResponse, error) {
	url := fmt.Sprintf("/api/files/upload?destination=%s", destination)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&EdgeUploadResponse{}).
		SetFileReader("file", fileName, file).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*EdgeUploadResponse), nil
}

func (inst *Client) UploadLocalFile(path, fileName, destination string) (*EdgeUploadResponse, error) {
	fileAndPath := filepath.FromSlash(fmt.Sprintf("%s/%s", path, fileName))
	reader, err := os.Open(fileAndPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open file: %s err: %s", fileAndPath, err.Error()))
	}
	resp, err := inst.Rest.R().
		SetResult(&EdgeUploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(fmt.Sprintf("/api/files/upload?destination=%s", destination))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() > 299 {
		return nil, errors.New(resp.String())
	}
	return resp.Result().(*EdgeUploadResponse), nil
}

type DownloadResponse struct {
	FileName    string `json:"file,omitempty"`
	Path        string `json:"path,omitempty"`
	Destination string `json:"destination"`
}

// DownloadFile download a file
func (inst *Client) DownloadFile(path, file, destination string) (*DownloadResponse, error) {
	url := fmt.Sprintf("/api/files/download?path=%s&file=%s", path, file)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetOutput(destination).
		Post(url))
	if err != nil {
		return nil, err
	}
	return &DownloadResponse{FileName: file, Path: path, Destination: destination}, nil
}
