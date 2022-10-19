package edgebioscli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

func (inst *BiosClient) Ping() (*model.Message, error) {
	url := fmt.Sprintf("/api/system/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Message), nil
}
