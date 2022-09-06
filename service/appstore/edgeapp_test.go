package appstore

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"testing"
)

func TestStore_generateServiceFile(t *testing.T) {
	var err error
	appName := "flow-framework"
	appVersion := "v0.6.0"
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir: "/data",
		},
		Perm: nonRoot,
	})
	fmt.Println(err)

	exe := []string{"g=/data/bacnet-server-c", "a=/data"}

	tmpDir, absoluteServiceFileName, err := appStore.generateServiceFile(&ServiceFile{
		Name:                    appName,
		Version:                 appVersion,
		ServiceDescription:      "",
		RunAsUser:               "",
		ServiceWorkingDirectory: "",
		AppSpecificExecStart:    "app",
		CustomServiceExecStart:  "",
		EnvironmentVars:         exe,
	})
	fmt.Println(tmpDir, absoluteServiceFileName, err)
	if err != nil {
		return
	}
}
