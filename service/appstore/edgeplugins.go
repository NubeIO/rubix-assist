package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"os"
	"path"
)

func (inst *Store) EdgeUploadPlugin(hostUUID, hostName string, plugin *Plugin) (*model.EdgeUploadResponse, error) {
	pluginsStorePluginFile, err := inst.GetPluginsStorePluginFile(plugin)
	if err != nil {
		return nil, err
	}
	tmpDir, err := global.App.MakeTmpDirUpload()
	if err != nil {
		return nil, err
	}
	zip, err := fileutils.UnZip(pluginsStorePluginFile, tmpDir, os.FileMode(global.App.FileMode))
	if err != nil {
		return nil, err
	}
	if len(zip) != 1 {
		return nil, errors.New("the plugin folder contents multiple files")
	}
	binaryName := zip[0]
	err = inst.ValidateBinaryPlugin(binaryName)
	if err != nil {
		return nil, err
	}
	pluginInstallationPath := getPluginInstallationPath()
	file := path.Join(tmpDir, binaryName)
	uploadResp, err := inst.EdgeUploadLocalFile(hostUUID, hostName, file, pluginInstallationPath)
	if err != nil {
		return nil, err
	}
	err = fileutils.RmRF(tmpDir)
	if err != nil {
		return nil, err
	}
	return uploadResp, nil
}

func (inst *Store) EdgeListPlugins(hostUUID, hostName string) ([]Plugin, error) {
	p := getPluginInstallationPath()
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	files, err := client.ListFiles(p)
	if err != nil {
		return nil, err
	}
	var pluginDetails []Plugin
	for _, file := range files {
		pluginDetails = append(pluginDetails, *inst.GetPluginDetails(file))
	}
	return pluginDetails, nil
}

func (inst *Store) EdgeDeletePlugin(hostUUID, hostName string, plugin *Plugin) (*model.Message, error) {
	if plugin == nil {
		return nil, errors.New("plugin is nil, can not be empty")
	}
	if plugin.Name == "" {
		return nil, errors.New("plugin name, can not be empty")
	}
	if plugin.Arch == "" {
		return nil, errors.New("plugin arch, can not be empty")
	}
	p := getPluginInstallationPath()
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	pluginName := fmt.Sprintf("%s-%s.so", plugin.Name, plugin.Arch)
	p = fmt.Sprintf("%s/%s", p, pluginName)
	return client.DeleteFile(p)
}

func (inst *Store) EdgeDeleteAllPlugins(hostUUID, hostName string) (*model.Message, error) {
	p := getPluginInstallationPath()
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteAllFiles(p)
}
