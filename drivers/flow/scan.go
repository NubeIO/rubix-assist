package flow

import (
	"fmt"
	"github.com/NubeIO/lib-networking/scanner"
)

type Scan struct {
	IP, Iface string
	Debug     bool
}

func (s *Scan) Scan(ip string, count int, interfaceName string) {
	if ip == "" {
		ip = "192.168.15.1-254"
	}
	if count == 0 {
		count = 254
	}
	ports := []string{"22", "1414", "1883", "1660", "502", "80", "1313"}

	address, err := scanner.New().ResoleAddress("", count, interfaceName)
	if err != nil {
		fmt.Println("err msg", err)
		return
	}
	scanner.New().IPScanner(address, ports, true)
	return
}
