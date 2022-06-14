package tasks

import "errors"

type TaskType int

// go generate ./...

//go:generate stringer -type=TaskType
const (
	SubTask TaskType = iota
	PingHost
	InstallApp
	InstallPlugin
	RemoveApp
	StopApp
	StartApp
	RestartApp
	EnableApp
	DisableApp
	FileTransfer
	DeleteFile
	Unzip
	CopyDir
	DeleteDir
	RebootHost
)

func CheckTask(s string) error {
	switch s {
	case SubTask.String():
		return nil
	case PingHost.String():
		return nil
	case InstallApp.String():
		return nil
	case InstallPlugin.String():
		return nil
	case RemoveApp.String():
		return nil
	case StopApp.String():
		return nil
	case StartApp.String():
		return nil
	case RestartApp.String():
		return nil
	case EnableApp.String():
		return nil
	case DisableApp.String():
		return nil
	case FileTransfer.String():
		return nil
	case DeleteFile.String():
		return nil
	case Unzip.String():
		return nil
	case CopyDir.String():
		return nil
	case DeleteFile.String():
		return nil
	case DeleteDir.String():
		return nil
	case RebootHost.String():
		return nil
	}
	return errors.New("invalid action type, try InstallApp or PingHost")

}
