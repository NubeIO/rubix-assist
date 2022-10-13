package appstore

import (
	"errors"
	"github.com/NubeIO/lib-rubix-installer/installer"
	base "github.com/NubeIO/rubix-assist/database"
	"os"
)

const flowFramework = "flow-framework"

type Store struct {
	App *installer.App
	DB  *base.DB
}

func New(store *Store) (*Store, error) {
	if store == nil {
		return nil, errors.New("appstore can not be empty")
	}
	if store.App == nil {
		return nil, errors.New("app can not be empty")
	}
	store.App = installer.New(store.App)
	err := store.initMakeAllDirs()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (inst *Store) initMakeAllDirs() error {
	if err := os.MkdirAll(inst.getAppsStorePath(), os.FileMode(inst.App.FileMode)); err != nil {
		return err
	}
	if err := os.MkdirAll(inst.getPluginsStorePath(), os.FileMode(inst.App.FileMode)); err != nil {
		return err
	}
	return nil
}
