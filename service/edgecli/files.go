package edgecli

import (
	"fmt"
	"io"
)

type UploadResponse struct {
	Message     interface{} `json:"message,omitempty"`
	Destination string      `json:"destination,omitempty"`
	File        string      `json:"file,omitempty"`
	Size        string      `json:"size,omitempty"`
	UploadTime  string      `json:"upload_time,omitempty"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
}

func (inst *Client) UploadFile(fileName, destination string, reader io.Reader) (*UploadResponse, error) {
	resp, err := inst.Rest.R().
		SetResult(&UploadResponse{}).
		SetError(&UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(fmt.Sprintf("/api/files/upload?destination=%s", destination))
	if err != nil {
		return nil, err
	}
	data := &UploadResponse{}
	if resp.IsSuccess() {
		data = resp.Result().(*UploadResponse)
		return data, err
	} else {
		data = resp.Error().(*UploadResponse)
		return data, err
	}
}
