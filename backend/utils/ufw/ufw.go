package ufw

import (
	"github.com/NubeIO/rubix-updater/utils/command"
	"strings"
)

type UFW struct {
	PortsCurrentState map[string]map[string]bool
}

func (ufw *UFW) UFWLoadStatus(asSudo bool) (*UFW, error) {

	ufw.PortsCurrentState = map[string]map[string]bool{}
	cmd := "ufw status | grep ALLOW"
	if asSudo {
		cmd = "sudo ufw status | grep ALLOW"
	}
	output, err := command.RunCMD(cmd, true)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line != "" {
			if strings.Contains(strings.ToLower(line), "reject") == true {
				continue
			}
			for cc := 20; cc > 0; cc-- {
				replace := ""
				for ttt := 0; ttt < cc; ttt++ {
					replace += " "
				}

				line = strings.Replace(line, replace, " ", -1)
			}
			tokens := strings.Split(line, " ")
			address := tokens[2]
			tokens1 := strings.Split(tokens[0], "/")
			protocol := tokens1[0]
			port := tokens1[0]
			if address != "" && protocol != "" && port != "" {
				_, ok := ufw.PortsCurrentState[address]
				if ok == false {
					ufw.PortsCurrentState[address] = map[string]bool{}
				}
				ufw.PortsCurrentState[address][protocol+":"+port] = true
			}
		}
	}

	return ufw, nil
}


