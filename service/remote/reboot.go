package remote

import (
	"github.com/NubeIO/rubix-assist/service/remote/command"
	log "github.com/sirupsen/logrus"
)

func (inst *Admin) HostReboot() (ok bool, err error) {
	cmd := "sudo shutdown -r now"
	inst.SSH.CMD.Commands = command.Builder(cmd)
	inst.SSH.RunCommand()
	if err != nil {
		log.Error("admin: HostReboot Error: ", err)
		return ok, err
	}
	return ok, err
}
