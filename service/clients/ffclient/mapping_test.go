package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"testing"
)

var cli = NewLocalClient(&Connection{
	Ip:   "0.0.0.0",
	Port: 1660,
})

var mapping = &Mapping{
	EdgeDeviceName:    "bgis-test",
	FlowNetworkUUID:   "fln_2c7d8c775574461b",
	NetworkUUID:       "net_4896cac168a24bf1",
	NetworkUUIDRemote: "net_bbf64ecb6f694f0b",
}

// MAKE REMOTE NETWORKS

func TestFlowClient_MakeRemoteConsumers(t *testing.T) {

	err := cli.MakeRemoteConsumers(mapping)
	//
	fmt.Println(err)

}

func TestFlowClient_MakeRemoteDevicePoints(t *testing.T) {

	err := cli.MakeRemoteDevicePoints(mapping)
	//
	fmt.Println(err)

}

// MAKE LOCAL STREAMS

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
