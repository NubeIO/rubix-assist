package assistcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) EdgeBiosPing(hostIDName string) (*model.Message, error) {
	url := fmt.Sprintf("/api/edge-bios/system/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}
