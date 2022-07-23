package store

import (
	"fmt"
	"github.com/NubeIO/lib-command/command"
	"github.com/NubeIO/lib-command/unixcmd"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type CheckApp struct {
	Name              string `json:"name"`    // rubix-wires
	Version           string `json:"version"` // v1.1.1
	HasBuild          bool   `json:"has_build"`
	ServiceFileExists bool   `json:"service_file_exists"`
	*MatchBuild
}

func (inst *Store) CheckApp(app *App) (*CheckApp, error) {
	return inst.checkApp(app.Name, app.Version, app.ServiceFile)
}

func (inst *Store) CheckApps(apps []App) ([]CheckApp, error) {
	return inst.checkApps(apps)
}

func (inst *Store) checkApps(apps []App) ([]CheckApp, error) {
	var out []CheckApp
	for _, app := range apps {
		checkApp, err := inst.checkApp(app.Name, app.Version, app.ServiceFile)
		if err != nil {
			//return nil, err
		}
		out = append(out, *checkApp)
	}
	return out, nil

}

func (inst *Store) checkApp(appName, version, serviceName string) (*CheckApp, error) {
	checkApp := &CheckApp{}
	checkApp.Name = appName
	checkApp.Version = version
	check := inst.App.ConfirmStoreAppVersionDir(appName, version)
	checkApp.HasBuild = check
	if !check {
		return checkApp, errors.New(fmt.Sprintf("app store dir not found for app:%s version:%s", appName, version))
	}

	err := inst.serviceFileExists(appName, version, serviceName)
	if err == nil {
		checkApp.ServiceFileExists = true
	}
	matchBuild, err := inst.BuildCheck(appName, version, "")
	checkApp.MatchBuild = matchBuild

	return checkApp, nil
}

type MatchBuild struct {
	MatchName          bool   `json:"match_name"`
	MatchNamePartly    bool   `json:"match_name_partly"`
	MatchVersion       bool   `json:"match_version"`
	MatchVersionPartly bool   `json:"match_version_partly"`
	BuildZipName       string `json:"file_name"`
	Arch               string `json:"arch"`
}

func (inst *Store) BuildCheck(appName, version, path string) (*MatchBuild, error) {
	if path == "" {
		path = fmt.Sprintf("%s/apps/%s/%s", inst.App.GetStoreDir(), appName, version)
	}
	details, err := getFileDetails(path)
	if err != nil {
		return nil, err
	}
	checks := &MatchBuild{}
	for _, detail := range details {
		if detail.Extension == ".zip" {
			fileName := detail.Name

			if strings.Contains(fileName, appName) {
				checks.MatchNamePartly = true
			}
			if strings.Contains(fileName, version) {
				checks.MatchVersionPartly = true
			}
			match, count, versionCheck, archMatch, arch := matchRepoName(fileName, appName)
			if !match {
				errMsg := fmt.Sprintf("failed on match uploaded app, match-count:%d zip file name:%s repo-name:%s arch%s", count, fileName, appName, arch)
				log.Errorln(errMsg)
			} else {
				checks.MatchName = true
			}
			if !archMatch {
				errMsg := fmt.Sprintf("failed on match arch, zip file name:%s repo-name:%s arch%s", fileName, appName, arch)
				log.Errorln(errMsg)
			}
			if version != fmt.Sprintf("v%s", versionCheck) {
				errMsg := fmt.Sprintf("failed on match arch, zip file name:%s repo-name:%s arch%s", fileName, appName, arch)
				log.Errorln(errMsg)
			} else {
				checks.MatchVersionPartly = true
				checks.MatchVersion = true
			}

			checks.BuildZipName = fileName
			checks.Arch = arch
		}
	}
	return checks, nil

}

func (inst *Store) serviceFileExists(appName, version, serviceName string) error {
	path := fmt.Sprintf("%s/apps/%s/%s/%s", inst.App.GetStoreDir(), appName, version, serviceName)
	if !inst.App.FileExists(path) {
		return errors.New(fmt.Sprintf("failed to find service file path:%s", path))
	}
	return nil

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

// matchRepoName get the tag name from the zip eg, wires-builds-0.5.5-1575cf89.amd64.zip => wires-builds
// 	returns
// 	- true if is a match if it is a match
// 	- match count
// 	- string version name
// 	- arch match`
// 	- arch type
func matchRepoName(zipName, repoName string) (bool, int, string, bool, string) {
	parts := strings.Split(zipName, "-")
	repoNameParts := strings.Split(repoName, "-")
	count := 0
	version := ""
	arch := ""
	archMatch := false
	repoMatch := false
	for i, part := range parts {
		p := strings.Split(part, ".")
		// if len is 3 eg, 0.0.1
		isNum := 0
		if len(p) == 3 || len(p) == 4 {
			// check if they are numbers
			for _, s := range p {
				if _, err := strconv.Atoi(s); err == nil {
					isNum++
				}
			}
			if isNum == 3 {
				count = i
				version = part
				version = strings.Trim(version, ".zip")
			}
		}
	}
	match := 0
	for i := 0; i < count; i++ {
		if isMatch(parts, repoNameParts[i]) {
			match++
		}
	}
	if match == count {
		repoMatch = true
	}
	if repoName != "wires-builds" { // wires can run on any os
		arch, _ = getArch()
		if contains(parts, arch) {
			if repoName == "wires-builds" {

			}
			archMatch = true
		}
	} else {
		archMatch = true
	}
	return repoMatch, count, version, archMatch, arch
}

var cmd = unixcmd.New(&command.Command{})

func getArch() (string, error) {
	arch, err := cmd.DetectArch()
	if err != nil {
		return "", err
	}
	return arch.ArchModel, err
}

func isMatch(s []string, term string) bool {
	count := 0
	for _, item := range s {
		if item == term {
			count++
			return true
		}
	}
	return false
}

func contains(s []string, term string) bool {
	count := 0
	for _, item := range s {
		if strings.Contains(item, term) {
			count++
			return true
		}
	}
	return false
}
