package appstore

import (
	"fmt"
	"io/ioutil"
)

// Making an app
// make the app appstore dirs

type App struct {
	Name    string `json:"name"`    // rubix-wires
	Version string `json:"version"` // v1.1.1
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
	if err := inst.makeAppsStoreDir(); err != nil {
		return nil, err
	}
	if err := inst.makeAppsStoreAppDir(appName); err != nil {
		return nil, err
	}
	if err := inst.makeAppsStoreAppWithVersionDir(appName, version); err != nil {
		return nil, err
	}
	return app, nil
}

type ListApps struct {
	Name     string   `json:"name,omitempty"`
	Path     string   `json:"path,omitempty"`
	Versions []string `json:"versions,omitempty"`
}

func (inst *Store) ListAppsWithVersions() ([]ListApps, error) {
	listApps, err := inst.ListApps()
	var apps []ListApps
	for _, app := range listApps {
		versions, err := inst.ListAppVersions(app.Name)
		if err != nil {
			return nil, err
		}
		app.Versions = versions
	}
	return apps, err
}

func (inst *Store) ListApps() ([]ListApps, error) {
	var apps []ListApps
	var app ListApps
	files, err := ioutil.ReadDir(inst.getAppsStorePath())
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		app.Name = file.Name()
		app.Path = fmt.Sprintf(inst.getAppsStoreAppPath(file.Name()))
		apps = append(apps, app)
	}
	return apps, err
}

func (inst *Store) ListAppVersions(appName string) ([]string, error) {
	var versions []string
	files, err := ioutil.ReadDir(inst.getAppsStoreAppPath(appName))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		versions = append(versions, file.Name())
	}
	return versions, err
}
