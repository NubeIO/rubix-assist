package assistcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

// EdgeWriteConfig replace the config file of a nube app
func (inst *Client) EdgeWriteConfig(hostIDName string, app *assistmodel.EdgeConfig) (*Message, error) {
	url := fmt.Sprintf("/api/edge/config")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&Message{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeReadConfig(hostIDName, appName, configName string) (*assistmodel.EdgeConfigResponse, error, error) {
	url := fmt.Sprintf("/api/edge/config?app_name=%s&config_name=%s", appName, configName)
	resp, connectionError, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&assistmodel.EdgeConfigResponse{}).
		Get(url))
	if connectionError != nil || requestErr != nil {
		return nil, connectionError, requestErr
	}
	return resp.Result().(*assistmodel.EdgeConfigResponse), nil, nil
}
