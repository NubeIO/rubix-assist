package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (base *Controller) clearDir(uuid, path string) (result bool, err error) {
	c, _ := base.newClient(uuid)
	defer c.Close()
	command := fmt.Sprintf("sudo rm %s/*", path)
	_, err = c.Run(command)
	if err != nil {
		return false, err
	}
	return true, err
}

func (base *Controller) mkDir2(uuid, dir string, sudo bool) (result bool, err error) {
	c, _ := base.newClient(uuid)
	defer c.Close()
	command := fmt.Sprintf("mkdir %s", dir)
	_, err = c.Run(command)
	if err != nil {
		return false, err
	}
	return true, err
}

func (base *Controller) rmDir(uuid, dir string, sudo bool) (result bool, err error) {
	c, _ := base.newClient(uuid)
	defer c.Close()
	command := fmt.Sprintf("rm -r %s", dir)
	_, err = c.Run(command)
	if err != nil {
		return false, err
	}
	return true, err
}

func (base *Controller) ClearDir(ctx *gin.Context) {
	body := dirBody(ctx)
	uuid := ctx.Params.ByName("uuid")
	dir, err := base.clearDir(uuid, body.Path)
	if err != nil {
		reposeHandler(nil, err, ctx)
	} else {
		reposeHandler(dir, err, ctx)
	}
}
