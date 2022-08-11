package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
)

type UploadResponse struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	UploadedOk   bool   `json:"uploaded_ok"`
	TmpFile      string `json:"tmp_file"`
	UploadedFile string `json:"uploaded_file"`
}

func (inst *Store) AddUploadStoreApp(app *installer.Upload) (*UploadResponse, error) {
	var archType = app.Arch
	var productType = app.Product
	if app.Name == "" {
		return nil, errors.New("app name can not be empty")
	}
	if app.Version == "" {
		return nil, errors.New("app name can not be empty")
	}
	if archType == "" {
		return nil, errors.New("arch type can not be empty, try armv7 amd64")
	}
	if productType == "" {
		return nil, errors.New("product type can not be empty, try RubixCompute, RubixComputeIO, RubixCompute5, Server, Edge28, Nuc")
	}
	_, err := inst.AddApp(&App{
		Name:    app.Name,
		Version: app.Version,
	})
	if err != nil {
		return nil, err
	}
	var appName = app.Name
	var version = app.Version
	var file = app.File
	uploadResp := &UploadResponse{
		Name:         appName,
		Version:      version,
		UploadedOk:   false,
		TmpFile:      "",
		UploadedFile: "",
	}
	resp, err := inst.App.Upload(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("upload app:%s", err.Error()))
	}
	uploadResp.TmpFile = resp.TmpFile
	source := resp.UploadedFile
	dest := fmt.Sprintf("%s/apps/%s/%s/%s", inst.App.GetStoreDir(), appName, version, resp.FileName)
	check := inst.App.FileExists(source)
	if !check {
		return nil, errors.New(fmt.Sprintf("upload file tmp dir not found:%s", source))
	}
	uploadResp.UploadedFile = dest
	err = inst.App.MoveFile(source, dest, true)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("move build error:%s", err.Error()))
	}
	uploadResp.UploadedOk = true
	return uploadResp, nil
}
