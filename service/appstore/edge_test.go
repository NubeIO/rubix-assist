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
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	fmt.Println(err)

	exe := []string{"g=/data/bacnet-server-c", "a=/data"}

	file, s, s2, err := appStore.generateServiceFile(&ServiceFile{
		Name:                    appName,
		Version:                 appVersion,
		ServiceDependency:       "",
		ServiceDescription:      "",
		RunAsUser:               "",
		ServiceWorkingDirectory: "",
		AppSpecficExecStart:     "app",
		CustomServiceExecStart:  "",
		EnvironmentVars:         exe,
	})
	fmt.Println(file, s, s2)
	if err != nil {
		return
	}
}
