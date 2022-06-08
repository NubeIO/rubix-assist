package client

import (
	"fmt"
	"testing"
	"time"
)

func TestHostNetwork(*testing.T) {

	client := New("0.0.0.0", 8080)

	hosts, _ := client.GetHostNetworks()
	fmt.Println(222, hosts)
	uuid := ""
	fmt.Println(hosts)
	for _, host := range hosts {
		uuid = host.UUID
	}
	if uuid == "" {
		return
	}

	host, res := client.GetHostNetwork(uuid)
	fmt.Println(res.StatusCode)
	if res.StatusCode != 200 {
		//return
	}
	fmt.Println(host)
	host.Name = fmt.Sprintf("name_%d", time.Now().Unix())
	host, res = client.AddHostNetwork(host)
	host.Name = "get fucked_" + fmt.Sprintf("name_%d", time.Now().Unix())
	if res.GetStatus() != 200 {
		//return
	}
	fmt.Println("NEW host", host.Name)
	host, res = client.UpdateHostNetwork(host.UUID, host)
	if res.GetStatus() != 200 {
		//return
	}
	fmt.Println(host.Name, host.UUID)
	fmt.Println(res.GetStatus())
	res = client.DeleteHostNetwork(host.UUID)
	fmt.Println(res.Message)
	if res.GetStatus() != 200 {
		//return
	}

}
