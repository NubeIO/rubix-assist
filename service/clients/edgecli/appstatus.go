package edgecli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/helpers"
	"github.com/NubeIO/rubix-assist/namings"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
)

func (inst *Client) AppsStatus() (*[]amodel.AppsStatus, error) {
	files, err := inst.ListFiles(global.Installer.AppsInstallDir)
	if err != nil {
		return nil, err
	}
	ch := make(chan *amodel.AppsStatus)
	for _, file := range files {
		appName := namings.GetAppNameFromRepoName(file.Name)
		go inst.appStatusChannel(appName, ch)
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

func (inst *Client) appStatusChannel(appName string, ch chan<- *amodel.AppsStatus) {
	appStatus, _, _ := inst.GetAppStatus(appName)
	ch <- appStatus
}

func (inst *Client) GetAppStatus(appName string) (*amodel.AppsStatus, error, error) {
	version, connectionErr, requestErr := inst.getAppVersion(appName)
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	serviceName := namings.GetServiceNameFromAppName(appName)
	state, connectionErr, requestErr := inst.appState(serviceName)
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	appStatus := amodel.AppsStatus{
		Name:        appName,
		Version:     version,
		ServiceName: serviceName,
		State:       state,
	}
	return &appStatus, nil, nil
}

func (inst *Client) appState(unit string) (*systemctl.SystemState, error, error) {
	url := fmt.Sprintf("/api/systemctl/state?unit=%s", unit)
	res, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetResult(&systemctl.SystemState{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	return res.Result().(*systemctl.SystemState), nil, nil
}

func (inst *Client) getAppVersion(appName string) (string, error, error) {
	file := global.Installer.GetAppInstallPath(appName)
	files, connectionErr, requestErr := inst.ListFilesV2(file)
	if connectionErr != nil || requestErr != nil {
		return "", connectionErr, requestErr
	}
	for _, f := range files {
		if f.IsDir {
			if helpers.CheckVersionBool(f.Name) {
				return f.Name, nil, nil
			}
		}
	}
	return "", nil, errors.New("version can't be nil")
}
