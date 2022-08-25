package assistapi

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

func (inst *Client) EdgeReadConfig(hostIDName, appName, configName string) (*appstore.EdgeConfig, error) {
	url := fmt.Sprintf("/api/edge/config?name=%s&config=%s", appName, configName)
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&appstore.EdgeConfig{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.EdgeConfig), nil
}
