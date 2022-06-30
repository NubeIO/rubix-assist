package ffclient

import (
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestFlowClient_GetPoints(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	var err error
	points, err := cli.GetPoint("pnt_dc192335373041a7")
	if err != nil {
		return
	}
	pprint.PrintJOSN(points)
}
