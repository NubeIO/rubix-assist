package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (base *Controller) clearDir(id, path string) (result bool, err error) {
	c, _ := base.newClient(id)
	defer c.Close()
	command := fmt.Sprintf("sudo rm %s/*", path)
	_, err = c.Run(command)
	if err != nil {
		return false, err
	}
	return true, err
}

func (base *Controller) mkDir2(id, dir string, sudo bool) (result bool, err error) {
	c, _ := base.newClient(id)
	defer c.Close()
	command := fmt.Sprintf("mkdir %s", dir)
	_, err = c.Run(command)
	if err != nil {
		return false, err
	}
	return true, err
}

func (base *Controller) rmDir(id, dir string, sudo bool) (result bool, err error) {
	c, _ := base.newClient(id)
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
	id := ctx.Params.ByName("id")
	dir, err := base.clearDir(id, body.Path)
	if err != nil {
		reposeHandler(nil, err, ctx)
	} else {
		reposeHandler(dir, err, ctx)
	}
}
