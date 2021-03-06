package store

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestStore_generateServiceFile(t *testing.T) {
	var err error
	appName := "flow-framework"
	appVersion := "v0.6.0"
	appBuildName := "flow-framework"
	appStore, err := New(&Store{
		App: &installer.App{
			DataDir:  "/data",
			FilePerm: nonRoot,
		},
		Perm: nonRoot,
	})
	fmt.Println(err)
	resp, err := appStore.GenerateUploadEdgeService("", "rc", &ServiceFile{
		Name:                    appName,
		Version:                 appVersion,
		BuildName:               appBuildName,
		ServiceDescription:      "",
		RunAsUser:               "",
		ServiceWorkingDirectory: "",
		AppSpecficExecStart:     "app -p 1660 -g /data/flow-framework -d data -prod",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJOSN(resp)

}
