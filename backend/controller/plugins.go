package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TUpdatePlugins struct {
	RemoveOldPlugins string `json:"remove_old_plugins_pass"`
	MakeUploadDir string `json:"make_upload_dir_pass"`
	UpLoadPlugins string `json:"upload_plugins_pass"`
	CleanUp string `json:"clean_up_pass"`

}

func (base *Controller) UpdatePlugins(ctx *gin.Context) {
	body := uploadBody(ctx)
	id := ctx.Params.ByName("id")
	result := new(TUpdatePlugins)
	result.RemoveOldPlugins = "pass"
	result.MakeUploadDir = "pass"
	result.UpLoadPlugins = "pass"
	result.CleanUp = "restarted service"
	_, err := base.clearDir(id, "/data/flow-framework/data/plugins")
	if err != nil {
		fmt.Println("RemoveOldPlugins", err)
		result.RemoveOldPlugins = "failed to remove existing OR there where no existing plugins installed"
	}
	_, err = base.mkDir(id, body.ToPath, false)
	if err != nil {
		fmt.Println("MakeUploadDir", err)
		result.MakeUploadDir = "failed to make OR new dir or was exiting"
	}
	err = base.uploadZip(id, body)
	if err != nil {
	fmt.Println("UpLoadPlugins", err)
		result.UpLoadPlugins = fmt.Sprint(err)
	}
	_, err = base.rmDir(id, body.ToPath, false)
	if err != nil {
		fmt.Println("CleanUp", err)
		result.CleanUp = fmt.Sprint(err)
	}
	_, err = base.runCommand(id, "sudo systemctl restart nubeio-flow-framework.service", true)
	if err != nil {
		fmt.Println("SystemctlRestart", err)
		result.CleanUp = "fail"
	}
	reposeHandler(result, err, ctx)
}

func (base *Controller) UploadPlugins(ctx *gin.Context) {
	body := uploadBody(ctx)
	id := ctx.Params.ByName("id")
	err := base.uploadZip(id, body)
	if err != nil {
		//return
	}
	reposeHandler("string(out)", err, ctx)
}

func (base *Controller) DeleteAllPlugins(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	dir, err := base.clearDir(id, "/data/flow-framework/data/plugins")
	if err != nil {
		reposeHandler(dir, err, ctx)
	} else {
		reposeHandler(dir, err, ctx)
	}
}
