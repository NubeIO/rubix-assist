package main

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/utils/command"
	"log"
	"strings"

)

type UFW struct {

	IPToPortsCurrentState map[string]map[string]bool

}



func (ufw *UFW) LoadCurrentPolicy() error {
	ufw.IPToPortsCurrentState = map[string]map[string]bool{}

	cmd := "sudo ufw status | grep ALLOW"
	output, err := command.RunCMD(cmd, true)
	fmt.Println(222, err)
	fmt.Println(string(output))

	if err != nil {
		return err
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

			log.Println("Loading rule: " + line)

			tokens := strings.Split(line, " ")
			log.Println("tokens: ", tokens)
			address := tokens[2]
			log.Println("address: " + address)
			tokens1 := strings.Split(tokens[0], "/")
			log.Println("tokens1: ", tokens1)
			protocol := tokens1[0]
			port := tokens1[0]

			fmt.Println(address + " " + port + " " + protocol)

			if address != "" && protocol != "" && port != "" {
				_, ok := ufw.IPToPortsCurrentState[address]
				if ok == false {
					ufw.IPToPortsCurrentState[address] = map[string]bool{}
				}

				ufw.IPToPortsCurrentState[address][protocol+":"+port] = true
			}

		}
	}

	return nil
}

func main()  {

	u := &UFW{

	}
	err := u.LoadCurrentPolicy()
	if err != nil {
		return 
	}
	fmt.Println(999, u.IPToPortsCurrentState)
	
}

//func (ufw *UFW) FlushNewPolicy() error {
//	return FlushNewPolicy(&ufw.IPToPortsCurrentState, &ufw.IPToPortsNewState, ufw.AddCommand, ufw.RemoveCommand)
//}