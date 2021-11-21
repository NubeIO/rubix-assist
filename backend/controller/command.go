package controller

import "fmt"

func (base *Controller) runCommand(id, cmd string, sudo bool) (result bool, err error) {
	c, _ := base.newClient(id)
	defer c.Close()
	command := fmt.Sprintf("%s", cmd)
	_, err = c.Run(command)
	if err != nil {
		return false, err
	}
	return true, err
}