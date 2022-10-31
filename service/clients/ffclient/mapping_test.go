package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"testing"
)

var cli = NewLocalClient(&Connection{
	Ip:   "123.209.193.75",
	Port: 1660,
})

var mapping = &Mapping{
	EdgeDeviceName:    "RAAF Richmond 3",
	FlowNetworkUUID:   "fln_b43baddc03bf4590",
	NetworkUUID:       "net_f8f05815339b4ff9",
	NetworkUUIDRemote: "net_aaf0666fdbe5429b",
}

// MAKE REMOTE NETWORKS

func TestFlowClient_MakeRemoteConsumers(t *testing.T) {
	err := cli.MakeRemoteConsumers(mapping)
	fmt.Println(err)

}

// make all the remote points
func TestFlowClient_MakeRemoteDevicePoints(t *testing.T) {
	err := cli.MakeRemoteDevicePoints(mapping)
	fmt.Println(err)

}

// MAKE LOCAL STREAMS
// make all the streams
func TestFlowClient_MakeLocalStreams(t *testing.T) {
	err := cli.MakeLocalStreams(mapping)
	fmt.Println(err)
}

func TestFlowClient_AddStreamToExistingFlow(t *testing.T) {
	err, _ := cli.AddStreamToExistingFlow("fln_a06a1c6606ee493f", &model.Stream{
		CommonStream: model.CommonStream{
			CommonName: model.CommonName{
				Name: "test",
			},
		},
	}, false, Remote{
		FlowNetworkUUID: "fln_a06a1c6606ee493f",
	})
	fmt.Println(err)
}
