package appstore

import (
	"fmt"
	"io/ioutil"
)

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
