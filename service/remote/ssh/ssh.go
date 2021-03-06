package ssh

import (
	"fmt"
	"strings"

	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/remote/command"
	sh "github.com/helloyi/go-sshclient"
)

type Host struct {
	Host   *model.Host
	CMD    *command.Opts
	IsSudo bool
}

//RunCommand will run a local or remote command, if CommandOpts.Sudo is true then a sudo is added to the existing command (cmd = "sudo " + CommandOpts.CMD)
func (h *Host) RunCommand() (res *command.Response) {
	var err error
	cmd := h.CMD
	res = &command.Response{}
	if nils.BoolIsNil(h.Host.IsLocalhost) {
		res = command.Run(cmd)
		cmdOut := res.Out
		err = res.Err
		if err != nil {
			res.Err = err
			return res
		}
		res.Ok = true
		res.Out = cmdOut
		return res
	} else {
		host := fmt.Sprintf("%s:%d", h.Host.IP, h.Host.Port)
		c, err := sh.DialWithPasswd(host, h.Host.Username, h.Host.Password)
		if err != nil {
			res.Err = err
			return res
		}
		defer c.Close()
		if len(cmd.Commands) <= 0 {
			return res
		}
		out, err := c.Cmd(cmd.Commands[0]).Output()
		if err != nil {
			res.Err = err
			return res
		}
		outAsString := strings.TrimRight(string(out), "\n")
		res.Ok = true
		res.Out = outAsString
		return res
	}
}
