package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"mime"
	"os"
)

func (inst *Client) CreateSnapshot() ([]byte, string, error) {
	url := fmt.Sprintf("/api/snapshots/create")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Post(url))
	if err != nil {
		return nil, "", err
	}
	_, param, err := mime.ParseMediaType(resp.RawResponse.Header.Get("Content-Disposition"))
	if err != nil {
		return nil, "", err
	}
	return resp.Body(), param["filename"], nil
}

func (inst *Client) RestoreSnapshot(filename string, reader *os.File, useGlobalUUID string) error {
	url := fmt.Sprintf("/api/snapshots/restore")
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetQueryParam("use_global_uuid", useGlobalUUID).
		SetFileReader("file", filename, reader).
		Post(url))
	return err
}
