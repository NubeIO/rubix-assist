package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func (inst *Controller) ZipUpload(ctx *gin.Context) {
	//curl -X POST http://localhost:8080/api/tools/zip   -F "file=@/home/aidan/Downloads/file.zip"   -H "Content-Type: multipart/form-data"
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}
	ctx.String(http.StatusOK, "", file.Filename)

}
