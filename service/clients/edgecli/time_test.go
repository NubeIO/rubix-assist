package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-edge/service/system"
	"testing"
	"time"
)

func TestClient_GetHardwareTZ(t *testing.T) {
	cli := New(&Client{
		Ip:   deviceIP,
		Port: 0,
	})

	resp, err := cli.GetHardwareTZ()
	fmt.Println(err)
	fmt.Println(resp)
}

func TestClient_GetTimeZoneList(t *testing.T) {
	cli := New(&Client{
		Ip:   deviceIP,
		Port: 0,
	})

	resp, err := cli.GetTimeZoneList()
	fmt.Println(err)
	fmt.Println(resp)
}

func TestClient_SetSystemTime(t *testing.T) {
	cli := New(&Client{
		Ip:   deviceIP,
		Port: 0,
	})
	newTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(newTime)
	resp, err := cli.SetSystemTime(system.DateBody{
		DateTime: newTime,
	})
	fmt.Println(err)
	fmt.Println(resp)
}

func TestClient_SystemTime(t *testing.T) {
	cli := New(&Client{
		Ip:   deviceIP,
		Port: 0,
	})

	resp, err := cli.SystemTime()
	fmt.Println(err)
	fmt.Println(resp)
}

func TestClient_UpdateTimezone(t *testing.T) {
	cli := New(&Client{
		Ip:   deviceIP,
		Port: 0,
	})

	resp, err := cli.UpdateTimezone(system.DateBody{

		TimeZone: "Australia/Sydney",
	})
	fmt.Println(err)
	fmt.Println(resp)
}
