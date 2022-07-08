package ffclient

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestFlowClient_BacnetWhoIs(t *testing.T) {

	cli := NewLocalClient(&Connection{})
	var err error
	//devices, err := cli.BacnetWhoIs(&WhoIsOpts{GlobalBroadcast: true}, "net_fcbc53dec7ea4329", false)
	//fmt.Println(err)
	//pprint.PrintJOSN(devices)

	points, err := cli.BacnetDevicePoints("dev_e8f4fc599d16485b", true, true)
	fmt.Println(err)
	pprint.PrintJOSN(points)
}
