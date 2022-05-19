package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"github.com/NubeIO/rubix-assist/model"
)

type commandOpts struct {
	uuid        string
	cmd         string
	sudo, debug bool
	host        model.Host
}

func (inst *Controller) runCommand(commandOpts commandOpts, remoteCommand bool) (out []byte, result bool, err error) {
	if !remoteCommand {
		fmt.Println("HOST:", commandOpts.host.IP, "COMMAND", commandOpts.cmd)
		c, err := inst.newRemoteClient(commandOpts.host)
		if err != nil {
			fmt.Println("REMOTE-COMMAND-ERROR", err)
			return nil, false, err
		}
		defer c.Close()
		out, err = c.Run(commandOpts.cmd)
		if err != nil {
			fmt.Println("REMOTE-COMMAND-ERROR", err)
			return nil, false, err
		}
		return out, true, err
	} else {
		out, err := command.RunCMD(commandOpts.cmd, commandOpts.debug)
		if err != nil {
			return nil, false, err
		}
		return out, false, err
	}
}
