package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"path"
)

type UploadResponse struct {
	Name         string `json:"name,omitempty"`
	Version      string `json:"version,omitempty"`
	UploadedOk   bool   `json:"uploaded_ok,omitempty"`
	TmpFile      string `json:"tmp_file,omitempty"`
	UploadedFile string `json:"uploaded_file,omitempty"`
}

func (inst *Store) UploadAddOnAppStore(app *installer.Upload) (*UploadResponse, error) {
	if app.Name == "" {
		return nil, errors.New("app_name can not be empty")
	}
	if app.Version == "" {
		return nil, errors.New("app_version can not be empty")
	}
	if app.Arch == "" {
		return nil, errors.New("arch_type can not be empty, try armv7 amd64")
	}
	err := inst.makeAppsStoreAppWithVersionDir(app.Name, app.Version)
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
		return nil, errors.New(fmt.Sprintf("upload app: %s", err.Error()))
	}
	uploadResp.TmpFile = resp.TmpFile
	source := resp.UploadedFile
	destination := path.Join(inst.getAppsStoreAppWithVersionPath(appName, version), resp.FileName)
	check := fileutils.FileExists(source)
	if !check {
		return nil, errors.New(fmt.Sprintf("upload file tmp dir not found:%s", source))
	}
	uploadResp.UploadedFile = destination
	err = fileutils.MoveFile(source, destination)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("move build error: %s", err.Error()))
	}
	uploadResp.UploadedOk = true
	return uploadResp, nil
}
