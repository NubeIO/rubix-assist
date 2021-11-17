package controller

import (
	"github.com/gin-gonic/gin"
)

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

