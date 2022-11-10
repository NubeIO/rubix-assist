package appstore

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestStore_GenerateUploadEdgeService(t *testing.T) {
	var err error
	appName := "flow-framework"
	appVersion := "v0.6.0"
	appStore, err := New(&Store{})
	fmt.Println(err)
	resp, err := appStore.GenerateServiceFileAndEdgeUpload("", "rc", &ServiceFile{
		Name:                    appName,
		Version:                 appVersion,
		ServiceDescription:      "",
		RunAsUser:               "",
		ServiceWorkingDirectory: "",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.PrintJSON(resp)
}
