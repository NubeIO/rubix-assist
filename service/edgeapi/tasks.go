package edgeapi

import "errors"

type TaskType int

// go generate ./...

//go:generate stringer -type=TaskType
const (
	PingHost TaskType = iota
	InstallApp
	RemoveApp
)

func CheckTask(s string) error {
	switch s {
	case PingHost.String():
		return nil
	case InstallApp.String():
		return nil
	case RemoveApp.String():
		return nil
	}
	return errors.New("invalid action type, try installApp or ping")

}
