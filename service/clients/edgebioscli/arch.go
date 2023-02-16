package edgebioscli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *BiosClient) GetArch() (*amodel.Arch, error) {
	url := fmt.Sprintf("/api/system/arch")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&amodel.Arch{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*amodel.Arch), nil
}
