package edgecli

import (
	"fmt"
	"testing"
)

func TestClient_RestartNetworking(t *testing.T) {
	cli := New(&Client{
		URL:   deviceIP,
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
