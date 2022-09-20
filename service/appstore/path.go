package appstore

import (
	"os"
	"path"
)

// getAppsStorePath => /data/appstore/apps
func (inst *Store) getAppsStorePath() string {
	p := path.Join(inst.App.StoreDir, "apps")
	return p
}

// getAppsStoreAppPath => /data/appstore/apps/<app_name>
func (inst *Store) getAppsStoreAppPath(appName string) string {
	p := path.Join(inst.getAppsStorePath(), appName)
	return p
}

// getAppsStoreAppWithVersionPath get the full app install path and version
func (inst *Store) getAppsStoreAppWithVersionPath(appName, version string) string {
	p := path.Join(inst.getAppsStoreAppPath(appName), version)
	return p
}

// getAppsStoreAppWithVersionPath get the full app install path and version
func (inst *Store) getAppsStoreAppWithVersionAndFile(appName, version, file string) string {
	p := path.Join(inst.getAppsStoreAppPath(appName), version, file)
	return p
}

// MakeAppDir  => /data/appstore
func (inst *Store) makeStoreDir() error {
	return os.MkdirAll(inst.App.StoreDir, os.FileMode(inst.App.FileMode))
}

// MakeAppDir  => /data/appstore/apps
func (inst *Store) makeAppsStoreDir() error {
	p := inst.getAppsStorePath()
	return os.MkdirAll(p, os.FileMode(inst.App.FileMode))
}

// MakeApp  => /data/appstore/apps/flow-framework
func (inst *Store) makeAppsStoreAppDir(appName string) error {
	p := inst.getAppsStoreAppPath(appName)
	return os.MkdirAll(p, os.FileMode(inst.App.FileMode))
}

// MakeAppVersionDir  => /data/appstore/apps/flow-framework/v1.1.1
func (inst *Store) makeAppsStoreAppWithVersionDir(appName, version string) error {
	p := inst.getAppsStoreAppWithVersionPath(appName, version)
	return os.MkdirAll(p, os.FileMode(inst.App.FileMode))
}

func (inst *Store) getAppWorkingDir(appName, appVersion string) string {
	return inst.App.GetAppInstallPathWithVersionPath(appName, appVersion)
}

// getPluginsStorePath => /data/appstore/plugins
func (inst *Store) getPluginsStorePath() string {
	p := path.Join(inst.App.StoreDir, "plugins")
	return p
}

// getPluginsStorePath => /data/appstore/plugins
func (inst *Store) getPluginsStoreWithFile(fileName string) string {
	p := path.Join(inst.getPluginsStorePath(), fileName)
	return p
}
