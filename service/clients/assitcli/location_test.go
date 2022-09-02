package assitcli

import (
	"fmt"
	"testing"
	"time"
)

func TestHostLocation(*testing.T) {
	hosts, res := client.GetHostSchema()
	fmt.Println(hosts)
	uuid := ""
	fmt.Println(hosts)
	fmt.Println(res.StatusCode)
	host, res := client.GetLocation(uuid)
	fmt.Println(res.StatusCode)
	if res.StatusCode != 200 {
		// return
	}
	fmt.Println(host)
	host.Name = fmt.Sprintf("name_%d", time.Now().Unix())
	host, res = client.AddLocation(host)
	host.Name = "get fucked_" + fmt.Sprintf("name_%d", time.Now().Unix())
	fmt.Println(res.StatusCode)
	if res.StatusCode != 200 {

	}

	fmt.Println("NEW host", host.Name)
	host, res = client.UpdateLocation(host.UUID, host)
	if res.StatusCode != 200 {
		// return
	}
}
