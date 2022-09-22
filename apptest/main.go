package main

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/appstore"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli"
)

const (
	appName     = "flow-framework"
	appVersion  = "v0.6.1"
	arch        = "amd64"
	product     = "Server"
)

// TODO: have this on rubix-ui

func addUploadApp() error {
	client := assitcli.New(&assitcli.Client{})
	// listStore, err := client.ListAppsBuildDetails()
	// fmt.Println(err)
	// if err != nil {
	// 	return err
	// }
	//
	// if len(listStore) == 0 {
	// 	return errors.New("no apps are added")
	// }
	// pprint.PrintJSON(listStore)
	app, err := client.EdgeUploadApp("rc", &installer.Upload{
		Name:        appName,
		Version:     appVersion,
		Product:     product,
		Arch:        arch,
	})
	if err != nil {
		return err
	}
	pprint.PrintJSON(app)
	return nil
}

func uploadService() (string, error) {
	client := assitcli.New(&assitcli.Client{})
	service, err := client.EdgeUploadService("rc", &appstore.ServiceFile{
		Name:                    appName,
		Version:                 appVersion,
		ServiceDescription:      "",
		RunAsUser:               "",
		ServiceWorkingDirectory: "",
		AppSpecificExecStart:    "app -p 1660 -g /data/flow-framework -d data -prod",
	})
	if err != nil {
		return "", err
	}
	pprint.PrintJSON(service)
	return service.UploadedFile, nil
}

func installService(source string) error {
	client := assitcli.New(&assitcli.Client{})
	service, err := client.InstallEdgeService("rc", &installer.Install{
		Name:        appName,
		Version:     appVersion,
		Source:      source,
	})

	if err != nil {
		return err
	}
	pprint.PrintJSON(service)
	return nil
}

func main() {
	err := addUploadApp()
	fmt.Println("addUploadApp")
	fmt.Println(err)
	if err != nil {
		return
	}
	source, err := uploadService()
	fmt.Println("uploadService")
	fmt.Println(err)
	if err != nil {
		return
	}
	err = installService(source)
	fmt.Println("installService")
	fmt.Println(err)
	if err != nil {
		return
	}
}
