package edgecli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func Test_ListAppsAndService(*testing.T) {
	cli := New(&Client{})
	apps, err := cli.ListAppsStatus()
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJSON(apps)
}
