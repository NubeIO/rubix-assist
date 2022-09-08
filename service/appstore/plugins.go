package appstore

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"io/ioutil"
	"os"
	"strings"
)

// PluginDetails influx-amd64.so
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

// PluginDetail takes in the name (influx-amd64.so) and returns the info
func (inst *Store) PluginDetail(pluginName string) *Plugin {
	appName, arch, _ := inst.PluginDetails(pluginName)
	return &Plugin{
		PluginName: appName,
		Arch:       arch,
	}
}

// CheckBinaryPlugin check if all the details of a binary name is correct (influx-amd64.so)
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
	Version    string `json:"version,omitempty"`
}

func (inst *Store) GetPluginPath(plugin *Plugin) (path, zipName string, err error) {
	if plugin == nil {
		return "", "", errors.New("plugin is nil, cant not be empty")
	}
	plugins, err := inst.StoreListPlugins()
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

func (inst *Store) StoreListPlugins() ([]installer.BuildDetails, error) {
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

// PluginZipDetails list the details from the name of the plugin zip, as in name, version and arch
func (inst *Store) PluginZipDetails(pluginName string) *installer.BuildDetails {
	return inst.App.GetZipBuildDetails(pluginName)
}

func (inst *Store) StoreUploadPlugin(app *installer.Upload) (*UploadResponse, error) {
	var file = app.File
	uploadResp := &UploadResponse{}
	resp, err := inst.App.Upload(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("upload plugin:%s", err.Error()))
	}
	pluginStore := fmt.Sprintf("%s/plugins", inst.App.GetStoreDir())
	err = inst.App.MakeDirectoryIfNotExists(pluginStore, os.FileMode(inst.App.FileMode))
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
