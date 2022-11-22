package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

// ListAppsStatus get all the apps status
func (inst *Client) ListAppsStatus() ([]model.AppsStatus, error) {
	url := fmt.Sprintf("/api/apps/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.AppsStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]model.AppsStatus)
	return *data, nil
}
