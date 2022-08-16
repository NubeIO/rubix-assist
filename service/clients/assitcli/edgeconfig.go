package assitcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// EdgeReplaceConfig replace the config file of a nube app
func (inst *Client) EdgeReplaceConfig(hostIDName string, app *appstore.EdgeReplaceConfig) (*appstore.EdgeReplaceConfigResp, error) {
	url := fmt.Sprintf("/api/edge/apps/add")
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
