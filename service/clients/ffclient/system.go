package ffclient

import "github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"

type Ping struct {
	Health   string `json:"health"`
	Database string `json:"database"`
}

func (inst *FlowClient) Ping() (*Ping, error) {
	resp, err := nresty.FormatRestyResponse(inst.client.R().
		SetResult(&Ping{}).
		Get("/api/system/ping"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Ping), nil
}
