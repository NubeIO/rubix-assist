package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// GetFlowNetworkClones an object
func (inst *FlowClient) GetFlowNetworkClones(withStreams ...bool) ([]model.FlowNetworkClone, error) {
	url := fmt.Sprintf("/api/flow_network_clones")
	if len(withStreams) > 0 {
		if withStreams[0] == true {
			url = fmt.Sprintf("/api/flow_network_clones/?with_streams=true")
		}
	}
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.FlowNetworkClone{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.FlowNetworkClone
	out = *resp.Result().(*[]model.FlowNetworkClone)
	return out, nil
}

// GetFlowNetworkClone an object
func (inst *FlowClient) GetFlowNetworkClone(uuid string, withStreams ...bool) (*model.FlowNetworkClone, error) {
	url := fmt.Sprintf("/api/flow_network_clones/%s", uuid)
	if len(withStreams) > 0 {
		if withStreams[0] == true {
			url = fmt.Sprintf("/api/flow_network_clones/%s?with_streams=true", uuid)
		}
	}
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.FlowNetworkClone{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetworkClone), nil
}

// GetFlowNetworkClonesWithChild an object
func (inst *FlowClient) GetFlowNetworkClonesWithChild() ([]model.FlowNetworkClone, error) {
	url := fmt.Sprintf("/api/flow_network_clones?with_streams=true&with_producers=true&with_consumers=true&with_writers=true&with_tags=true")
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&[]model.FlowNetworkClone{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.FlowNetworkClone
	out = *resp.Result().(*[]model.FlowNetworkClone)
	return out, nil
}

// GetFlowNetworkCloneWithChild an object
func (inst *FlowClient) GetFlowNetworkCloneWithChild(uuid string) (*model.FlowNetworkClone, error) {
	url := fmt.Sprintf("/api/flow_network_clones/%s?with_streams=true&with_producers=true&with_consumers=true&with_writers=true&with_tags=true", uuid)
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&model.FlowNetworkClone{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetworkClone), nil
}

// DeleteFlowNetworkClone an object
func (inst *FlowClient) DeleteFlowNetworkClone(uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.client.R().
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/api/flow_network_clones/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
