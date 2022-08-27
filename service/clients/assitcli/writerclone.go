package assitcli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"strconv"
)

// GetWriterClones all objects
func (inst *Client) GetWriterClones(hostIDName string) ([]model.WriterClone, error) {
	url := fmt.Sprintf("proxy/ff/api/producers/writer_clones")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&[]model.WriterClone{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.WriterClone
	out = *resp.Result().(*[]model.WriterClone)
	return out, nil
}

// GetWriterClone an object
func (inst *Client) GetWriterClone(hostIDName, uuid string) (*model.WriterClone, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.WriterClone{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("proxy/ff/api/producers/writer_clones/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.WriterClone), nil
}

// EditWriterClone edit an object
func (inst *Client) EditWriterClone(hostIDName, uuid string, body model.WriterClone, updateProducer bool) (*model.WriterClone, error) {
	param := strconv.FormatBool(updateProducer)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.WriterClone{}).
		SetBody(body).
		SetPathParams(map[string]string{"uuid": uuid}).
		SetQueryParam("update_producer", param).
		Patch("proxy/ff/api/producers/writer_clones/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.WriterClone), nil
}

// CreateWriterClone edit an object
func (inst *Client) CreateWriterClone(hostIDName string, body model.WriterClone) (*model.WriterClone, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetResult(&model.WriterClone{}).
		SetBody(body).
		Post("proxy/ff/api/producers/writer_clones"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.WriterClone), nil
}

// DeleteWriterClone delete
func (inst *Client) DeleteWriterClone(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host_uuid", hostIDName).
		SetHeader("host_name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("proxy/ff/api/producers/writer_clones/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}
