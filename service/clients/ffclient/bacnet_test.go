package ffclient

import (
	"fmt"
	"github.com/NubeDev/bacnet"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestFlowClient_BacnetWhoIs(t *testing.T) {

	cli := NewLocalClient(&Connection{})
	var err error
	devices, err := cli.BacnetWhoIs(&bacnet.WhoIsOpts{}, "net_1dc89f79b04e4874", true)
	fmt.Println(err)
	pprint.PrintJOSN(devices)

	//points, err := cli.BacnetDevicePoints("dev_a99cd05a7d504c66", true)
	//fmt.Println(err)
	//pprint.PrintJOSN(points)
}
