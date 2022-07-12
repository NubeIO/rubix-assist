package ffclient

import (
	"fmt"
	"testing"
)

func TestFlowClient_NetworkSchema(t *testing.T) {
	cli := NewLocalClient(&Connection{})
	var err error

	schema, err := cli.NetworkSchema("system")
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(schema, err)
}
