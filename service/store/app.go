package store

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Making an app
// make the app store dirs

type App struct {
	Name    string `json:"name"`    // rubix-wires
	Version string `json:"version"` // v1.1.1
}

type AppsLibrary struct {
	Name        string   `json:"name"`    // rubix-wires
	Version     string   `json:"version"` // v1.1.1
	ArchVersion []string `json:"arch_version"`
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

type Builds struct {
	Name    string   `json:"name"`
	Version []string `json:"version,omitempty"`
	Arch    string   `json:"arch"`
	Path    string   `json:"path"`
}

type ListApps struct {
	Name     string   `json:"name,omitempty"`
	Path     string   `json:"path,omitempty"`
	Versions []string `json:"versions,omitempty"`
	Builds   []Builds `json:"builds,omitempty"`
}

//func (inst *Store) ListApps() ([]ListApps, error) {
//	rootDir := inst.App.GetStoreDir()
//	var apps []ListApps
//	var app ListApps
//	err := filepath.Walk(fmt.Sprintf("%s/apps", rootDir),
//		func(path string, info os.FileInfo, err error) error {
//			if err != nil {
//				return err
//			}
//			parts := strings.Split(path, "/")
//			if len(parts) >= 5 {
//				if len(parts) == 5 { // app name
//					app.Name = parts[4]
//					app.Path = path
//					apps = append(apps, app)
//				}
//			}
//			return nil
//		})
//	if err != nil {
//		return nil, err
//	}
//	return apps, nil
//}

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

func (inst *Store) ListAppsVersions(path string) ([]ListApps, error) {
	var apps []ListApps
	var app ListApps
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		app.Name = file.Name()
		app.Path = fmt.Sprintf("%s/apps/%s", path, file.Name())
		apps = append(apps, app)
	}
	return apps, err
}

//func (inst *Store) ListAppsArchTypes(path string) ([]ListApps, error) {
//	raw, err := inst.listAppsArchTypesRaw(path)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, apps := range raw {
//		//apps
//
//	}
//}

func (inst *Store) listAppsBuilds(path string) ([]ListApps, error) {
	var apps []ListApps
	var app ListApps
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

//func (inst *Store) ListAppsVersions(path string) ([]ListApps, error) {
//	//rootDir := inst.App.GetStoreDir()
//
//	files, err := ioutil.ReadDir(path)
//	if err != nil {
//		//log.Fatal(err)
//	}
//	for _, file := range files {
//		fmt.Println(file.Name(), file.IsDir())
//	}
//}

//func (inst *Store) ListStore() ([]string, error) {
//	rootDir := inst.App.GetStoreDir()
//	var files []AppsLibrary
//	app := AppsLibrary{}
//	var archVersion []string
//	err := filepath.Walk(fmt.Sprintf("%s/apps", rootDir),
//		func(path string, info os.FileInfo, err error) error {
//			if err != nil {
//				return err
//			}
//			parts := strings.Split(path, "/")
//			if len(parts) >= 5 {
//				fmt.Println(path)
//				if len(parts) == 5 { // app name
//					fmt.Println(parts[4])
//					app.Name = parts[4]
//				}
//				if len(parts) == 6 { // version
//					app.Version = parts[5]
//				}
//				if len(parts) == 7 { // build
//					details := inst.App.GetZipBuildDetails(parts[6])
//					if details != nil {
//						if details.MatchedArch != "" {
//							archVersion = append(archVersion, details.MatchedArch)
//							app.ArchVersion = archVersion
//						}
//						if details.ZipName != "" {
//							files = append(files, app)
//						}
//					}
//				}
//			}
//			return nil
//		})
//	if err != nil {
//		return nil, err
//	}
//	return files, nil
//}

// list apps
// list version for an app
// list all arch for an app

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
