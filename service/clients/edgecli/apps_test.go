package edgecli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestHost(*testing.T) {

	cli := New("", 0)

	file, err := cli.UploadLocalFile("/home/aidan/rubix/builds", "flow-framework-0.5.6-b1b21422.amd64.zip", "/home/aidan")
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(file)

}
