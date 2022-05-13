package remote

import (
	"github.com/NubeIO/rubix-assist/service/remote/command"
	"github.com/NubeIO/rubix-assist/service/remote/ssh"
)

type Response struct {
	Ok  bool
	Out string
	Err error
}

type Admin struct {
	SSH *ssh.Host
}

func New(admin *Admin) *Admin {
	opts := &command.Opts{}
	admin.SSH.CMD = opts
	return admin
}

func (inst *Admin) Uptime() (res *command.Response) {
	cmd := "uptime"
	inst.SSH.CMD.Commands = command.Builder(cmd)
	res = inst.SSH.RunCommand()
	return
}
