package ffclient

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestFlowClient_GetPlugins(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	var err error
	p, err := cli.GetPlugins()
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(p)

	e, err := cli.GetPlugin("plg_892ec6f5d4044d4a")
	fmt.Println(e, err)
}
