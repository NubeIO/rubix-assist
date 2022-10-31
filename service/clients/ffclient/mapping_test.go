package ffclient

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"testing"
)

// MAKE REMOTE NETWORKS

func TestFlowClient_MakeRemoteDevicePoints(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	err := cli.MakeRemoteDevicePoints(&Mapping{
		FlowNetworkUUID:   "fln_a06a1c6606ee493f",
		NetworkUUID:       "net_4896cac168a24bf1",
		NetworkUUIDRemote: "net_f81aa2a182794790",
	})
	//
	fmt.Println(err)

}

// MAKE LOCAL STREAMS

func TestFlowClient_MakeLocalStreams(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	err := cli.MakeLocalStreams(&Mapping{
		FlowNetworkUUID: "fln_a06a1c6606ee493f",
		NetworkUUID:     "net_4896cac168a24bf1",
	})
	fmt.Println(err)
}

func TestFlowClient_AddStreamToExistingFlow(t *testing.T) {
	cli := NewLocalClient(&Connection{})
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
