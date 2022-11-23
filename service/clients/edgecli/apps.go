package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

// ListAppsStatus get all the apps status
func (inst *Client) ListAppsStatus() ([]amodel.AppsStatus, error) {
	url := fmt.Sprintf("/api/apps/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]amodel.AppsStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]amodel.AppsStatus)
	return *data, nil
}
