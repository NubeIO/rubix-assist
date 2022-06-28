package ffclient

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

func (a *FlowClient) SyncStream(body *model.SyncStream) (*model.StreamClone, error) {
	resp, err := nresty.FormatRestyResponse(a.client.R().
		SetResult(&model.StreamClone{}).
		SetBody(body).
		Post("/api/sync/stream"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.StreamClone), nil
}
