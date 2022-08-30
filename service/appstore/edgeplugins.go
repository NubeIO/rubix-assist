package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	"os"
)

//EdgeUploadPlugin to edge device
// rubix-ui to pass in a plugin RA and then unzip it to tmp dir and check its arch
// then upload it to edge
// restart FF if as an option
func (inst *Store) EdgeUploadPlugin(hostUUID, hostName string, plugin *Plugin) (*EdgeUploadResponse, error) {
	pluginPath, pluginName, err := inst.GetPluginPath(plugin)
	if err != nil {
		return nil, err
	}
	pluginPathName := fmt.Sprintf("%s/%s", pluginPath, pluginName)
	tmpDir, err := inst.App.MakeTmpDirUpload()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	zip, err := fileutils.UnZip(pluginPathName, tmpDir, os.FileMode(FilePerm))
	if err != nil {
		return nil, err
	}
	if len(zip) == 1 {
	} else {
		return nil, errors.New("the plugin folder contents was greater then one")
	}
	binaryName := zip[0]
	err = inst.CheckBinaryPlugin(binaryName)
	if err != nil {
		return nil, err
	}
	flowPath := inst.App.GetAppPath(flowFramework)
	err = fileutils.DirExistsErr(flowPath)
	if err != nil {
		return nil, errors.New("flow-framework has not be installed yet")
	}
	flowPathPluginPath := fmt.Sprintf("%s/data/plugins", flowPath)
	err = fileutils.DirExistsErr(flowPathPluginPath)
	if err != nil {
		err := inst.App.MakeDirectoryIfNotExists(flowPathPluginPath, os.FileMode(FilePerm))
		if err != nil {
			return nil, err
		}
	}
	uploadResp, err := inst.EdgeUploadLocalFile(hostUUID, hostName, tmpDir, binaryName, flowPathPluginPath)
	if err != nil {
		return nil, err
	}
	err = fileutils.RmRF(tmpDir)
	if err != nil {
		return nil, err
	}
	return uploadResp, nil
}

func (inst *Store) EdgeGetPluginPath() (string, error) {
	flowPath := inst.App.GetAppPath(flowFramework)
	err := fileutils.DirExistsErr(flowPath)
	if err != nil {
		return "", errors.New("flow-framework has not be installed yet")
	}
	return fmt.Sprintf("%s/data/plugins", flowPath), nil
}

func (inst *Store) EdgeListPlugins(hostUUID, hostName string) ([]Plugin, error) {
	path, err := inst.EdgeGetPluginPath()
	if err != nil {
		return nil, err
	}
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	files, err := client.ListFiles(path)
	if err != nil {
		return nil, err
	}
	var pluginDetails []Plugin
	for _, file := range files {
		pluginDetails = append(pluginDetails, *inst.PluginDetail(file))
	}
	return pluginDetails, nil

}

func (inst *Store) EdgeDeletePlugin(hostUUID, hostName string, plugin *Plugin) (*edgecli.Message, error) {
	if plugin == nil {
		return nil, errors.New("plugin is nil, cant not be empty")
	}
	if plugin.PluginName == "" {
		return nil, errors.New("plugin name, cant not be empty")
	}
	if plugin.Arch == "" {
		return nil, errors.New("plugin arch, cant not be empty")
	}
	path, err := inst.EdgeGetPluginPath()
	if err != nil {
		return nil, err
	}
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	pluginName := fmt.Sprintf("%s-%s.so", plugin.PluginName, plugin.Arch)
	path = fmt.Sprintf("%s/%s", path, pluginName)
	return client.DeleteFile(path)
}

func (inst *Store) EdgeDeleteAllPlugins(hostUUID, hostName string) (*edgecli.Message, error) {
	path, err := inst.EdgeGetPluginPath()
	if err != nil {
		return nil, err
	}
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return nil, err
	}
	return client.DeleteAllFiles(path)
}
