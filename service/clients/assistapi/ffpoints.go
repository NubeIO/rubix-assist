package assistapi

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/ffclient/nresty"
)

// EdgeGetPoints an objects
func (inst *Client) EdgeGetPoints(hostIDName string) ([]model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.Point{}).
		Get("ff/proxy/api/points/"))
	if err != nil {
		return nil, err
	}
	var out []model.Point
	out = *resp.Result().(*[]model.Point)
	return out, nil
}
