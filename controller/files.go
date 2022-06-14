package controller

import (
	"errors"

	"github.com/NubeIO/rubix-assist/controller/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const dirPath = "/home/aidan/" //TODO add in config

func ConcatPath(localSystemFilePath string) string {
	localSystemFilePath = path.Join(filepath.Dir(dirPath), localSystemFilePath)
	return localSystemFilePath
}

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

func (inst *Controller) GetFile(ctx *gin.Context) {
	localSystemFilePath := ctx.Param("filePath")
	localSystemFilePath = ConcatPath(localSystemFilePath)

	fileInfo, err := os.Stat(localSystemFilePath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			response.WithData(ctx, http.StatusOK, response.FILENOTEXIST, err)
		} else {
			response.WithData(ctx, http.StatusOK, response.ERROR, err)
		}
		return
	}

	var dirContent []string
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(localSystemFilePath)
		if err != nil {
			response.WithData(ctx, http.StatusOK, response.ERROR, err)
			return
		}
		for _, file := range files {
			dirContent = append(dirContent, file.Name())
		}
	} else {
		byteFile, err := ioutil.ReadFile(localSystemFilePath)
		if err != nil {
			response.WithData(ctx, http.StatusOK, response.ERROR, err)
			return
		}

		ctx.Header("Content-Disposition", "attachment; filename=Readme.md")
		ctx.Data(http.StatusOK, "application/octet-stream", byteFile)
	}
	response.WithData(ctx, http.StatusOK, response.SUCCESS, gin.H{"path": dirContent})
}
