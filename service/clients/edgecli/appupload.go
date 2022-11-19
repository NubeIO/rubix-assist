package edgecli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/NubeIO/rubix-assist/service/clients/edgebioscli/ebmodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func (inst *Client) AppUpload(body *model.AppUpload) (*model.Message, error) {
	uploadLocation := global.Installer.GetAppDownloadPath(body.AppName)
	url := fmt.Sprintf("/api/dirs/delete-all?path=%s", uploadLocation)
	_, _ = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))

	url = fmt.Sprintf("/api/dirs/create?path=%s", uploadLocation)
	_, _ = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))

	appStoreFile, err := findAppOnAppStoreFile(body.AppName, body.Arch, body.Version)
	if err != nil {
		return nil, err
	}
	url = fmt.Sprintf("/api/files/upload?destination=%s", uploadLocation)
	reader, err := os.Open(*appStoreFile)
	if err != nil {
		return nil, err
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&ebmodel.UploadResponse{}).
		SetFileReader("file", filepath.Base(body.File), reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	upload := resp.Result().(*ebmodel.UploadResponse)

	url = fmt.Sprintf("/api/zip/unzip?source=%s&destination=%s", upload.Destination, uploadLocation)
	resp, err = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]string{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	unzippedFiles := resp.Result().(*[]string)
	url = fmt.Sprintf("/api/files/delete?file=%s", upload.Destination)
	resp, err = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}

	if body.MoveExtractedFileToNameApp {
		for _, f := range *unzippedFiles {
			from := path.Join(uploadLocation, f)
			to := path.Join(uploadLocation, "app")
			url = fmt.Sprintf("/api/files/move?from=%s&to=%s", from, to)
			resp, err = nresty.FormatRestyResponse(inst.Rest.R().
				SetResult(&model.Message{}).
				Post(url))
			if err != nil {
				return nil, err
			}
		}
	}
	if body.MoveOneLevelInsideFileToOutside {
		url = fmt.Sprintf("/api/files/list-details?path=%s", upload.Destination)

	}
	return nil, nil
}

func findAppOnAppStoreFile(appName, arch, version string) (*string, error) {
	storePath := global.Installer.GetAppsStoreAppPathWithArchVersion(appName, arch, version)
	files, err := ioutil.ReadDir(storePath)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New(fmt.Sprintf("%s store file doesn't exist (arch: %s, version: %s)", appName, arch, version))
	}
	appStoreFile := path.Join(storePath, files[0].Name())
	return &appStoreFile, nil
}
