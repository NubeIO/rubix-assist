package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/model"
	"strings"
)

func (base *Controller) mkDir(host *model.Host, name string, checkIfExists, ifExistsClear bool) (out []byte, err error) {
	opts := commandOpts{
		host: *host,
	}
	dirExists := false
	if checkIfExists {
		check, err := base.dirExists(host, name)
		if err != nil {
			return nil, err
		}
		dirExists = check
		if check && ifExistsClear {
			_, err := base.wipeDir(host, name)
			if err != nil {
				return nil, err
			}
		}
	}
	if !dirExists {
		opts = commandOpts{
			cmd:  fmt.Sprintf("mkdir %s", name),
			host: *host,
		}
		check, _, err := base.runCommand(opts, host.IsLocalhost)
		if err != nil {
			return nil, err
		}
		return check, nil
	}
	return nil, nil
}

func (base *Controller) dirExists(host *model.Host, name string) (out bool, err error) {
	opts := commandOpts{
		cmd:  fmt.Sprintf("if [ -d %s ]; then echo 'yes'; else echo 'no'; fi", name),
		host: *host,
	}
	dirExists := false
	check, _, err := base.runCommand(opts, host.IsLocalhost)
	if err != nil {
		return false, err
	}
	if strings.Contains(string(check), "yes") {

		dirExists = true
	}
	fmt.Println("dirExists()", "result:", string(check))
	return dirExists, nil
}

func (base *Controller) wipeDir(host *model.Host, name string) (out bool, err error) {
	opts := commandOpts{
		cmd:  fmt.Sprintf("rm  %s/*", name),
		host: *host,
	}
	_, _, err = base.runCommand(opts, host.IsLocalhost)
	if err != nil {
		fmt.Println("wipeDir()", "result:", "fail", "command:", opts.cmd)
		return false, nil
	}
	fmt.Println("wipeDir()", "result:", "pass")
	return true, nil

}
