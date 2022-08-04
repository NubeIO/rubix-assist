package main

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli"
	"github.com/NubeIO/rubix-assist/service/store"
	"github.com/pkg/errors"
)

var appName = "flow-framework"
var appVersion = "v0.6.1"
var source = ""
var arch = "amd64"
var product = "Server"

func addUploadApp() error {
	client := assitcli.New("0.0.0.0", 1662)
	listStore, err := client.ListStore()
	fmt.Println(err)
	if err != nil {
		return err
	}

	if len(*listStore) == 0 {

		if err != nil {
			return errors.New("no apps are added")
		}
	}
	pprint.PrintJOSN(listStore)

	app, err := client.AddUploadEdgeApp("rc", &store.EdgeApp{
		Name:    appName,
		Version: appVersion,
		Product: product,
		Arch:    arch,
	})
	if err != nil {
		return err
	}
	pprint.PrintJOSN(app)
	return nil
}

func uploadService() error {
	client := assitcli.New("0.0.0.0", 1662)
	service, err := client.UploadEdgeService("rc", &store.ServiceFile{
		Name:                    appName,
		Version:                 appVersion,
		ServiceDescription:      "",
		RunAsUser:               "",
		ServiceWorkingDirectory: "",
		AppSpecficExecStart:     "app -p 1660 -g /data/flow-framework -d data -prod",
	})
	if err != nil {
		return err
	}
	pprint.PrintJOSN(service)
	source = service.UploadedFile
	return nil
}

func installService() error {
	client := assitcli.New("0.0.0.0", 1662)
	service, err := client.InstallEdgeService("rc", &installer.Install{
		Name:        appName,
		Version:     appVersion,
		ServiceName: "",
		Source:      source,
	})

	if err != nil {
		return err
	}
	pprint.PrintJOSN(service)
	return nil
}

func main() {
	err := addUploadApp()
	fmt.Println("addUploadApp")
	fmt.Println(err)
	if err != nil {
		return
	}
	err = uploadService()
	fmt.Println("uploadService")
	fmt.Println(err)
	if err != nil {
		return
	}
	err = installService()
	fmt.Println("installService")
	fmt.Println(err)
	if err != nil {
		return
	}
}
