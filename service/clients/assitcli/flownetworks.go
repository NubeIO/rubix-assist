package assitcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

func buildUrl(overrideUrl ...string) string {
	if len(overrideUrl) > 0 {
		if overrideUrl[0] != "" {
			return overrideUrl[0]
		}
	}
	return ""
}

// GetFlowNetwork an object
func (inst *Client) GetFlowNetwork(hostIDName, uuid string, withStreams bool, overrideUrl ...string) (*model.FlowNetwork, error) {
	url := fmt.Sprintf("proxy/ff/api/flow_networks/%s", uuid)
	if withStreams == true {
		url = fmt.Sprintf("proxy/ff/api/flow_networks/%s?with_streams=true", uuid)
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.FlowNetwork{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetwork), nil
}

// GetFlowNetworks an objects
func (inst *Client) GetFlowNetworks(hostIDName string, withStreams bool, overrideUrl ...string) ([]model.FlowNetwork, error) {
	url := fmt.Sprintf("proxy/ff/api/flow_networks")
	if withStreams == true {
		url = fmt.Sprintf("proxy/ff/api/flow_networks/?with_streams=true")
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.FlowNetwork{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.FlowNetwork
	out = *resp.Result().(*[]model.FlowNetwork)
	return out, nil
}
