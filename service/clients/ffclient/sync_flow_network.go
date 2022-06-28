package ffclient

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

func (a *FlowClient) SyncFlowNetwork(body *model.FlowNetwork) (*model.FlowNetworkClone, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&model.FlowNetworkClone{}).
		SetBody(body).
		Post("/api/sync/flow_network"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FlowNetworkClone), nil
}
