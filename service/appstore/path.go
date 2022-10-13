package appstore

import (
	"path"
)

// getAppsStorePath => /data/store/apps
func (inst *Store) getAppsStorePath() string {
	p := path.Join(inst.App.StoreDir, "apps")
	return p
}

// getAppsStoreAppPath => /data/store/apps/<app_name>
func (inst *Store) getAppsStoreAppPath(appName string) string {
	p := path.Join(inst.getAppsStorePath(), appName)
	return p
}

// getAppsStoreAppWithArchVersionPath /data/store/apps/<app_name>/<arch>/<version>
func (inst *Store) getAppsStoreAppWithArchVersionPath(appName, arch, version string) string {
	p := path.Join(inst.getAppsStoreAppPath(appName), arch, version)
	return p
}

func (inst *Store) getAppWorkingDir(appName, appVersion string) string {
	return inst.App.GetAppInstallPathWithVersionPath(appName, appVersion)
}

// getPluginsStorePath => /data/store/plugins
func (inst *Store) getPluginsStorePath() string {
	p := path.Join(inst.App.StoreDir, "plugins")
	return p
}

// getPluginsStorePath => /data/store/plugins
func (inst *Store) getPluginsStoreWithFile(fileName string) string {
	p := path.Join(inst.getPluginsStorePath(), fileName)
	return p
}
