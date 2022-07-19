package store

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

func (inst *Store) UploadApp(app *installer.Upload) (*UploadResponse, error) {
	if app.Name == "" {
		return nil, errors.New("app name can not be empty")
	}
	if app.Version == "" {
		return nil, errors.New("app name can not be empty")
	}
	check := inst.App.ConfirmStoreDir()
	if !check {
		return nil, errors.New("app store dir not found")
	}
	check = inst.App.ConfirmStoreAppDir(app.Name)
	if !check {
		return nil, errors.New(fmt.Sprintf("app store dir not found for app:%s", app.Name))
	}
	check = inst.App.ConfirmStoreAppVersionDir(app.Name, app.Version)
	if !check {
		return nil, errors.New(fmt.Sprintf("app store dir not found for app:%s version:%s", app.Name, app.Version))
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
	check = inst.App.FileExists(source)
	if !check {
		return nil, errors.New(fmt.Sprintf("upload file tmp dir not found:%s", source))
	}
	uploadResp.UploadedFile = dest
	err = inst.App.MoveFile(source, dest, false)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("move build error:%s", err.Error()))
	}
	uploadResp.UploadedOk = true
	return uploadResp, nil
}
