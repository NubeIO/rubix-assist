package assitcli

import (
	"fmt"
	"testing"
	"time"
)

func TestHost(*testing.T) {

	client := New("0.0.0.0", 1662)

	hosts, _ := client.GetHosts()
	uuid := ""
	for _, host := range hosts {
		uuid = host.UUID
	}
	if uuid == "" {
		return
	}

	host, res := client.GetHost(uuid)
	fmt.Println(res.StatusCode)
	if res.StatusCode != 200 {
		//return
	}
	fmt.Println(host)
	host.Name = fmt.Sprintf("name_%d", time.Now().Unix())
	host, res = client.AddHost(host)
	host.Name = "get fucked_" + fmt.Sprintf("name_%d", time.Now().Unix())
	if res.StatusCode != 200 {
		return
	}
	fmt.Println("NEW host", host.Name)
	host, res = client.UpdateHost(host.UUID, host)
	if res.StatusCode != 200 {
		return
	}
	fmt.Println(host.Name, host.UUID)
	fmt.Println(res.StatusCode)
	res = client.DeleteHost(host.UUID)
	fmt.Println(res.Message)
	if res.StatusCode != 200 {
		return
	}

}
