package appstore

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	log "github.com/sirupsen/logrus"
)

type EdgeReplaceConfig struct {
	AppName              string `json:"app_name"`
	FileName             string `json:"file_name"`
	RestartApp           bool   `json:"restart_app"`
	DeleteFileFromAssist bool   `json:"delete_file_from_assist"`
}

type EdgeReplaceConfigResp struct {
	AppName            string              `json:"app_name"`
	EdgeUploadResponse *EdgeUploadResponse `json:"edge_upload_response,omitempty"`
	RestartMessage     string              `json:"restart_message,omitempty"`
}

func (inst *Store) EdgeWriteConfig(hostUUID, hostName string, body *EdgeReplaceConfig) (*EdgeReplaceConfigResp, error) {
	appName := body.AppName
	fileName := body.FileName
	//restartApp := body.RestartApp
	deleteFileFromAssist := body.DeleteFileFromAssist
	if appName == "" {
		return nil, errors.New("assist-config-file: app name, cant not be empty")
	}
	if fileName == "" {
		return nil, errors.New("assist-config-file: file name, cant not be empty, try config.yml, config.json or .env")
	}
	dir := inst.App.StoreDir
	dest := fmt.Sprintf("/data/%s/config", appName)
	fileNamePath := fmt.Sprintf("%s/%s", dir, appName)
	var restartMsg = fmt.Sprintf("restarted app ok %s", appName)
	log.Infof("assist-config-file: try and upload config file to edge file:%s", dest)
	file, err := inst.EdgeUploadLocalFile(hostUUID, hostName, dir, fileName, dest)
	if err != nil {
		return nil, err
	}
	//if restartApp {
	//	_, err := inst.EdgeCtlAction(hostUUID, hostName, &installer.CtlBody{
	//		AppName: appName,
	//		Action:  "restart",
	//	})
	//	if err != nil {
	//		restartMsg = fmt.Sprintf("assist-config-file: failed to restart app:%s err%s", appName, err.Error)
	//		log.Errorf(restartMsg)
	//		return nil, errors.New(restartMsg)
	//	}
	//}
	if deleteFileFromAssist {
		err := fileutils.New().Rm(fileNamePath)
		if err != nil {
			log.Errorf("assist-config-file: failed to delete uploaded config file app:%s err%s", appName, err.Error())
		}
	}
	return &EdgeReplaceConfigResp{
		AppName:            appName,
		EdgeUploadResponse: file,
		RestartMessage:     restartMsg,
	}, nil

}

// EdgeReplaceConfig
// file needs to first be uploaded to /data/store
// delete existing file
// upload the new file (use assist file upload)
// restart the service
// deleteFileFromAssist delete the file that was uploaded after the upload to the edge device is completed
func (inst *Store) EdgeReplaceConfig(hostUUID, hostName string, body *EdgeReplaceConfig) (*EdgeReplaceConfigResp, error) {
	appName := body.AppName
	fileName := body.FileName
	//restartApp := body.RestartApp
	deleteFileFromAssist := body.DeleteFileFromAssist
	if appName == "" {
		return nil, errors.New("assist-config-file: app name, cant not be empty")
	}
	if fileName == "" {
		return nil, errors.New("assist-config-file: file name, cant not be empty, try config.yml, config.json or .env")
	}
	dir := inst.App.StoreDir
	dest := fmt.Sprintf("/data/%s/config", appName)
	fileNamePath := fmt.Sprintf("%s/%s", dir, appName)
	var restartMsg = fmt.Sprintf("restarted app ok %s", appName)
	log.Infof("assist-config-file: try and upload config file to edge file:%s", dest)
	file, err := inst.EdgeUploadLocalFile(hostUUID, hostName, dir, fileName, dest)
	if err != nil {
		return nil, err
	}
	//if restartApp {
	//	_, err := inst.EdgeCtlAction(hostUUID, hostName, &installer.CtlBody{
	//		AppName: appName,
	//		Action:  "restart",
	//	})
	//	if err != nil {
	//		restartMsg = fmt.Sprintf("assist-config-file: failed to restart app:%s err%s", appName, err.Error)
	//		log.Errorf(restartMsg)
	//		return nil, errors.New(restartMsg)
	//	}
	//}
	if deleteFileFromAssist {
		err := fileutils.New().Rm(fileNamePath)
		if err != nil {
			log.Errorf("assist-config-file: failed to delete uploaded config file app:%s err%s", appName, err.Error())
		}
	}
	return &EdgeReplaceConfigResp{
		AppName:            appName,
		EdgeUploadResponse: file,
		RestartMessage:     restartMsg,
	}, nil

}
