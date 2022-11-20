package edgecli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-assist/installer"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/constants"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"os"
	"path"
	"path/filepath"
)

func (inst *Client) PluginUpload(body *installer.Plugin) (*model.Message, error) {
	uploadLocation := global.Installer.GetAppPluginDownloadPath(constants.FlowFramework)
	if body.ClearBeforeUploading {
		url := fmt.Sprintf("/api/files/delete-all?path=%s", uploadLocation)
		_, _ = nresty.FormatRestyResponse(inst.Rest.R().Delete(url))
	}

	url := fmt.Sprintf("/api/dirs/create?path=%s", uploadLocation)
	_, _ = nresty.FormatRestyResponse(inst.Rest.R().Post(url))

	pluginFile, err := global.Installer.GetPluginsStorePluginFile(&installer.Plugin{
		Name:      body.Name,
		Arch:      body.Arch,
		Version:   body.Version,
		Extension: body.Extension,
	})
	if err != nil {
		return nil, err
	}
	tmpDir, err := global.Installer.MakeTmpDirUpload()
	if err != nil {
		return nil, err
	}
	fileDetails, err := fileutils.Unzip(pluginFile, tmpDir, os.FileMode(global.Installer.FileMode))
	if err != nil {
		return nil, err
	}
	if len(fileDetails) != 1 {
		return nil, errors.New(fmt.Sprintf("plugins extraction count mismatch %d", len(fileDetails)))
	}
	extractedPluginFile := path.Join(tmpDir, fileDetails[0].Name)
	reader, err := os.Open(extractedPluginFile)
	if err != nil {
		return nil, err
	}
	url = fmt.Sprintf("/api/files/upload?destination=%s", uploadLocation)
	_, err = nresty.FormatRestyResponse(inst.Rest.R().
		SetFileReader("file", filepath.Base(extractedPluginFile), reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	if err = fileutils.RmRF(tmpDir); err != nil {
		return nil, err
	}
	return &model.Message{Message: "successfully uploaded the plugin"}, nil
}
