package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
)

// EdgeReplaceConfig replace the config file of a nube app
func (inst *Client) EdgeReplaceConfig(hostIDName string, app *appstore.EdgeReplaceConfig) (*appstore.EdgeReplaceConfigResp, error) {
	url := fmt.Sprintf("/api/edge/config")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&appstore.EdgeReplaceConfigResp{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.EdgeReplaceConfigResp), nil
}

func (inst *Client) EdgeReadConfig(hostIDName, appName, configName string) (*appstore.EdgeConfig, error) {
	url := fmt.Sprintf("/api/edge/config?name=%s&config=%s", appName, configName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&appstore.EdgeConfig{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*appstore.EdgeConfig), nil
}
