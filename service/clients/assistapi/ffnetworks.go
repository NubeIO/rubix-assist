package assistapi

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// EdgeAddNetwork an object
func (inst *Client) EdgeAddNetwork(hostIDName string, body *model.Network, restartPlugin bool) (*model.Network, error) {
	url := fmt.Sprintf("proxy/ff/api/networks?restart_plugin=%t", restartPlugin)
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.Network{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}
