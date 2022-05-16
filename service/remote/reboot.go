package remote

import (
	"github.com/NubeIO/rubix-assist/service/remote/command"
)

func (inst *Admin) HostReboot() (res *command.Response) {
	cmd := "sudo shutdown -r now"
	inst.SSH.CMD.Commands = command.Builder(cmd)
	res = inst.SSH.RunCommand()
	return
}
