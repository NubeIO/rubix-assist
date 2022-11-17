package appstore

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/constants"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"path"
)

// getAppsStorePath => /data/store/apps
func (inst *Store) getAppsStorePath() string {
	p := path.Join(global.App.StoreDir, "apps")
	return p
}

// getAppsStoreAppPath => /data/store/apps/<app_name>
func (inst *Store) getAppsStoreAppPath(appName string) string {
	p := path.Join(inst.getAppsStorePath(), appName)
	return p
}

// GetAppsStoreAppWithArchVersionPath /data/store/apps/<app_name>/<arch>/<version>
func (inst *Store) GetAppsStoreAppWithArchVersionPath(appName, arch, version string) string {
	p := path.Join(inst.getAppsStoreAppPath(appName), arch, version)
	return p
}

// getPluginsStorePath => /data/store/plugins
func (inst *Store) getPluginsStorePath() string {
	p := path.Join(global.App.StoreDir, "plugins")
	return p
}

// getPluginsStorePath => /data/store/plugins
func (inst *Store) getPluginsStoreWithFile(fileName string) string {
	p := path.Join(inst.getPluginsStorePath(), fileName)
	return p
}

func getPluginInstallationPath() string {
	flowPath := global.App.GetAppDataPath(constants.FlowFramework)
	return fmt.Sprintf("%s/data/plugins", flowPath)
}
