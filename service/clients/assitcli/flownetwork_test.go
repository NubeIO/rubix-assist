package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_AddFlowNetwork(t *testing.T) {
	// fln_cd7d562b1fa04c1e

	cli := New(&Client{
		Rest:        nil,
		URL:         "",
		Port:        0,
		HTTPS:       false,
		AssistToken: "",
	})
	// stream := &model.Stream{
	//	CommonStream: model.CommonStream{
	//		CommonName: model.CommonName{
	//			Name: fmt.Sprintf("%d", time.Now().Unix()),
	//		},
	//	},
	// }
	// flow, err := cli.AddStreamToExistingFlow("rc", "fln_cd7d562b1fa04c1e", stream)
	// fmt.Println(err)
	// if err != nil {
	//	return
	// }
	// pprint.PrintJSON(flow)

	network, err := cli.GetStreamsByFlowNetwork("rc", "fln_cd7d562b1fa04c1e")
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(network)
}
