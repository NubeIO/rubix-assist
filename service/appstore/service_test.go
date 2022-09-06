package appstore

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestStore_GenerateUploadEdgeService(t *testing.T) {
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
	resp, err := appStore.GenerateServiceFileAndEdgeUpload("", "rc", &ServiceFile{
		Name:                    appName,
		Version:                 appVersion,
		ServiceDescription:      "",
		RunAsUser:               "",
		ServiceWorkingDirectory: "",
		AppSpecificExecStart:    "app -p 1660 -g /data/flow-framework -d data -prod",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(resp)
}
