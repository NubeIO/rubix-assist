package assistcli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/edgebioscli"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"strconv"
	"testing"
)

func TestClient_ListAppsWithVersions(t *testing.T) {
	url := fmt.Sprintf("/api/tokens/%s", "tok_41cd4496aa05")
	inst := edgebioscli.New(&edgebioscli.BiosClient{Ip: "test.nube-iiot.com", Port: 1659, ExternalToken: "$2a$10$Ws1FXt1p8SREKX16Wck8w.zJ6cBc1uKA65oAmcMYrHef0IUjQZ.jq"})
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().Delete(url))
	out, _ := strconv.ParseBool(resp.String())
	fmt.Println("out", out)
	fmt.Println("RESP>>>", resp.String())
	fmt.Println("err", err)

}
