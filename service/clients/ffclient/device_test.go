package ffclient

import (
	"fmt"
	"testing"
)

func TestFlowClient_GetFirstDevice(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	var err error
	dev, err := cli.GetFirstDevice()
	if err != nil {
		return
	}
	fmt.Println(dev)
}
