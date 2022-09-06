package appstore

import (
	"fmt"
	"os"
)

// getAppStorePathAndVersion get the full app install path and version
func (inst *Store) getAppStorePath(appName string) string {
	path := fmt.Sprintf("%s/apps/%s", inst.App.GetStoreDir(), appName)
	return filePath(path)
}

// getAppStorePathAndVersion get the full app install path and version
func (inst *Store) getAppStorePathAndVersion(appName, version string) string {
	path := fmt.Sprintf("%s/apps/%s/%s", inst.App.GetStoreDir(), appName, version)
	return filePath(path)
}

// MakeAppDir  => /data/appstore/
func (inst *Store) makeStoreDir() error {
	return inst.App.MakeDirectoryIfNotExists(inst.App.GetStoreDir(), os.FileMode(FilePerm))
}

// MakeAppDir  => /data/appstore/apps/
func (inst *Store) makeAppDir() error {
	path := fmt.Sprintf("%s/%s", inst.App.GetStoreDir(), "/apps")
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

// MakeApp  => /data/appstore/apps/flow-framework
func (inst *Store) makeApp(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/apps/%s", inst.App.GetStoreDir(), appName)
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

// MakeAppVersionDir  => /data/appstore/apps/flow-framework/v1.1.1
func (inst *Store) makeAppVersionDir(appName, version string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkVersion(version); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/apps/%s/%s", inst.App.GetStoreDir(), appName, version)
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

func (inst *Store) getServiceWorkingDir(appName, appVersion string) string {
	return inst.App.GetAppInstallPathAndVersion(appName, appVersion)
}
