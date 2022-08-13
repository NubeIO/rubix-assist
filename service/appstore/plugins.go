package appstore

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"io/ioutil"
	"os"
	"strings"
)

//PluginDetails influx-amd64.so
func (inst *Store) PluginDetails(pluginName string) (appName, arch, fileExtension string) {
	parts := strings.Split(pluginName, "-")
	for _, part := range parts {
		appName = parts[0]
		if strings.Contains(part, "amd64") {
			arch = "amd64"
		}
		if strings.Contains(part, "armv7") {
			arch = "armv7"
		}
		if strings.Contains(part, ".so") {
			fileExtension = ".so"
		}
	}
	return appName, arch, fileExtension
}

//CheckBinaryPlugin check if all the details of a binary name is correct (influx-amd64.so)
func (inst *Store) CheckBinaryPlugin(pluginName string) error {
	appName, arch, fileExtension := inst.PluginDetails(pluginName)
	if appName == "" {
		return errors.New(fmt.Sprintf("plugin name is incorrect:%s", pluginName))
	}
	if arch == "" {
		return errors.New(fmt.Sprintf("plugin arch is incorrect:%s", pluginName))
	}
	if fileExtension == "" {
		return errors.New(fmt.Sprintf("plugin fileExtension is incorrect:%s", pluginName))
	}
	return nil
}

type Plugin struct {
	PluginName string `json:"plugin_name"`
	Arch       string `json:"arch"`
	Version    string `json:"version"`
}

//UploadPluginToEdge to edge device
// rubix-ui to pass in a plugin RA and then unzip it to tmp dir and check its arch
// then upload it to edge
// restart FF if as an option
func (inst *Store) UploadPluginToEdge(hostUUID, hostName string, plugin *Plugin) (*EdgeUploadResponse, error) {
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
	zip, err := fileutils.New().UnZip(pluginPathName, tmpDir, os.FileMode(FilePerm))
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
	flowPathPluginPath := fmt.Sprintf("%s/data/plugins", flowPath)
	uploadResp, err := inst.EdgeUploadLocalFile(hostUUID, hostName, tmpDir, binaryName, flowPathPluginPath)
	if err != nil {
		return nil, err
	}
	//err = fileutils.New().RmRF(tmpDir)
	//if err != nil {
	//	return nil, err
	//}
	return uploadResp, nil
}

func (inst *Store) GetPluginPath(plugin *Plugin) (path, zipName string, err error) {
	if plugin == nil {
		return "", "", errors.New("plugin is nil, cant not be empty")
	}
	plugins, err := inst.ListPlugins()
	if err != nil {
		return "", "", err
	}
	var matchName bool
	var matchArch bool
	for _, plg := range plugins {
		if plg.MatchedName == plugin.PluginName {
			matchName = true
			if plg.MatchedArch == plugin.Arch {
				matchArch = true
				if plg.MatchedVersion == plugin.Version {
					return fmt.Sprintf("%s/plugins", inst.App.GetStoreDir()), plg.ZipName, nil
				}
			}
		}
	}
	if !matchName {
		return "", "", errors.New(fmt.Sprintf("failed to find plugin name:%s", plugin.PluginName))
	}
	if !matchArch {
		return "", "", errors.New(fmt.Sprintf("failed to find plugin arch:%s", plugin.Arch))
	}
	return "", "", errors.New(fmt.Sprintf("failed to find plugin version:%s", plugin.Version))
}

func (inst *Store) ListPlugins() ([]installer.BuildDetails, error) {
	pluginStore := fmt.Sprintf("%s/plugins", inst.App.GetStoreDir())
	files, err := ioutil.ReadDir(pluginStore)
	if err != nil {
		return nil, err
	}
	var plugins []installer.BuildDetails
	for _, file := range files {
		plugins = append(plugins, *inst.PluginZipDetails(file.Name()))
	}
	return plugins, err
}

func (inst *Store) PluginZipDetails(pluginName string) *installer.BuildDetails {
	return inst.App.GetZipBuildDetails(pluginName)
}

func (inst *Store) UploadStorePlugin(app *installer.Upload) (*UploadResponse, error) {
	var file = app.File
	uploadResp := &UploadResponse{}
	resp, err := inst.App.Upload(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("upload plugin:%s", err.Error()))
	}
	pluginStore := fmt.Sprintf("%s/plugins", inst.App.GetStoreDir())
	err = inst.App.MakeDirectoryIfNotExists(pluginStore, os.FileMode(FilePerm))
	if err != nil {
		return nil, err
	}
	uploadResp.TmpFile = resp.TmpFile
	source := resp.UploadedFile
	dest := fmt.Sprintf("%s/plugins/%s", inst.App.GetStoreDir(), resp.FileName)
	check := inst.App.FileExists(source)
	if !check {
		return nil, errors.New(fmt.Sprintf("upload file tmp dir not found:%s", source))
	}
	uploadResp.UploadedFile = dest
	err = inst.App.MoveFile(source, dest, true)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("move plugin error:%s", err.Error()))
	}
	uploadResp.UploadedOk = true
	return uploadResp, nil
}
