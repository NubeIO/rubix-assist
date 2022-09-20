package edgecli

import (
	"fmt"
	"testing"
)

func TestClient_RestartNetworking(t *testing.T) {
	cli := New(&Client{
		Ip:    deviceIP,
		Port:  0,
		HTTPS: false,
	})

	exists, err := cli.RestartNetworking()
	fmt.Println(exists)
	fmt.Println(err)
	if err != nil {
		return
	}
}
