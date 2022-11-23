package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/helpers"
	"github.com/NubeIO/rubix-assist/namings"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	log "github.com/sirupsen/logrus"
)

func (inst *Client) AppsStatus() (*[]amodel.AppsStatus, error) {
	files, err := inst.ListFiles(global.Installer.AppsInstallDir)
	if err != nil {
		return nil, err
	}
	ch := make(chan *amodel.AppsStatus)
	for _, file := range files {
		go inst.getAppStatus(file.Name, ch)
	}
	appsStatus := make([]*amodel.AppsStatus, len(files))
	for i := range appsStatus {
		appsStatus[i] = <-ch
	}
	notNullAppsStatus := make([]amodel.AppsStatus, 0)
	for _, appStatus := range appsStatus {
		if appStatus != nil {
			notNullAppsStatus = append(notNullAppsStatus, *appStatus)
		}
	}
	return &notNullAppsStatus, nil
}

func (inst *Client) getAppStatus(fileName string, ch chan<- *amodel.AppsStatus) {
	appName := namings.GetAppNameFromRepoName(fileName)
	version := inst.getAppVersion(appName)
	if version == nil {
		ch <- nil
	}
	serviceName := namings.GetServiceNameFromAppName(appName)
	state, err := inst.appState(serviceName)
	if err != nil {
		ch <- nil
	}
	appStatus := amodel.AppsStatus{
		Name:        appName,
		Version:     *version,
		ServiceName: serviceName,
		State:       state,
	}
	ch <- &appStatus
}

func (inst *Client) appState(unit string) (*systemctl.SystemState, error) {
	url := fmt.Sprintf("/api/systemctl/state?unit=%s", unit)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&systemctl.SystemState{}).
		Get(url))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return resp.Result().(*systemctl.SystemState), nil
}

func (inst *Client) getAppVersion(appName string) *string {
	file := global.Installer.GetAppInstallPath(appName)
	files, err := inst.ListFiles(file)
	if err != nil {
		return nil
	}
	for _, f := range files {
		if f.IsDir {
			if helpers.CheckVersionBool(f.Name) {
				return &f.Name
			}
		}
	}
	return nil
}
