package flow

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/portscanner"
)

type Scan struct {
	IP, Iface string
	Debug     bool
}

func (s *Scan) Scan() (hosts portscanner.Hosts) {
	ip := s.IP
	if ip == "" {
		ip = "192.168.15.1-254"
	}

	ports := []string{"22", "1313", "1414", "1616", "1615", "502", "1883"}

	// IP sequence is defined by a '-' between first and last IP address .
	ipsSequence := []string{ip}

	// result returns a map with open ports for each IP address.
	hosts = portscanner.IPScanner(ipsSequence, ports, s.Debug)
	return
}
