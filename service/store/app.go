package store

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"os"
)

// Making an app
// make the app store dirs

// ListApps make all the app store dirs
func (inst *Store) ListApps() ([]installer.AppResponse, error) {
	return inst.App.DiscoverStoreInstalled()
}

// AddApp make all the app store dirs
func (inst *Store) AddApp(appName, version string) error {
	if err := inst.App.MakeStoreAll(); err != nil {
		return err
	}
	if err := inst.MakeAppDir(); err != nil {
		return err
	}
	if err := inst.MakeApp(appName); err != nil {
		return err
	}
	if err := inst.MakeAppVersionDir(appName, version); err != nil {
		return err
	}
	return nil
}

//MakeAppDir  => /data/store/apps/
func (inst *Store) MakeAppDir() error {
	path := fmt.Sprintf("%s/%s", inst.App.GetStoreDir(), "/apps")
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

//MakeApp  => /data/store/apps/flow-framework
func (inst *Store) MakeApp(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/apps/%s", inst.App.GetStoreDir(), appName)
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

//MakeAppVersionDir  => /data/store/apps/flow-framework/v1.1.1
func (inst *Store) MakeAppVersionDir(appName, version string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkVersion(version); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/apps/%s/%s", inst.App.GetStoreDir(), appName, version)
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}
