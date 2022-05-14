package flow

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/networking"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/portscanner"
)

type Scan struct {
	IP, Iface string
	Debug     bool
}

func (s *Scan) Scan() (hosts *portscanner.Hosts) {
	nets := networking.NewNets()
	interfaceName := s.Iface
	ip := s.IP
	if ip == "" {
		if interfaceName == "" { // if nothing is provided then take a guess of the user network
			gateway, err := nets.GetNetworksThatHaveGateway()
			if err != nil {
				//responseHandler(nil, err, ctx)
				return
			}
			for i, net := range gateway {
				if i == 0 {
					interfaceName = net.Interface
				}
			}
			net, err := nets.GetNetworkByIface(interfaceName)
			if err != nil {
				//responseHandler(nil, err, ctx)
			}
			ip = net.Gateway
		}
	} else {
		ip = "192.168.15"
	}

	ip = fmt.Sprintf("%s-254", ip)
	ports := []string{"22", "1313", "1414", "1616", "1615", "502", "1883"}

	// IP sequence is defined by a '-' between first and last IP address .
	ipsSequence := []string{ip}

	// result returns a map with open ports for each IP address.
	hosts = portscanner.IPScanner(ipsSequence, ports, s.Debug)
	return
}
