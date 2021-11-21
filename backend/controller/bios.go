package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/pkg/config"
	"github.com/gin-gonic/gin"
)

func (base *Controller) InstallBios(ctx *gin.Context) {
	getConfig := config.GetConfig()
	body := uploadBody(ctx)
	toPath := body.ToPath
	if body.ToPath == "" {
		toPath = getConfig.Path.ToPath
	}
	fmt.Println(toPath)
	id := ctx.Params.ByName("id")
	d, err := base.GetHostDB(id)
	if err != nil {
		reposeHandler(d, err, ctx)
	} else {
		c, _ := base.newClient(id)
		defer c.Close()

		commands := []string{"sudo rm -r /data",
			"sudo rm -r /data/rubix-bios/apps/install"}

		for _, command := range commands {
			out, err := c.Run(command)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(out)
			}
		}

		reposeHandler("string(out)", err, ctx)
	}

}
