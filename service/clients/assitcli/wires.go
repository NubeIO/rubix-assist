package assitcli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/service/wirescli"
)

func (inst *Client) WiresUpload(hostID string, body interface{}) (data interface{}, response *Response) {
	path := fmt.Sprintf("%s/upload", Paths.Wires.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		SetBody(body).
		Post(path)
	return resp.String(), response.buildResponse(resp, err)
}

func (inst *Client) WiresBackup(hostID string) (data interface{}, err error) {
	path := fmt.Sprintf("%s/backup", Paths.Wires.Path)
	resp, err := inst.Rest.R().
		SetHeader("host_uuid", hostID).
		SetHeader("host_name", hostID).
		Get(path)
	if resp.IsSuccess() {
		r := &wirescli.WiresExport{}
		json.Unmarshal(resp.Body(), r)
		return r.Objects, err
	} else {
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("failed to backup wires: %s", resp.String()))
	}
}
