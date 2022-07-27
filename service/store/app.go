package store

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Making an app
// make the app store dirs

type App struct {
	Name        string `json:"name"`         // rubix-wires
	Version     string `json:"version"`      // v1.1.1
	ServiceFile string `json:"service_file"` // nubeio-rubix-wires
}

// AddApp make all the app store dirs
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

func (inst *Store) ListStore() ([]App, error) {
	rootDir := inst.App.GetStoreDir()
	var files []App
	app := App{}
	err := filepath.WalkDir(rootDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && strings.Count(p, string(os.PathSeparator)) == 5 {
			parts := strings.Split(p, "/")
			if len(parts) >= 4 { // app name
				app.Name = parts[4]
			}
			if len(parts) >= 5 { // version
				app.Version = parts[5]
			}
			files = append(files, app)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

//func (inst *Store) GetAppZipName(appName, version string) (fileName string, path string, match *installer.MatchBuild, err error) {
//	path = inst.getAppStorePathAndVersion(appName, version)
//	check, err := inst.App.BuildCheck(appName, version, path)
//	if err != nil {
//		return "", path, nil, err
//	}
//	return filePath(check.BuildZipName), path, check, err
//}

// getAppStorePathAndVersion get the full app install path and version
func (inst *Store) getAppStorePathAndVersion(appName, version string) string {
	path := fmt.Sprintf("%s/apps/%s/%s", inst.App.GetStoreDir(), appName, version)
	return filePath(path)
}

//MakeAppDir  => /data/store/
func (inst *Store) makeStoreDir() error {
	return inst.App.MakeDirectoryIfNotExists(inst.App.GetStoreDir(), os.FileMode(FilePerm))
}

//MakeAppDir  => /data/store/apps/
func (inst *Store) makeAppDir() error {
	path := fmt.Sprintf("%s/%s", inst.App.GetStoreDir(), "/apps")
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

//MakeApp  => /data/store/apps/flow-framework
func (inst *Store) makeApp(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/apps/%s", inst.App.GetStoreDir(), appName)
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(FilePerm))
}

//MakeAppVersionDir  => /data/store/apps/flow-framework/v1.1.1
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
