package openvpncli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	log "github.com/sirupsen/logrus"
)

func (inst *OpenVPNClient) GetOpenVPNConfig(name string) (*amodel.OpenVPNConfig, error) {
	getURL := fmt.Sprintf("/api/openvpn/%s", name)
	resp, connectionErr, responseErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetResult(&amodel.OpenVPNConfig{}).
		Get(getURL))
	if connectionErr != nil {
		return nil, connectionErr
	}
	if responseErr != nil {
		log.Info(fmt.Sprintf("OpenVPN is not found for %s, so generating for it", name))
		postURL := "/api/openvpn"
		_, err := nresty.FormatRestyResponse(inst.Rest.R().
			SetBody(amodel.OpenVPNBody{Name: name}).
			SetResult(&amodel.Message{}).
			Post(postURL))
		if err != nil {
			return nil, err
		}
		resp, err = nresty.FormatRestyResponse(inst.Rest.R().
			SetResult(&amodel.OpenVPNConfig{}).
			Get(getURL))
		if err != nil {
			return nil, err
		}
		openVPNConfig := resp.Result().(*amodel.OpenVPNConfig)
		return openVPNConfig, nil
	}
	openVPNConfig := resp.Result().(*amodel.OpenVPNConfig)
	return openVPNConfig, nil
}
