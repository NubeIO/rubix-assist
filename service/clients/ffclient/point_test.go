package ffclient

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestFlowClient_GetPoints(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	var err error
	nets, err := cli.GetNetworkByPluginName("bacnetmaster", true)
	if err != nil {
		return
	}
	dev := &model.Device{}
	for _, device := range nets.Devices {
		dev = device
		break
	}
	points, err := cli.AddPoint(&model.Point{
		Name:       uuid.ShortUUID("name"),
		DeviceUUID: dev.UUID,
		ObjectType: "analogInput",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(points)
}
