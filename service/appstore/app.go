package appstore

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Making an app
// make the app appstore dirs

type App struct {
	Name    string `json:"name"`    // rubix-wires
	Version string `json:"version"` // v1.1.1
}

type AppsLibrary struct {
	Name        string   `json:"name"`    // rubix-wires
	Version     string   `json:"version"` // v1.1.1
	ArchVersion []string `json:"arch_version"`
}

// AddApp make all the app appstore dirs
func (inst *Store) AddApp(app *App) (*App, error) {
	appName := app.Name
	version := app.Version
	if err := inst.App.MakeDataDir(); err != nil {
		return nil, err
	}
	if err := inst.makeStoreDir(); err != nil {
		return nil, err
	}
	if err := inst.makeAppDir(); err != nil {
		return nil, err
	}
	if err := inst.makeApp(appName); err != nil {
		return nil, err
	}
	if err := inst.makeAppVersionDir(appName, version); err != nil {
		return nil, err
	}
	return app, nil
}

type Builds struct {
	Name    string   `json:"name"`
	Version []string `json:"version,omitempty"`
	Arch    string   `json:"arch"`
	Path    string   `json:"path"`
}

type ListApps struct {
	Name     string   `json:"name,omitempty"`
	Path     string   `json:"path,omitempty"`
	Version  string   `json:"version,omitempty"`
	Versions []string `json:"versions,omitempty"`
	Builds   []Builds `json:"builds,omitempty"`
}

func (inst *Store) ListApps() ([]ListApps, error) {
	rootDir := inst.App.GetStoreDir()
	var apps []ListApps
	var app ListApps
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/apps", rootDir))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		app.Name = file.Name()
		app.Path = fmt.Sprintf("%s/apps/%s", rootDir, file.Name())
		apps = append(apps, app)
	}
	return apps, err
}

func (inst *Store) ListAppsWithVersions() ([]ListApps, error) {
	listApps, err := inst.ListApps()
	var apps []ListApps
	for _, app := range listApps {
		versions, err := inst.ListAppVersions(app.Name)
		if err != nil {
			return nil, err
		}
		for _, version := range versions {
			apps = append(apps, version)
		}
	}
	return apps, err
}

func (inst *Store) ListAppsBuildDetails() ([]installer.BuildDetails, error) {
	listApps, err := inst.ListAppsWithVersions()
	var apps []installer.BuildDetails
	for _, app := range listApps {
		versions, err := inst.ListAppArchTypes(app.Name, app.Version)
		if err != nil {
			return nil, err
		}
		for _, version := range versions {
			apps = append(apps, version)
		}
	}
	return apps, err
}

func (inst *Store) ListAppVersions(appName string) ([]ListApps, error) {
	var apps []ListApps
	var app ListApps
	var path = fmt.Sprintf("%s", inst.getAppStorePath(appName))
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		app.Name = appName
		app.Version = file.Name()
		app.Path = fmt.Sprintf("%s/apps/%s", path, file.Name())
		apps = append(apps, app)
	}
	return apps, err
}

func (inst *Store) ListAppArchTypes(appName, version string) ([]installer.BuildDetails, error) {
	var apps []installer.BuildDetails
	files, err := inst.listAppBuilds(appName, version)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		details := inst.App.GetZipBuildDetails(file.Name)
		if details.MatchedName == "" {
			details.MatchedName = appName
		}
		apps = append(apps, *details)
	}
	return apps, err
}

func (inst *Store) listAppBuilds(appName, version string) ([]ListApps, error) {
	var apps []ListApps
	var app ListApps
	var path = fmt.Sprintf("%s", inst.getAppStorePathAndVersion(appName, version))
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		app.Name = file.Name()
		app.Path = fmt.Sprintf("%s/%s", path, file.Name())
		apps = append(apps, app)
	}
	return apps, err
}

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

//MakeAppDir  => /data/appstore/
func (inst *Store) makeStoreDir() error {
	return inst.App.MakeDirectoryIfNotExists(inst.App.GetStoreDir(), os.FileMode(FilePerm))
}

//MakeAppDir  => /data/appstore/apps/
func (inst *Store) makeAppDir() error {
	path := fmt.Sprintf("%s/%s", inst.App.GetStoreDir(), "/apps")
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

//MakeApp  => /data/appstore/apps/flow-framework
func (inst *Store) makeApp(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/apps/%s", inst.App.GetStoreDir(), appName)
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

//MakeAppVersionDir  => /data/appstore/apps/flow-framework/v1.1.1
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

type fileDetails struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
	IsDir     bool   `json:"is_dir"`
}

func getFileDetails(dir string) ([]fileDetails, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var out []fileDetails
	var f fileDetails
	for _, file := range files {
		var extension = filepath.Ext(file.Name())
		f.Extension = extension
		f.Name = file.Name()
		f.IsDir = file.IsDir()
		out = append(out, f)
	}
	return out, nil
}
