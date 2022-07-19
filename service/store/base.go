package store

import (
	"errors"
	"github.com/NubeIO/lib-rubix-installer/installer"
)

const nonRoot = 0700
const root = 0777

var FilePerm = root

type Store struct {
	App  *installer.App
	Perm int `json:"file_perm"`
}

func New(store *Store) (*Store, error) {
	if store == nil {
		return nil, errors.New("store can not be empty")
	}
	if store.App == nil {
		return nil, errors.New("app can not be empty")
	}
	if store.Perm == 0 {
		store.Perm = FilePerm
	}
	if store.App.FilePerm == 0 {
		store.App.FilePerm = FilePerm
	}
	if store.App.DataDir == "" {
		store.App.DataDir = "/data"
	}
	store.App = installer.New(store.App)
	return store, nil
}
