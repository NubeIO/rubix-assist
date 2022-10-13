package appstore

import (
	"errors"
	"fmt"
	"strings"
)

type Plugin struct {
	Name      string `json:"name"`
	Arch      string `json:"arch"`
	Version   string `json:"version,omitempty"`
	Extension string `json:"extension"`
}

// GetPluginDetails takes in the name (influx-amd64.so) and returns the info
func (inst *Store) GetPluginDetails(pluginName string) *Plugin {
	parts := strings.Split(pluginName, "-")
	plugin := Plugin{}
	for _, part := range parts {
		plugin.Name = parts[0]
		if strings.Contains(part, "amd64") {
			plugin.Arch = "amd64"
		}
		if strings.Contains(part, "armv7") {
			plugin.Arch = "armv7"
		}
		if strings.Contains(part, ".so") {
			plugin.Extension = ".so"
		}
	}
	return &plugin
}

// ValidateBinaryPlugin check if all the details of a binary name is correct (influx-amd64.so)
func (inst *Store) ValidateBinaryPlugin(pluginName string) error {
	plugin := inst.GetPluginDetails(pluginName)
	if plugin.Name == "" {
		return errors.New(fmt.Sprintf("plugin name is incorrect: %s", pluginName))
	}
	if plugin.Arch == "" {
		return errors.New(fmt.Sprintf("plugin arch is incorrect: %s", pluginName))
	}
	if plugin.Extension == "" {
		return errors.New(fmt.Sprintf("plugin extension is incorrect: %s", pluginName))
	}
	return nil
}

func (inst *Store) GetPluginsStorePluginFile(plugin *Plugin) (pluginsPath string, err error) {
	if plugin == nil {
		return "", errors.New("plugin is nil, can not be empty")
	}
	plugins, err := inst.GetPluginsStorePlugins()
	if err != nil {
		return "", err
	}
	var matchName bool
	var matchArch bool
	for _, plg := range plugins {
		if plg.Name == plugin.Name {
			matchName = true
			if plg.Arch == plugin.Arch {
				matchArch = true
				if plg.Version == plugin.Version {
					return inst.getPluginsStoreWithFile(plg.ZipName), nil
				}
			}
		}
	}
	if !matchName {
		return "", errors.New(fmt.Sprintf("failed to find plugin name: %s", plugin.Name))
	}
	if !matchArch {
		return "", errors.New(fmt.Sprintf("failed to find plugin arch: %s", plugin.Arch))
	}
	return "", errors.New(fmt.Sprintf("failed to find plugin: %s, version: %s", plugin.Name, plugin.Version))
}
