package installer

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"path"
)

const fileMode = 0755
const defaultTimeout = 30

type Installer struct {
	DataDir         string // /data
	StoreDir        string // <data_dir>/store
	TmpDir          string // /data/tmp
	FileMode        int    // 0755
	DefaultTimeout  int    // 30
	AppsDownloadDir string // <data_dir>/rubix-service/apps/download
	AppsInstallDir  string // <data_dir>/rubix-service/apps/install
	SystemCtl       *systemctl.SystemCtl
}

func New(app *Installer) *Installer {
	if app == nil {
		app = &Installer{}
	}
	if app.DataDir == "" {
		app.DataDir = "/data"
	}
	if app.FileMode == 0 {
		app.FileMode = fileMode
	}
	if app.DefaultTimeout == 0 {
		app.DefaultTimeout = defaultTimeout
	}
	if app.StoreDir == "" {
		app.StoreDir = path.Join(app.DataDir, "store")
	}
	if app.TmpDir == "" {
		app.TmpDir = path.Join(app.DataDir, "tmp")
	}
	if app.AppsDownloadDir == "" {
		app.AppsDownloadDir = path.Join(app.DataDir, "rubix-service/apps/download")
	}
	if app.AppsInstallDir == "" {
		app.AppsInstallDir = path.Join(app.DataDir, "rubix-service/apps/install")
	}
	app.SystemCtl = systemctl.New(false, app.DefaultTimeout)
	return app
}
